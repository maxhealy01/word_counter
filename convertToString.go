package main

import (
	"fmt"
	"io"
	"os"

	"code.sajari.com/docconv"
)

//	func convertToString(f io.Reader) string {
//		var text string
//			switch {
//		case strings.HasSuffix(f.Name(), ".odt"):
//			text = odtToString(f)
//		case strings.HasSuffix(f.Name(), ".doc"):
//			text = docToString(f)
//		case strings.HasSuffix(f.Name(), ".docx"):
//			text = docxToString(f)
//		case strings.HasSuffix(f.Name(), ".pdf"):
//			text = pdfToString(f)
//		}
//		return text
//	}
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
		fmt.Println("Error: ", err)
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
