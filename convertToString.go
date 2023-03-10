package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"code.sajari.com/docconv"
)

func openAndConvertFileToString(s string) string {
	if isFileEmpty(s) {
		return ""
	}

	f, err := os.Open(s)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer f.Close()
	fmt.Println(f.Name())
	// Perform necessary conversions to string
	// if ODT, doc, etc
	var text string
	switch {
	case strings.HasSuffix(f.Name(), ".odt"):
		text = odtToString(f)
	case strings.HasSuffix(f.Name(), ".doc"):
		text = docToString(f)
	case strings.HasSuffix(f.Name(), ".docx"):
		text = docxToString(f)
	case strings.HasSuffix(f.Name(), ".pdf"):
		text = pdfToString(f)
	}
	return text
}

func isFileEmpty(s string) bool {
	sizeCheck, err := os.Stat(s)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if sizeCheck.Size() == 0 {
		return true
	}
	return false
}
func odtToString(f io.Reader) string {
	body, _, err := docconv.ConvertODT(f)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return body
}

func docToString(f io.Reader) string {
	body, _, err := docconv.ConvertDoc(f)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return body
}

func docxToString(f io.Reader) string {
	body, _, err := docconv.ConvertDocx(f)
	if err != nil {
		fmt.Println("Hmm? ", err)
		os.Exit(1)
	}
	return body
}

func pdfToString(f io.Reader) string {
	body, _, err := docconv.ConvertPDF(f)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return body
}
