package service

import (
	"net/http"
	"github.com/adiclepcea/furnir/pdf"
	"net/http/httputil"
)
//PrintOperations generates a PDF with the barcodes used for operations
func PrintOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type","application/pdf")
	operationBarcodes := []string{"palet sursa", "palet dest","print sursa", "print dest","palet nou"}
	pdf.GenerateSimplePDF(operationBarcodes, "Operatii paleti", httputil.NewChunkedWriter(w))	
}