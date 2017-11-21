package dao

import (
	"database/sql"
	"fmt"
	"log"
	//mysql driver
	"github.com/adiclepcea/furnir/models"
)

//PieceDao defines the db operations that can be done
//on a piece for persistence
type PieceDao struct {
}

//SavePiece will insert or update a piece
func (pieceDao PieceDao) SavePiece(piece models.Piece) (*models.Piece, error) {
	log.Printf("Save piece: %d\r\n", piece.ID)
	if piece.PalletsID == 0 {
		return nil, fmt.Errorf("A piece must have a pallet")
	}
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if piece.ID == 0 {
		var res sql.Result
		if piece.Essence.ID != 0 {
			res, err = db.Exec("Insert into pieces(pallets_id, essences_id,barcode,code,length,width,sheets,pallet) values(?,?,?,?,?,?,?,?)",
				piece.PalletsID,
				piece.Essence.ID,
				piece.Barcode,
				piece.Scanned.Code,
				piece.Scanned.Length,
				piece.Scanned.Width,
				piece.Scanned.SheetCount,
				piece.Scanned.OriginalPallet)
		} else {
			res, err = db.Exec("Insert into pieces(pallets_id, barcode,code, length, width, sheets,pallet) values(?,?,?,?,?,?,?)",
				piece.PalletsID,
				piece.Barcode,
				piece.Scanned.Code,
				piece.Scanned.Length,
				piece.Scanned.Width,
				piece.Scanned.SheetCount,
				piece.Scanned.OriginalPallet)
		}
		if err != nil {
			log.Printf("Error saving piece: %s\r\n", err.Error())
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		piece.ID = id

	} else {
		_, err = db.Exec("Update pieces set pallets_id=?, essences_id=?, barcode=?,code=?, length=?, width=?, sheets=?,pallet=? where pieces_id=?",
			piece.PalletsID,
			piece.Essence.ID,
			piece.Barcode,
			piece.Scanned.Code,
			piece.Scanned.Length,
			piece.Scanned.Width,
			piece.Scanned.SheetCount,
			piece.Scanned.OriginalPallet,
			piece.ID)
		if err != nil {
			log.Printf("Error saving piece: %s\r\n", err.Error())
			return nil, err
		}
	}

	return pieceDao.FindPieceByID(piece.ID)
}

//FindPieceByID finds the piece with the selected id
func (pieceDao PieceDao) FindPieceByID(id int64) (*models.Piece, error) {
	log.Printf("Find piece by id: %d\r\n", id)
	piece := models.Piece{}
	piece.Essence = models.Essence{}
	piece.Scanned = models.ScannedPiece{}
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query(`Select p.pieces_id, p.pallets_id, p.barcode, p.code, p.length, p.width,p.sheets,p.pallet, e.essences_id, e.name, e.code  
		from pieces p left outer join essences e on p.essences_id=e.essences_id where pieces_id=?`, id)

	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&piece.ID,
			&piece.PalletsID,
			&piece.Barcode,
			&piece.Scanned.Code,
			&piece.Scanned.Length,
			&piece.Scanned.Width,
			&piece.Scanned.SheetCount,
			&piece.Scanned.OriginalPallet,
			&piece.Essence.ID,
			&piece.Essence.Name,
			&piece.Essence.Code)

		return &piece, nil
	}

	return nil, nil
}

//FindPiecesByBarcode finds the pieces by barcode
func (pieceDao PieceDao) FindPiecesByBarcode(code string) ([]models.Piece, error) {
	log.Printf("Find piece by barcode: %s\r\n", code)
	var pieces []models.Piece
	pieces = make([]models.Piece, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query(`Select p.pieces_id, p.pallets_id, p.barcode, p.code, p.length, p.width,p.sheets,p.pallet, e.essences_id, e.name, e.code  
		from pieces p left outer join essences e on p.essences_id=e.essences_id where p.barcode=?`, code)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		piece := models.Piece{}
		piece.Essence = models.Essence{}
		piece.Scanned = models.ScannedPiece{}
		res.Scan(&piece.ID,
			&piece.PalletsID,
			&piece.Barcode,
			&piece.Scanned.Code,
			&piece.Scanned.Length,
			&piece.Scanned.Width,
			&piece.Scanned.SheetCount,
			&piece.Scanned.OriginalPallet,
			&piece.Essence.ID,
			&piece.Essence.Name,
			&piece.Essence.Code)
		pieces = append(pieces, piece)
	}
	return pieces, nil
}

