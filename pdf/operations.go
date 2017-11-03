package pdf

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode"
	"log"
	"io"
	"image/jpeg"
	"bytes"
	"github.com/jung-kurt/gofpdf"
)

//GenerateSimplePDF will generate a PDF containing the barcodes and with the title
func GenerateSimplePDF(barcodes []string, title string, w io.WriteCloser) error{

	pdf := gofpdf.New("P","mm", "A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","B",20)
	pdf.WriteAligned(0,3,title,"C")
	pdf.SetFont("Arial","",12)
	pdf.Ln(20)
	
	
	imgOpt := gofpdf.ImageOptions{ImageType:"JPG",ReadDpi:false}
	for _, ob := range barcodes{
		bc, err := code128.Encode(ob)
		if err!=nil{
			log.Println(err.Error())
			continue
		}	
		bcScaled, err := barcode.Scale(bc,400,60)
		if err!=nil{
			log.Println(err.Error())
			continue
		}
		buf := new(bytes.Buffer)
		jpeg.Encode(buf,bcScaled,nil)
		pdf.RegisterImageOptionsReader(ob,imgOpt,bytes.NewReader(buf.Bytes()))
		pdf.ImageOptions(ob,55,-1,0,0,true,imgOpt,0,"")
		pdf.Ln(3)
		pdf.WriteAligned(0,-2,ob,"C")
		pdf.Ln(10)
		pdf.Line(10,pdf.GetY()+1,200,pdf.GetY()+1)
		pdf.Ln(3)
	}

	return pdf.OutputAndClose(w)
}