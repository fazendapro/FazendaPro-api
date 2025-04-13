package main

import (
	"log"
	"net/http"

	"github.com/fazendapro/FazendaPro-api/api/handlers"
	"github.com/fazendapro/FazendaPro-api/config"
	"github.com/fazendapro/FazendaPro-api/internal/repository"
	"github.com/fazendapro/FazendaPro-api/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar configuração:", err)
	}

	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FazendaPro API rodando!"))
	})

	mux.HandleFunc("/user", userHandler.GetUser)

	log.Printf("Servidor rodando na porta %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
