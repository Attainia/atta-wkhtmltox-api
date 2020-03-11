package internal

import (
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

	// validate content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "text/html" {
		http.Error(w, "Content-Type header must be `text/html`.", http.StatusBadRequest)
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
	if _, err = exec.Command("xvfb-run", config.WKHTMLTOPDFPath, htmlPath, outPath).Output(); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
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
