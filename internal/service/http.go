package service

import (
	"net/http"
	"time"
)

// getNewClient - cria um novo client http com timeout de 30 segundos
func getNewClient() http.Client {
	var client http.Client
	client.Timeout = 30 * time.Second

	return client
}
