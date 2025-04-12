package main

import (
	"log"
	"net/http"

	"github.com/fazendapro/FazendaPro-api/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar config:", err)
	}

	server := http.NewServeMux()

	log.Printf("Servidor rodando na porta %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, server))
}
