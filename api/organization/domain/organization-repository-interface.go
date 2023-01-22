package domain

type IOrganizationRepository interface {
	Create(*Organization) error
	Delete(int64) error
	Update(*Organization) error
	GetList(page, limit uint) ([]Organization, int, error)
	GetAll() ([]Organization, error)
	GetByID(id int64) (*Organization, error)
}
