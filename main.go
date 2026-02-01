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
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// debug log untuk memastikan env terbaca
	log.Printf("Loaded config: PORT=%s, DB_CONN=%s", cfg.Port, cfg.DBConn)

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

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"Server running on port %s"}`, cfg.Port)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)
	http.HandleFunc("/products", productHandler.HandleProducts)
	http.HandleFunc("/products/", productHandler.HandleProductByID)

	addr := "0.0.0.0:" + cfg.Port
	log.Println("Server running at", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
