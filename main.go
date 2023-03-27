package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/gousb"
)

var counter = 0

func main() {
	fmt.Println(gousb.ClassHID)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9999")
	var wg sync.WaitGroup
	potok, _ := strconv.Atoi(os.Args[1])
	fmt.Println("potok ", potok)
	for i := 0; i < potok; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				counter++
			}
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("counter", counter)
}
