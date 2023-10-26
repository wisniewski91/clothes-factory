package clients

import (
	"database/sql"
	"path/filepath"

	"github.com/wisniewski91/clothes-factory/internal/utils"
)

const tableName = "clients"

var (
	modulePath = []string{"modules", "clients"}
	routePath  = "clients"
)

type ClientRecord struct {
	Id          int
	Name        string
	AddressData utils.AddressData
	Status      int8
	UpdatedAt   string
	CreatedAt   string
}

type ClientsRepository interface {
	Fetch(id int) (*ClientRecord, error)
	FetchAll() (*[]ClientRecord, error)
	FetchPage(page int, limit int) (*[]ClientRecord, error)
	Insert(data *ClientRecord) (*ClientRecord, error)
	Update(data *ClientRecord) (*ClientRecord, error)
	Delete(id int) error
}

type MysqlClientsRepository struct {
	DB *sql.DB
}

func NewMysqlClientsRepository(db *sql.DB) *MysqlClientsRepository {
	return &MysqlClientsRepository{
		DB: db,
	}
}

type ClientsService struct {
	repository ClientsRepository
}

func NewClientsService(r ClientsRepository) *ClientsService {
	return &ClientsService{
		repository: r,
	}
}

func renderPath(template ...string) string {
	path := append(modulePath, template...)
	return filepath.Join(path...)
}

func (c ClientRecord) Validate() error {
	if c.Name == "" {
		return utils.ValidateError{
			Code:    -10,
			Object:  "ClientRecord",
			Message: "Name not provided",
		}
	}

	return nil
}
