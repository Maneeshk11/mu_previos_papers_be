package store

import (
	"fmt"
	"math/big"
	"mu_previous_papers_be/model"
	"mu_previous_papers_be/server"
	"sync"

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
