package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nariett/arox-pkg/config"
	"github.com/go-chi/chi/v5"
)

func StartHttpServer(cfg *config.Config, r chi.Router) {
	log.Printf("Starting server at port :%s\n", cfg.Port)
	port := fmt.Sprintf(":%s", cfg.Port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Println("Error starting the server:", err)
	}
}
