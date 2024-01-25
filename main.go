package main

import (
	"Contract-Service/app/configs"
	"Contract-Service/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
	"log"
	"net/http"
	"os"
)

func main() {
	// Connexion à la base de données
	configs.ConnectDB()
	// Créez un routeur Gin
	router := gin.Default()
	// Routes de l'application
	routes.ContractRoute(router)
	// Lancez votre application Go sur un port spécifié
	goPort := os.Getenv("GO_DOCKER_PORT")
	if goPort == "" {
		goPort = "5000" // Port par défaut si aucun port n'est spécifié
	}
	// Enregistrez l'instance auprès de Eureka
	eurekaClient := fargo.NewConn("http://localhost:8761/eureka/v2")
	instance := &fargo.Instance{
		App:      "contract-service", // Remplacez par le nom de votre application
		Port:     5000,               // Port sur lequel votre service écoute
		HostName: "localhost",        // Nom d'hôte de votre service
		Status:   fargo.UP,
	}
	if err := eurekaClient.RegisterInstance(instance); err != nil {
		log.Fatalf("Failed to register with Eureka: %v", err)
	}

	// Lancez votre serveur HTTP
	err := http.ListenAndServe(":"+goPort, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
