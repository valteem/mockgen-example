package storage

type Product interface {
	Insert() (int, error)
	Find(int) (*Product, error)
}
