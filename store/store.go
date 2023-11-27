package store

import (
	"fmt"
	"math/big"
	"mu_previous_papers_be/model"
	"mu_previous_papers_be/server"
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store interface {
	server.Store
	Gorm() *gorm.DB
}

type store struct {
	mu    *sync.RWMutex
	cache map[string]*big.Float
	db    *gorm.DB
}

func (s *store) Gorm() *gorm.DB {
	return s.db
}

func NewStore(db *gorm.DB) Store {
	return &store{
		mu:    new(sync.RWMutex),
		cache: make(map[string]*big.Float),
		db:    db,
	}
}

func (s *store) HealthCheck() error {
	return nil
}

func (s *store) GetSubjects() []string {
	var res []string
	query := `select distinct subject_name from qpapers_info`
	result := s.db.Raw(query).Scan(&res)
	if result.Error != nil {
		fmt.Println("Error querying the database:", result.Error)
		return res
	}

	return res
}

func (s *store) GetSubjectCodes() []string {
	var res []string
	query := `select distinct subject_code from qpapers_info`
	result := s.db.Raw(query).Scan(&res)
	if result.Error != nil {
		fmt.Println("Error querying the database:", result.Error)
		return res
	}
	return res
}

func (s *store) GetTitles(subject, code, year string) []model.QpapersInfo {
	var query string
	var res []model.QpapersInfo

	if subject != "" {
		if year != "" {
			query = fmt.Sprintf(`select q1.subject_code, q1.subject_name, q1.semester, q1.exam_type, q1.exam_occasion, q1.exam_year, q1.branch, q2.file_path 
			from qpapers_info as q1
			left join qpapers_loc as q2 on q1.id = q2.qpapers_id
			where q1.subject_name='%s' and q1.exam_year='%s'`, subject, year)

		} else {
			query = fmt.Sprintf(`select q1.subject_code, q1.subject_name, q1.semester, q1.exam_type, q1.exam_occasion, q1.exam_year, q1.branch, q2.file_path 
			from qpapers_info as q1
			left join qpapers_loc as q2 on q1.id = q2.qpapers_id
			where q1.subject_name='%s'`, subject)
		}
	} else if code != "" {
		if year != "" {
			query = fmt.Sprintf(`select q1.subject_code, q1.subject_name, q1.semester, q1.exam_type, q1.exam_occasion, q1.exam_year, q1.branch, q2.file_path 
			from qpapers_info as q1
			left join qpapers_loc as q2 on q1.id = q2.qpapers_id
			where q1.subject_code='%s' and exam_year='%s'`, code, year)
		} else {
			query = fmt.Sprintf(`select q1.subject_code, q1.subject_name, q1.semester, q1.exam_type, q1.exam_occasion, q1.exam_year, q1.branch, q2.file_path 
			from qpapers_info as q1
			left join qpapers_loc as q2 on q1.id = q2.qpapers_id
			where q1.subject_code='%s'`, code)
		}
	}
	result := s.db.Raw(query).Scan(&res)
	if result.Error != nil {
		fmt.Println("Error querying the database:", result.Error)
		return res
	}

	return res
}

func (s *store) PutInDB(obj model.QpapersInfo) string {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "Failed to generate UUID"
	}
	uuidString := uuidObj.String()

	query_1 := fmt.Sprintf(`insert into qpapers_info(id, subject_code, subject_name, semester, exam_type, exam_occasion, exam_year, branch) 
	values ('%s', '%s', '%s', %d, '%s', '%s', %d, '%s')`, uuidString, obj.Subject_code, obj.Subject_name, obj.Semester, obj.Exam_type, obj.Exam_occasion, obj.Exam_year, obj.Branch)

	query_2 := fmt.Sprintf(`insert into qpapers_loc(qpapers_id, file_path) values ('%s', '%s')`, uuidString, obj.File_path)

	result1 := s.db.Exec(query_1)
	if result1.Error != nil {
		fmt.Println("Error querying the database:", result1.Error)
		return "Failed to add to qpapers_info"
	}

	result2 := s.db.Exec(query_2)
	if result2.Error != nil {
		fmt.Println("Error querying the database:", result2.Error)
		return "Failed to add to qpapers_loc"
	}

	return "Successfully added to db"
}
