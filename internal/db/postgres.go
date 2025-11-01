package db

import (
	"database/sql"
	"game-assistant-server/internal/model"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func Connect(databaseUrl string) *sql.DB {
	var db *sql.DB
	var err error
	retryNumber := 10
	retrySleep := time.Duration(1) * time.Second
	for i := range retryNumber {
		db, err = sql.Open("postgres", databaseUrl)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Print("Successfully connected to the database")
				return db
			}
		}
		log.Printf("Attempt %d: Database not ready, retrying in %s...", i, retrySleep)
		time.Sleep(retrySleep)
	}
	log.Fatalf("Could not connect to database after %d attempts, last error: %v", retryNumber, err)
	return nil
}

func GetAllCards(db *sql.DB) ([]model.Card, error) {
	rows, err := db.Query("SELECT id, title, description FROM Cards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []model.Card
	for rows.Next() {
		var c model.Card
		if err := rows.Scan(&c.ID, &c.Title, &c.Description); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}

func GetCard(db *sql.DB, id int) (*model.Card, error) {
	card := &model.Card{ID: id}
	err := db.QueryRow("SELECT title, description FROM Cards WHERE id = $1",
		id).Scan(&card.Title, &card.Description)
	if err != nil {
		return nil, err
	}
	card.ID = id
	return card, nil
}

func InsertCard(db *sql.DB, card *model.Card) (int, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO Cards (title, description) VALUES ($1, $2) RETURNING id",
		card.Title, card.Description).Scan(&id)
	return id, err
}
