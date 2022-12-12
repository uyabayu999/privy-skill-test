package main

import (
	"fmt"
	"log"
	"net/http"
	"privy-test/routers"
)

func main() {
	r := routers.Router()

	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
