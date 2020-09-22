package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"ordergo.com/api"
	"ordergo.com/database"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	//connect DB
	connect := database.NewDB("root:password@1@tcp(localhost:3306)/test")

	//lib
	// *gin.Engine router =null;
	router := gin.Default()

	// Middleware
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Cache-Control", "no-cache")
		// context.Writer.Header().Set("Content-Type", "application/json")
	})

	//api
	controller := new(api.API)

	controller.Data = connect
	api := router.Group("/api")

	api.GET("", func(context *gin.Context) {

		name := "max"
		dept := "sombutjalearn"

		_, err := fmt.Fprintf(os.Stdout, "%s is a %s portal.\n", name, dept)
		if err != nil {
			fmt.Print(err)
			return
		}

		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
			return
		}
		defer f.Close()

		// log.SetOutput(f)
		// log.Println("This is a test log entry")

		d1 := []byte("hello\ngo\n")

		f.Write(d1)

		context.JSON(http.StatusOK, gin.H{
			"name": "",
			"id":   1,
		})
	})

	api.POST("customer/createCustomer", controller.InsertCustomer)
	api.POST("customer/searchCustomer", controller.SearchCustomer)
	api.GET("test", func(context *gin.Context) {
		fmt.Println("NNNtestNN")
		context.JSON(http.StatusOK, gin.H{
			"name": "",
			"id":   1,
		})
	})
	api.GET("testspeed", func(context *gin.Context) {
		sum := 0
		for i := 0; i < 100000; i++ {
			sum += i
		}
		fmt.Println(sum)
	})

	// router.Run("localhost:9999")
	router.Run("10.11.2.23:8080")
}
