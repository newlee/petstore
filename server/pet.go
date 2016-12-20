package server

type Pet struct {
	Name string `json:"name" form:"name" query:"name"`
}

type PetRepo interface {
	Create(Pet)
	All() []Pet
}

type InMemoryPetRepo struct {
	pets []Pet
}

func (repo *InMemoryPetRepo) Create(p Pet) {
	repo.pets = append(repo.pets, p)
}

func (repo *InMemoryPetRepo) All() []Pet {
	return repo.pets
}

func NewInMemoryPetRepo() *InMemoryPetRepo {
	return &InMemoryPetRepo{pets:make([]Pet, 0)}
}

