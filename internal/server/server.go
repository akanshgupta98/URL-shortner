package server

import (
	"log"
	"url_shortner/internal/config"
	"url_shortner/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
	Addr   string
}

func Initialize(cfg config.ServerConfig) (r Router, err error) {
	router := gin.Default()
	router.GET("/")
	router.GET("/api/url-shortner/:site", handlers.URLShortnerFetch)
	router.POST("/api/url-shortner", handlers.URLShortner)
	r.Router = router
	r.Addr = cfg.IP + ":" + cfg.Port
	return

}

func Run(r Router) (err error) {
	err = r.Router.Run(r.Addr)
	if err != nil {
		log.Println("Unable to start the server")
		return
	}
	return
}
