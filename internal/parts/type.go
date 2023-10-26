package parts

import (
	"database/sql"
	"path/filepath"

	"github.com/wisniewski91/clothes-factory/internal/utils"
)

const tableName = "parts"

var (
	modulePath = []string{"modules", "parts"}
	routePath  = "parts"
)

type PartRecord struct {
	Id        int
	Name      string
	Status    int8
	UpdatedAt string
	CreatedAt string
}

type PartsRepository interface {
	Fetch(id int) (*PartRecord, error)
	FetchAll() (*[]PartRecord, error)
	FetchPage(page int, limit int) (*[]PartRecord, error)
	Insert(data *PartRecord) (*PartRecord, error)
	Update(data *PartRecord) (*PartRecord, error)
	Delete(id int) error
}

type MysqlPartsRepository struct {
	DB *sql.DB
}

func NewMysqlPartsRepository(db *sql.DB) *MysqlPartsRepository {
	return &MysqlPartsRepository{
		DB: db,
	}
}

type PartsService struct {
	repository PartsRepository
}

func NewPartsService(r PartsRepository) *PartsService {
	return &PartsService{
		repository: r,
	}
}

func renderPath(template ...string) string {
	path := append(modulePath, template...)
	return filepath.Join(path...)
}

func (e PartRecord) Validate() error {
	if e.Name == "" {
		return utils.ValidateError{
			Code:    -10,
			Object:  "Part",
			Message: "Name not provided",
		}
	}
	return nil
}
