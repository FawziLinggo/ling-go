package domain

type Category struct {
	// Kesalahan kalau tidak huruf besar depannay (id, name) gak bisa
	// di akses dari luar wak
	Id   int
	Name string
}
