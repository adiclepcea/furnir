package models

//Transfer reprezents the data for transfering a piece between two pallets
type Transfer struct {
	SourcePalletID int64  `json:"source_pallet_id"`
	DestPalletID   int64  `json:"dest_pallet_id"`
	PieceBarcode   string `json:"piece"`
}
