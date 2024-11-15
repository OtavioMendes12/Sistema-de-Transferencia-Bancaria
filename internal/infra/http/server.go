package http

import (
	"log"
	"net/http"
)

// StartServer inicia o servidor HTTP
func StartServer(router http.Handler, port string) {
	log.Printf("Servidor rodando na porta %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}