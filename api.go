package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/fib/:n", fibHandler)

	router.Run()
}

func fibHandler(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		log.Fatal("Erro na conversão Atoi")
	}
	c.JSON(http.StatusOK, fmt.Sprintf("O resultado é %v", fib(uint(n))))
	//c.String(http.StatusOK)
}

// 0 1 1 2 3 5 8 13
// 0 1 2 3 4 5 6 7

func fib(n uint) uint {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
