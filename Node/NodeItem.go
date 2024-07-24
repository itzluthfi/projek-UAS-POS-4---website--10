package Node

type NodeItem struct {
	Id          int
	Nama        string
	JmlStock    int
	Harga       int
	HargaDiskon int
	Diskon      int
	CreateAt    string
}

type ItemLL struct {
	Item NodeItem
	Next *ItemLL
}
