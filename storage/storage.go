package storage

type Product interface {
	Create(id uint, description string) error
	Description() string
}

func ProductDescription(p Product) string {
	return p.Description()
}

type ProductCatalog interface {
	Insert(product Product) error
	Find(id uint) Product
}

func FindProductDescription(pc ProductCatalog, id uint) string {
	return pc.Find(id).Description()
}
