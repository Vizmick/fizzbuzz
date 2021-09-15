package fb

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFizzBuzz(c *gin.Context) {
	var fbRequest FizzbuzzRequest
	log.SetPrefix("/fizzbuzz :")
	if c.Bind(&fbRequest) != nil {
		log.Println("error reading request")
		c.IndentedJSON(http.StatusBadRequest, "")
		return
	}
	log.Println(fbRequest)
	list, err := fizzbuzzList(fbRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	log.Println(list)
	c.IndentedJSON(http.StatusOK, list)
}

func GetMostFrequent(c *gin.Context) {
	log.SetPrefix("/frequent :")
	response, err := mostFrequentRequest()
	if err != nil {
		log.Println("error in accessing database")
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}
