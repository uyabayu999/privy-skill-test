package models

import (
	"database/sql"
	"log"
	"privy-test/config"
)

type Cake struct {
	ID          int     `json:"id" validate:"isdefault"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float32 `json:"rating" validate:"required,numeric"`
	Image       string  `json:"image" validate:"required"`
	CreatedAt   string  `json:"created_at" validate:"isdefault"`
	UpdatedAt   string  `json:"updated_at" validate:"isdefault"`
}

func GetAllCakes() ([]Cake, error) {
	var cakes []Cake

	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM cakes ORDER BY rating DESC, title ASC`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cake Cake

		err = rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)

		if err != nil {
			log.Fatalf("Tidak bisa mengambil data. %v", err)
		}

		cakes = append(cakes, cake)
	}

	return cakes, err
}

func GetCakeById(id int) (Cake, error) {
	var cake Cake

	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM cakes WHERE id = ?`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)

	switch err {
	case sql.ErrNoRows:
		log.Printf("Data tidak ditemukan")
	case nil:
		return cake, nil
	default:
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	return cake, err
}

func CreateCake(cake Cake) int {
	db := config.CreateConnection()
	defer db.Close()

	var id int64

	sqlStatement := `INSERT INTO cakes (title, description, rating, image) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(sqlStatement, cake.Title, cake.Description, cake.Rating, cake.Image)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	} else {
		id, err = res.LastInsertId()

		if err != nil {
			panic(err)
		}
	}

	return int(id)
}

func UpdateCakeById(id int, cake Cake) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE cakes SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?`

	res, err := db.Exec(sqlStatement, cake.Title, cake.Description, cake.Rating, cake.Image, id)

	if err != nil {
		log.Fatalf("Tidal bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error ketika memeriksa data yang diupdate. %v", err)
	}

	return rowsAffected
}

func DeleteCakeById(id int) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM cakes WHERE id = ?`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Data tidak ditemukan. %v", err)
	}

	return rowsAffected
}
