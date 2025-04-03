package storage

type Product interface {
	Create(description string) error
	Description() string
}

func ProductDescription(p Product) string {
	return p.Description()
}
