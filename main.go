package main

import (
	"fmt"
	"mu_previous_papers_be/model"
	"mu_previous_papers_be/server"
	"mu_previous_papers_be/store"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	dns := os.Getenv("DATABASE_URL")
	db, err := model.NewDB(dns)
	if err != nil {
		fmt.Println("error connecting to pgsql server: ", err)
		return
	}
	str := store.NewStore(db)
	serverMain := server.NewServer(str)
	serverMain.Run()

}
