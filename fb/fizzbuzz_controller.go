package fb

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFizzBuzz(c *gin.Context) {
	var fbRequest FizzbuzzRequest
	log.Println("1")
	if c.Bind(&fbRequest) != nil {
		log.Println("error reading request")
		c.IndentedJSON(http.StatusBadRequest, "")
		return
	}
	log.Println(fbRequest)
	list, err := fizzbuzzList(fbRequest)
	log.Println(list)
	if err != nil {
		return
	}
	log.Println(list)
	c.IndentedJSON(http.StatusOK, list)
}

func GetMostFrequent(c *gin.Context) {
	response, err := mostFrequentRequest()
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}
