package models

// Nama struct diawali huruf kapital agar bisa dipakai di folder lain
type Dimsum struct {
	ID    string `json:"id" firestore:"-"`
	Nama  string `json:"nama" firestore:"nama"`
	Harga int    `json:"harga" firestore:"harga"`
	Stok  int    `json:"stok" firestore:"stok"`
}
