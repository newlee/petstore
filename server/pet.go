package server

type Pet struct {
	Name string `json:"name" form:"name" query:"name"`
}

type PetRepo interface {
	Create(Pet)
	All() []Pet
}

type InMemoryPetRepo struct {
	products []Pet
}

func (repo *InMemoryPetRepo) Create(p Pet) {
	repo.products = append(repo.products, p)
}

func (repo *InMemoryPetRepo) All() []Pet {
	return repo.products
}

func NewInMemoryPetRepo() *InMemoryPetRepo {
	return &InMemoryPetRepo{products:make([]Pet, 0)}
}