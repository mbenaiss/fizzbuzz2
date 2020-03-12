package repository

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mbenaiss/fizzbuzz/internal/fizzbuzz"
)

type Repository struct {
	db *sql.DB
}

//NewClient initializes a new database connection
func New(path string) (*Repository, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	err = createTable(db)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db: db,
	}, nil
}

func createTable(db *sql.DB) error {
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS 
			stats (params TEXT PRIMARY KEY, hits INTEGER)
		`)

	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

//Get returns the most used params
func (c *Repository) Get() (*fizzbuzz.FizzBuzz, int, error) {
	rows, _ := c.db.Query("SELECT params, MAX(hits) FROM stats")
	var params string
	var hits int
	for rows.Next() {
		err := rows.Scan(&params, &hits)
		if err != nil {
			return nil, 0, err
		}
	}
	f := &fizzbuzz.FizzBuzz{}
	err := json.Unmarshal([]byte(params), f)
	if err != nil {
		return nil, 0, err
	}
	return f, hits, nil
}

//UpsertQuery insert or update params into database.
//it increments the number of hits
func (c *Repository) UpsertQuery(params string) error {
	hits := 1
	q := `INSERT INTO stats (params, hits) VALUES (?, ?)
		ON CONFLICT(params) DO UPDATE SET
		hits=hits+1`
	statement, err := c.db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = statement.Exec(params, hits)
	if err != nil {
		return err
	}
	return nil
}
