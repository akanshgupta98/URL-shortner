package server

import (
	"log"
	"url_shortner/internal/config"
	"url_shortner/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Run(cfg config.ServerConfig) (err error) {
	router := gin.Default()
	router.GET("/")
	router.GET("/api/url-shortner/:site", handlers.URLShortnerFetch)
	router.POST("/api/url-shortner", handlers.URLShortner)
	err = router.Run(cfg.IP + ":" + cfg.Port)
	if err != nil {
		log.Println("Unable to start the server")
		return err
	}
	return

}
