package models

import (
	"testing"
)

func TestNewFromScanTooShort(t *testing.T) {
	_, err := ScannedPiece{}.NewFromScan("110021440000034202551")
	if err == nil {
		t.Error("Expected to get error while passing a string that is not long enough. Got nil")
	}
}

func TestNewFromScanInvalidOriginalPallet(t *testing.T) {
	_, err := ScannedPiece{}.NewFromScan("11002144000g034202551")
	if err == nil {
		t.Error("Expected to get error while passing a string with invalid OriginalPalet. Got nil")
	}
}
func TestNewFromScanInvalidSheetCount(t *testing.T) {
	_, err := ScannedPiece{}.NewFromScan("1100214400000342G2551")
	if err == nil {
		t.Error("Expected to get error while passing a string with invalid SheetCount. Got nil")
	}
}

func TestNewFromScanInvalidLength(t *testing.T) {
	_, err := ScannedPiece{}.NewFromScan("110021440000034202G51")
	if err == nil {
		t.Error("Expected to get error while passing a string with invalid Length. Got nil")
	}
}

func TestNewFromScanInvalidWidth(t *testing.T) {
	_, err := ScannedPiece{}.NewFromScan("11002144000003420255D")
	if err == nil {
		t.Error("Expected to get error while passing a string with invalid width. Got nil")
	}
}
func TestNewFromScan(t *testing.T) {
	sp, err := ScannedPiece{}.NewFromScan("1100214400000342025514")

	if err != nil {
		t.Errorf("Expected no error while creating a new ScannedPiece with a correct string. Got %s", err.Error())
	}

	if sp.Code != "002144-11" {
		t.Errorf("Expected %s, got %s for Code", "002144-11", sp.Code)
	}

	if sp.OriginalPallet != 34 {
		t.Errorf("Expected %d, got %d  for OriginalPallet", 34, sp.OriginalPallet)
	}

	if sp.SheetCount != 20 {
		t.Errorf("Expected %d, got %d for SheetCount", 20, sp.SheetCount)
	}

	if sp.Length != 255 {
		t.Errorf("Expected %d, got %d for Length", 255, sp.Length)
	}

	if sp.Width != 14 {
		t.Errorf("Expected %d, got %d for Length", 14, sp.Width)
	}
}

func TestScannedPieceArea(t *testing.T) {
	sp, err := ScannedPiece{}.NewFromScan("1100214400000342025514")

	if err != nil {
		t.Errorf("Expected no error while creating a new ScannedPiece with a correct string. Got %s", err.Error())
	}
	if sp.Area() != 7.14 {
		t.Errorf("Expected a calculated area of: %f, got %f", 7.14, sp.Area())
	}
}