//FindPiecesByPalletsID finds the pieces inside the selected pallet
func (pieceDao PieceDao) FindPiecesByPalletsID(palletsID int64) ([]models.Piece, error) {
	log.Printf("Find pieces by pallets ID: %d\r\n", palletsID)
	var pieces []models.Piece
	pieces = make([]models.Piece, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query(`Select p.pieces_id, p.pallets_id, p.barcode, p.code, p.length, p.width,p.sheets,p.pallet, e.essences_id, e.name, e.code  
		from pieces p left outer join essences e on p.essences_id=e.essences_id where p.pallets_id=?`, palletsID)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		piece := models.Piece{}
		piece.Essence = models.Essence{}
		piece.Scanned = models.ScannedPiece{}
		res.Scan(&piece.ID,
			&piece.PalletsID,
			&piece.Barcode,
			&piece.Scanned.Code,
			&piece.Scanned.Length,
			&piece.Scanned.Width,
			&piece.Scanned.SheetCount,
			&piece.Scanned.OriginalPallet,
			&piece.Essence.ID,
			&piece.Essence.Name,
			&piece.Essence.Code)
		pieces = append(pieces, piece)
	}
	return pieces, nil
}

//FindAllPieces returns all pieces in the system
func (pieceDao PieceDao) FindAllPieces() ([]models.Piece, error) {
	log.Println("Find all pieces")
	var pieces []models.Piece
	pieces = make([]models.Piece, 0)

	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query(`Select p.pieces_id, p.pallets_id, p.barcode, p.code, p.length, p.width,p.sheets,p.pallet, e.essences_id, e.name, e.code  
		from pieces p left outer join essences e on p.essences_id=e.essences_id`)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		piece := models.Piece{}
		piece.Essence = models.Essence{}
		piece.Scanned = models.ScannedPiece{}
		res.Scan(&piece.ID,
			&piece.PalletsID,
			&piece.Barcode,
			&piece.Scanned.Code,
			&piece.Scanned.Length,
			&piece.Scanned.Width,
			&piece.Scanned.SheetCount,
			&piece.Scanned.OriginalPallet,
			&piece.Essence.ID,
			&piece.Essence.Name,
			&piece.Essence.Code)
		pieces = append(pieces, piece)
	}
	return pieces, nil
}

//DeletePieceByID deletes the piece having the passed id
func (pieceDao PieceDao) DeletePieceByID(id int64) error {
	log.Printf("Delete piece by id: %d\r\n", id)
	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("Delete from pieces where pieces_id=?", id)
	if err != nil {
		return err
	}
	return nil
}

//TransferPieceByBarcode will transfer a piece with the giver barcode between two pallets
func (pieceDao PieceDao) TransferPieceByBarcode(code string, srcPalletID int64, destPalletID int64) error {
	log.Printf("Transfer piece %s from %d to %d\r\n", code, srcPalletID, destPalletID)
	var palletRepo PalletDao

	db, err := InitDB()
	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	dstPallet, err := palletRepo.FindPalletByID(destPalletID)

	if err != nil {
		return err
	}

	res, err := db.Query(`Select p.pieces_id, p.pallets_id, p.barcode, p.code, p.length, p.width,p.sheets,p.pallet, e.essences_id, e.name, e.code  
		from pieces p left outer join essences e on p.essences_id=e.essences_id where p.barcode=? and p.pallets_id=?`, code, srcPalletID)
	if err != nil {
		return err
	}
	var piece models.Piece
	if res.Next() {
		piece = models.Piece{}
		piece.Essence = models.Essence{}
		piece.Scanned = models.ScannedPiece{}
		res.Scan(&piece.ID,
			&piece.PalletsID,
			&piece.Barcode,
			&piece.Scanned.Code,
			&piece.Scanned.Length,
			&piece.Scanned.Width,
			&piece.Scanned.SheetCount,
			&piece.Scanned.OriginalPallet,
			&piece.Essence.ID,
			&piece.Essence.Name,
			&piece.Essence.Code)
	}

	if piece.ID != 0 {
		_, err = db.Exec("update pieces set pallets_id=? where pieces_id=?", destPalletID, piece.ID)
	} else {
		sc, err := models.ScannedPiece{}.NewFromScan(code)
		if err != nil {
			return err
		}
		piece.Scanned = *sc
		piece.Barcode = code
		piece.PalletsID = destPalletID
		piece.Essence = dstPallet.Essence

		_, err = pieceDao.SavePiece(piece)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()

	return err
}
