package main

import (
	"fitness-cli-tracker/internal/plot"
	"fitness-cli-tracker/internal/storage"
	"fitness-cli-tracker/internal/ui"
	"fmt"
	"log"
)

func main() {
	// Строка подключения к PostgreSQL
	connStr := "user=postgres dbname=tracker sslmode=disable password=123 host=localhost port=5432"
	db, err := storage.NewStorage(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	record, err := ui.GetInput()
	if err != nil {
		log.Fatalf("Failed to get input: %v", err)
	}

	if err := db.SaveRecord(record); err != nil {
		log.Fatalf("Failed to save record: %v", err)
	}

	records, err := db.GetRecords()
	if err != nil {
		log.Fatalf("Failed to get records: %v", err)
	}

	if err := plot.PlotWeight(records); err != nil {
		log.Fatalf("Failed to plot weight: %v", err)
	}

	fmt.Println("Данные сохранены и график построен!")
}
