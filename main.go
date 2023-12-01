package main

import (
	"baseapp/baseproject"
	"baseapp/config"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

// TLS: https://pkg.go.dev/github.com/gin-gonic/gin#Engine.RunTLS
func main() {
	router, config := baseproject.SetupRouter()
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(router, config, &wg)
	go startTLSServer(router, config, &wg)
	wg.Wait()
}

// https://apple.stackexchange.com/questions/393715/do-you-want-the-application-main-to-accept-incoming-network-connections-pop
func startServer(router *gin.Engine, config *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	port := config.HttpPortString()
	err := router.Run(port)
	if err != nil {
		log.Fatalf("Http server crashed: " + err.Error())
	}
}

func startTLSServer(router *gin.Engine, config *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	port := config.HttpsPortString()
	if config.CrtFile != "" && config.KeyFile != "" {
    	err := router.RunTLS(port, config.CrtFile, config.KeyFile)
        if err != nil {
            log.Fatalf("Https server crashed: " + err.Error())
        }
	}
}
