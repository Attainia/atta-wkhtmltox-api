package internal

import (
	"atta-wkhtmltox-api/options"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type WkhtmltoxView struct{}

func (v *WkhtmltoxView) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		v.convertContent(w, r)
	}
}

func (v *WkhtmltoxView) convertContent(w http.ResponseWriter, r *http.Request) {
	var err error
	config := GetConfig()
	workDir := filepath.Join(config.WorkDir, strings.Replace(uuid.New().String(), "-", "", 4))
	htmlPath := filepath.Join(workDir, "f.html")
	outPath := filepath.Join(workDir, "f.pdf")

	var flags = []string{"--quiet"}

	// validate content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "text/html" {
		http.Error(w, "Content-Type header must be `text/html`.", http.StatusBadRequest)
		return
	}

	// validate accept type
	acceptParts := strings.Split(r.Header.Get("Accept"), ";")
	acceptType := acceptParts[0]

	// set flags for accepted type
	if acceptType == "application/pdf" || acceptType == "application/pdf;" {
		var opts *options.PDFOptions
		opts, err = options.PDFOptionsFromHeader(r.Header.Get("Accept"))
		if err != nil {
			http.Error(w, "Received bad Accept options.", http.StatusBadRequest)
			return
		}

		flags = append(flags, opts.GetCollateFlag()...)
		flags = append(flags, opts.GetCopiesFlag()...)
		flags = append(flags, opts.GetGrayscaleFlag()...)
		flags = append(flags, opts.GetLowQualityFlag()...)
		flags = append(flags, opts.GetOrientationFlag()...)
		flags = append(flags, opts.GetPageSizeFlag()...)
	} else {
		http.Error(w, "Accept header is not a supported format.", http.StatusBadRequest)
		return
	}

	// create working directory
	if err = os.Mkdir(workDir, 0777); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer func() { _ = os.RemoveAll(workDir) }()

	// write html from request to disk
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if err = ioutil.WriteFile(htmlPath, content, 0644); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	//convert to pdf
	var args = []string{}
	args = append(args, config.WKHTMLTOPDFPath)
	args = append(args, flags...)
	args = append(args, htmlPath, outPath)
	if err := exec.Command("xvfb-run", args...).Run(); err != nil {
		http.Error(w, "error processing content.", http.StatusInternalServerError)
		return
	}

	// write output to response
	f, err := os.Open(outPath)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	out, err := ioutil.ReadAll(f)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	if _, err = w.Write(out); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
