package main

import (
	"fmt"
	"log"
	"net/http"

	"app-kasir/config"
	"app-kasir/handlers"
	"app-kasir/repositories"
	"app-kasir/services"
)

func main() {
	// load config pakai viper
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// koneksi database
	db, err := config.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// dependency injection
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"Server running on port %s"}`, cfg.Port)
	})
	http.HandleFunc("/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)

	addr := "localhost:" + cfg.Port
	fmt.Println("Server running at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
