package rest

import (
	"fmt"
	"github.com/Nariett/arox-pkg/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func StartHttpServer(cfg *config.Config, r chi.Router) {
	fmt.Printf("Starting server at port :%s\n", cfg.Port)
	port := fmt.Sprintf(":%s", cfg.Port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
