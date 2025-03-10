package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ReynoldArun09/blog-application-golang/database"
	"github.com/ReynoldArun09/blog-application-golang/routes"
	"github.com/ReynoldArun09/blog-application-golang/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Warning no env file found")
	}

	db := database.InitDB()

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, db)

	port := utils.GetEnvVariables("PORT")

	fmt.Println("Server is up and running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}
