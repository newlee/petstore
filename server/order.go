package server

type Order struct {
	PetName   string `json:"petName" form:"petName" query:"petName"`
	GuestName string `json:"guestName" form:"guestName" query:"guestName"`
	Telephone string `json:"telephone" form:"telephone" query:"telephone"`
	Address   string `json:"address" form:"address" query:"address"`
}

type OrderRepo interface {
	Create(Order)
	Find(string) *Order
}

type InMemoryOrderRepo struct {
	orders []Order
}

func (repo *InMemoryOrderRepo) Create(o Order) {
	repo.orders = append(repo.orders, o)
}

func (repo *InMemoryOrderRepo) Find(telephone string) *Order {
	for _, o := range repo.orders {
		if (o.Telephone == telephone) {
			return &o
		}
	}
	return nil
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{orders:make([]Order, 0)}
}


