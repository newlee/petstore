package server

type Product struct {
	Name string `json:"name" form:"name" query:"name"`
}

type ProductRepo interface {
	Create(Product)
	All() []Product
}

type InMemoryProductRepo struct {
	products []Product
}

func (repo *InMemoryProductRepo) Create(p Product) {
	repo.products = append(repo.products, p)
}

func (repo *InMemoryProductRepo) All() []Product {
	return repo.products
}

func NewInMemoryProductRepo() *InMemoryProductRepo  {
	return &InMemoryProductRepo{products:make([]Product, 0)}
}