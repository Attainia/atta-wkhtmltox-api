package options

import (
	"regexp"
)

var pageSizes []string = []string{
	"A0",
	"A1",
	"A2",
	"A3",
	"A4",
	"A5",
	"A6",
	"A7",
	"A8",
	"A9",
	"B0",
	"B1",
	"B2",
	"B3",
	"B4",
	"B5",
	"B6",
	"B7",
	"B8",
	"B9",
	"B10",
	"C5E",
	"Comm10E",
	"DLE",
	"Executive",
	"Folio",
	"Ledger",
	"Legal",
	"Letter",
	"Tabloid",
}

type PDFOptions struct {
	*Options
}

func (o *PDFOptions) GetCollateFlag() []string {
	if collate, exists := o.options["collate"]; exists {
		if collate == "0" {
			return []string{"--no-collate"}
		}
	}
	return []string{"--collate"}
}

func (o *PDFOptions) GetCopiesFlag() []string {
	if copies, exists := o.options["copies"]; exists {
		if match, _ := regexp.MatchString("^[0-9]*$", copies); match {
			return []string{"--copies", copies}
		}
	}

	return []string{"--copies", "1"}
}

func (o *PDFOptions) GetGrayscaleFlag() []string {
	if grayscale, exists := o.options["grayscale"]; exists {
		if grayscale == "1" {
			return []string{"--grayscale"}
		}
	}

	return []string{}
}

func (o *PDFOptions) GetLowQualityFlag() []string {
	if lowquality, exists := o.options["lowquality"]; exists {
		if lowquality == "1" {
			return []string{"--lowquality"}
		}
	}

	return []string{}
}

func (o *PDFOptions) GetOrientationFlag() []string {
	if orientation, exists := o.options["orientation"]; exists {
		if orientation == "Landscape" {
			return []string{"--orientation", "Landscape"}
		}
	}

	return []string{"--orientation", "Portrait"}
}

func (o *PDFOptions) GetPageSizeFlag() []string {
	if pageSize, exists := o.options["page-size"]; exists {
		for i := range pageSizes {
			if pageSize == pageSizes[i] {
				return []string{"--page-size", pageSize}
			}
		}
	}

	return []string{"--page-size", "A4"}
}

func PDFOptionsFromHeader(header string) (*PDFOptions, error) {
	opts, err := OptionsFromHeader(header)
	if err != nil {
		return nil, err
	}
	return &PDFOptions{opts}, nil
}
