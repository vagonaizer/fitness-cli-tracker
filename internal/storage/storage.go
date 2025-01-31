package storage

import (
	"database/sql"
	"fitness-cli-tracker/internal/models"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS records (
			date DATE PRIMARY KEY,
			weight FLOAT,
			trained BOOLEAN,
			calories INTEGER
		)	
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveRecord(record models.Record) error {
	_, err := s.db.Exec(`
		INSERT INTO records (date, weight, trained, calories)
		VALUES($1, $2, $3, $4)
		`, record.Date, record.Weight, record.Trained, record.Calories)

	return err
}

func (s *Storage) GetRecords() ([]models.Record, error) {
	rows, err := s.db.Query("SELECT date, weight, trained, calories FROM records ORDER BY date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []models.Record
	for rows.Next() {
		var r models.Record
		err := rows.Scan(&r.Date, &r.Weight, &r.Trained, &r.Calories)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}
