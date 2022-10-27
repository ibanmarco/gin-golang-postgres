package main

import (
	"context"
	"github.com/gin-gonic/gin"
	handlers "github.com/ibanmarco/gin-golang-postgres/handlers"
	"github.com/ibanmarco/gin-golang-postgres/initializers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	initializers.PostgresConnection()
}

func main() {
	router := gin.Default()
	
	srv := &http.Server{
		Addr:         ":12345",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router.GET("/", handlers.RootHandler)
	router.GET("/welcome/:name", handlers.WelcomeHandler)
	router.GET("/books", handlers.ListBooksHandler)
	router.GET("/books/:id", handlers.GetBookHandler)
	router.POST("/books", handlers.CreateBookHandler)
	router.PUT("/books/:id", handlers.UpdateBookHandler)
	router.DELETE("/books/:id", handlers.DeleteBookHandler)

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	sig := <-signalChannel
	log.Println("Signal killed, graceful shutdown.", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(timeoutContext)

}
