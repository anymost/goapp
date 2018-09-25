package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/sync/errgroup"
	"time"
	"log"
)

var (
	g errgroup.Group
)

func router1() http.Handler {
	r := gin.Default()
	r.GET("/router1", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "router1",
		})
	})
	return r
}

func router2() http.Handler {
	r := gin.Default()
	r.GET("/router2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "router2",
		})
	})
	return r
}



func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
