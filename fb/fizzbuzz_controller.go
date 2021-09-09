package fb

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FizzbuzzRequest struct {
	Int1  int    `form:"int1" json:"int1"`
	Int2  int    `form:"int2" json:"int2"`
	Limit int    `form:"limit" json:"limit"`
	Str1  string `form:"str1" json:"str1"`
	Str2  string `form:"str2" json:"str2"`
}

func GetFizzBuzz(c *gin.Context) {
	var fbRequest FizzbuzzRequest
	log.Println("1")
	if c.Bind(&fbRequest) != nil {
		log.Println("error reading request")
		c.IndentedJSON(http.StatusBadRequest, "")
		return
	}
	list, err := fizzbuzzList(fbRequest)
	log.Println(list)
	if err != nil {
		return
	}
	log.Println(list)
	c.IndentedJSON(http.StatusOK, list)
}
