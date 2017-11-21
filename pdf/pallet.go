package pdf

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode"
	"log"
	"fmt"
	"io"
	"image/jpeg"
	"bytes"
	"github.com/adiclepcea/furnir/models"
	"github.com/jung-kurt/gofpdf"
)

//GeneratePalletPDF will generate a PDF containing data about the selected pallet
func GeneratePalletPDF(pallet models.Pallet, pieces []models.Piece, w io.WriteCloser ) error{
	
	bc,err := code128.Encode(fmt.Sprintf("%012d",pallet.ID))
	if err!=nil{
		log.Println(err.Error())
	}
	
	bcScaled, err := barcode.Scale(bc,200,40)
	if err!=nil{
		log.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	jpeg.Encode(buf,bcScaled,nil)

	pdf := gofpdf.New("P","mm", "A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","B",10)
	io := gofpdf.ImageOptions{ImageType:"JPG",ReadDpi:false}
	pdf.RegisterImageOptionsReader("palletBarcode",io,bytes.NewReader(buf.Bytes()))
	pdf.ImageOptions("palletBarcode",80,-1,0,0,true,io,0,"")
	pdf.Ln(1)
	pdf.WriteAligned(0,3,fmt.Sprintf("%012d %s",pallet.ID, pallet.Essence.Name),"C")
	pdf.Ln(20)
	
	for _, piece := range pieces{
		bcPiece, err := code128.Encode(piece.Barcode)
		if err!=nil{
			log.Println(err.Error())
			continue
		}	
		bcScaled, err := barcode.Scale(bcPiece,200,20)
		if err!=nil{
			log.Println(err.Error())
			continue
		}
		buf := new(bytes.Buffer)
		jpeg.Encode(buf,bcScaled,nil)
		imgName:=fmt.Sprintf("piece%d",piece.ID)
		pdf.RegisterImageOptionsReader(imgName,io,bytes.NewReader(buf.Bytes()))
		pdf.ImageOptions(imgName,-1,-1,0,0,true,io,0,"")
		pdf.WriteAligned(0,-2,fmt.Sprintf("%s,P=%d, L=%d, l=%d, Foi=%d",piece.Scanned.Code,piece.Scanned.OriginalPallet,piece.Scanned.Length,piece.Scanned.Width, piece.Scanned.SheetCount),"R")
		pdf.Line(10,pdf.GetY()+1,200,pdf.GetY()+1)
		pdf.Ln(3)
	}

	return pdf.OutputAndClose(w)
}