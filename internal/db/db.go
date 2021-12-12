package db

import (
	"fmt"
	"log"
	"os"

	"example.com/microservice/rocket-service/internal/rocket"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type Store struct {
	db *sqlx.DB
}

// New - return a new storage connection.
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}

	return Store{
		db: db,
	}, nil
}

//GetRocketByID - retrive a rocket from database by id
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRow(
		"SELECT id, name, type FROM rockets where id=$1",
	)

	err := row.Scan(rkt.ID, rkt.Name, rkt.Type)
	if err != nil {
		log.Print(err.Error())
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

//InsertRocket - inserts a rocket into database.
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO rockets
		(id, name, type)
		VALUES(:id, :name, :type)`,
		rkt,
	)
	if err != nil {
		log.Print(err.Error())
		return rocket.Rocket{}, err
	}
	return rocket.Rocket{
		ID:   rkt.ID,
		Name: rkt.Name,
		Type: rkt.Type,
	}, nil
}

//DeleteRocket - deletes a rocket from databse
// if fails return a error
func (s Store) DeleteRocket(id string) error {
	uid, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(
		`DELETE FROM rockets where id = $1`,
		uid,
	)

	if err != nil {
		log.Print("Deletion failed.")
		return err
	}

	return nil
}
