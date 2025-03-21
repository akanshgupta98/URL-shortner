package handlers

import (
	"errors"
	"log"
	"net/http"
	"url_shortner/internal/service"

	"github.com/gin-gonic/gin"
)

type URLShortnerRequestData struct {
	InputURL string `json:"url"`
}

type URLShortnerResponseData struct {
	InputURL  string
	OutputURL string
	ErrMsg    string
}

// Fetch data from service layer, and return data.
func URLShortner(c *gin.Context) {
	var input URLShortnerRequestData
	var response URLShortnerResponseData
	log.Println("Repsonse is : ", response)
	if err := c.BindJSON(&input); err != nil {
		log.Println("Bad request: ", err.Error())
		response.ErrMsg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.InputURL = input.InputURL
	if input.InputURL == "" {
		log.Println("Bad request. URL empty value not allowed")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := service.URLShortner(input.InputURL)
	if err != nil {
		response.ErrMsg = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	log.Println("Complete URL is: ", c.Request.URL.Scheme, c.Request.Host, c.Request.URL)
	response.OutputURL = "http://" + c.Request.Host + c.Request.URL.String() + "/" + data
	c.JSON(http.StatusOK, response)
}

func URLShortnerFetch(c *gin.Context) {

	var response URLShortnerResponseData
	inputParam := c.Param("site")
	originalURL, err := service.URLShortnerFetch(inputParam)
	if err != nil {
		if errors.Is(err, service.ErrFailure) {
			response.ErrMsg = err.Error()
			c.JSON(http.StatusInternalServerError, response)
		} else if errors.Is(err, service.ErrInvalidRequest) {
			response.ErrMsg = err.Error()
			c.JSON(http.StatusBadRequest, response)
		}
		log.Printf("Error reason is: %s", err.Error())
		return
	}
	log.Println("original URL IS: ", originalURL)
	response.OutputURL = originalURL
	c.Redirect(http.StatusFound, response.OutputURL)
	// c.JSON(http.StatusOK, response)

}
