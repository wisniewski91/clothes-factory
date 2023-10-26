package jobtypes

import (
	"database/sql"
	"path/filepath"

	"github.com/wisniewski91/clothes-factory/internal/utils"
)

const tableName = "job_types"

var (
	modulePath = []string{"modules", "jobTypes"}
	routePath  = "jobtypes"
)

type JobtypeRecord struct {
	Id        int
	Name      string
	Status    int8
	UpdatedAt string
	CreatedAt string
}

type JobtypesRepository interface {
	Fetch(id int) (*JobtypeRecord, error)
	FetchAll() (*[]JobtypeRecord, error)
	FetchPage(page int, limit int) (*[]JobtypeRecord, error)
	Insert(data *JobtypeRecord) (*JobtypeRecord, error)
	Update(data *JobtypeRecord) (*JobtypeRecord, error)
	Delete(id int) error
}

type MysqlJobtypesRepository struct {
	DB *sql.DB
}

func NewMysqlJobtypesRepository(db *sql.DB) *MysqlJobtypesRepository {
	return &MysqlJobtypesRepository{
		DB: db,
	}
}

type JobtypesService struct {
	repository JobtypesRepository
}

func NewJobtypesService(r JobtypesRepository) *JobtypesService {
	return &JobtypesService{
		repository: r,
	}
}

func renderPath(template ...string) string {
	path := append(modulePath, template...)
	return filepath.Join(path...)
}

func (jt JobtypeRecord) Validate() error {
	if jt.Name == "" {
		return utils.ValidateError{
			Code:    -10,
			Object:  "JobType",
			Message: "Name not provided",
		}
	}
	return nil
}
