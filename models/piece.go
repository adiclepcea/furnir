package models

import (
	"fmt"
	"strconv"
)

//1100214400000342025514
//002144-11  ST  34
//20 255 14 7.14

//ScannedPiece represents the contents of a scanned package
type ScannedPiece struct {
	Code           string
	Essence        string
	OriginalPallet int64
	SheetCount     int
	Length         int
	Width          int
}

//NewFromScan will return a new ScannedPiece from the string passed in
func (scannedPiece ScannedPiece) NewFromScan(Scan string) (*ScannedPiece, error) {
	var err error
	newScan := ScannedPiece{}

	//the minimum length is 22 for a code
	if len(Scan) < 22 {
		return nil, fmt.Errorf("Code %s is to small", Scan)
	}

	newScan.Code = Scan[2:8] + "-" + Scan[0:2]
	newScan.OriginalPallet, err = strconv.ParseInt(Scan[8:15], 10, 64)
	if err != nil {
		return nil, err
	}

	newScan.SheetCount, err = strconv.Atoi(Scan[15:17])
	if err != nil {
		return nil, err
	}

	newScan.Length, err = strconv.Atoi(Scan[17:20])
	if err != nil {
		return nil, err
	}

	newScan.Width, err = strconv.Atoi(Scan[20:22])
	if err != nil {
		return nil, err
	}

	return &newScan, nil
}

//Area will calculate the total area of the scanned piece
func (scannedPiece *ScannedPiece) Area() float32 {
	return float32(scannedPiece.Length) * float32(scannedPiece.Width) / float32(10000) * float32(scannedPiece.SheetCount)
}