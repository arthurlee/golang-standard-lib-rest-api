package main

import (
	"fmt"
	"./utils/database"
	"./utils/caching"
	"./controllers"
	"./routes"
	"os"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("onnect to db ok")

	cache := &caching.Redis{
		Client:caching.Connect(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), 0),
	}
	fmt.Println("onnect to redis ok")

	mux := http.NewServeMux()

	userController := controllers.NewUserController(db, cache)
	routes.CreateRoutes(mux, userController)

	addr := ":8000"
	fmt.Printf("Server run on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
