package main

import (
	"fmt"
	"net/http"

	"github.com/nazordz/boutique/product-catalog-service/configs"
	"github.com/nazordz/boutique/product-catalog-service/models"
	"github.com/nazordz/boutique/product-catalog-service/routers"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	configs.Setup()
	models.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", configs.EnvConfigs.Port)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
