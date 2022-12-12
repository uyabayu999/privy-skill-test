package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"privy-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllCakes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cakes, err := models.GetAllCakes()

	if err != nil {
		log.Fatalf("Gagal mengambil data. %v", err)
	}

	json.NewEncoder(w).Encode(cakes)
}

func GetCakeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Gagal konversi string id ke int")
	}

	cake, err := models.GetCakeById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			res := response{
				Message: "Data tidak ditemukan",
			}

			w.WriteHeader(404)
			json.NewEncoder(w).Encode(res)
		} else {
			log.Fatalf("Gagal mengambil data. %v", err)
		}
	} else {
		json.NewEncoder(w).Encode(cake)
	}
}

func CreateCake(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var cake models.Cake

	err := json.NewDecoder(r.Body).Decode(&cake)

	if err != nil {
		log.Printf("Gagal mendecode dari request body. %v", err)
	}

	var res response

	validate := validator.New()
	err = validate.Struct(cake)

	if err != nil {
		res = response{
			Message: "Request tidak valid",
		}

		w.WriteHeader(400)
	} else {
		insertID := models.CreateCake(cake)
		res = response{
			ID:      insertID,
			Message: "Data berhasil ditambahkan",
		}

		w.WriteHeader(201)
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateCakeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Gagal konversi string id ke int")
	}

	var cake models.Cake

	err = json.NewDecoder(r.Body).Decode(&cake)

	if err != nil {
		log.Printf("Gagal mendecode dari request body. %v", err)
	}

	var res response

	validate := validator.New()
	err = validate.Struct(cake)

	if err != nil {
		res = response{
			Message: "Request tidak valid",
		}

		w.WriteHeader(400)
	} else {
		updatedRows := models.UpdateCakeById(id, cake)

		if updatedRows < 1 {
			w.WriteHeader(404)
			res = response{
				Message: "Data tidak ditemukan",
			}
		} else {
			msg := fmt.Sprintf("Data berhasil diupdate. %v rows terupdate.", updatedRows)

			res = response{
				ID:      id,
				Message: msg,
			}
		}
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteCakeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Gagal konversi string id ke int")
	}

	deletedRows := models.DeleteCakeById(id)

	var res response

	if deletedRows < 1 {
		w.WriteHeader(404)
		res = response{
			Message: "Data tidak ditemukan",
		}
	} else {
		msg := fmt.Sprintf("Data berhasil dihapus. %v rows terhapus.", deletedRows)

		res = response{
			ID:      id,
			Message: msg,
		}
	}

	json.NewEncoder(w).Encode(res)
}
