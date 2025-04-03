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

type StorageRoom interface {
	Capacity(floor int) float32
}

func StorageShare(sr StorageRoom, floor int, share float32) float32 {
	return sr.Capacity(floor) * share
}

type StoredItemsCount interface {
	Count() int
}

func ItemsStored(c StoredItemsCount) int {
	return c.Count()
}
