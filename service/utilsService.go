package service

import (
	"net/http"
	"net/http/httputil"

	"github.com/adiclepcea/furnir/pdf"
)

//PrintOperations generates a PDF with the barcodes used for operations
func PrintOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/pdf")
	operationBarcodes := []string{
		"palet sursa",
		"palet dest1",
		"palet dest2",
		"palet dest3",
		"palet dest4",
		"print sursa",
		"print dest1",
		"print dest2",
		"print dest3",
		"print dest4",
		"palet nou"}
	pdf.GenerateSimplePDF(operationBarcodes, "Operatii paleti", httputil.NewChunkedWriter(w))
}
