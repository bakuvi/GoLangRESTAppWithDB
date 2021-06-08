package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	Id   int       `json:"id"`
	Date time.Time `json:"date"`
}
type GetUserDate struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}
type Service struct {
	Conn *sqlx.DB
}

func (s *Service) Add(tw User) (id int64 , err error) {
	res, err := s.Conn.Exec("INSERT INTO tabelname (tweet_author, text) VALUES (?, ?)", tw.Id, tw.Date)
	if err != nil {
		return 0, fmt.Errorf("can't insert tweet into db: %w", err)
	}

	return res.LastInsertId()
}

func (s *Service) Get(dateFrom time.Time, dateTo time.Time) (id []int64,err error) {

var ids []int64
	err= s.Conn.Select(&ids, `SELECT id FROM first_table WHERE  date BETWEEN ? AND ?;`, dateFrom, dateTo )

	if err != nil {
		return nil, fmt.Errorf("can't query row: %w", err)
	}

return ids, nil
}

