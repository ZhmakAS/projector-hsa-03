package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brianvoe/gofakeit/v6"
	metrics "github.com/tevjef/go-runtime-metrics"
	_ "github.com/tevjef/go-runtime-metrics/expvar"

	"projector-test-app/pkg"
	"projector-test-app/pkg/http"
)

func main() {
	var cfg Env
	if err := cfg.Parse(); err != nil {
		panic(err)
	}

	faker := gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		osCall := <-c
		log.Printf("Stop system call:%+v", osCall)
		cancel()
	}()

	mongoDB, err := pkg.InitMongoDB(ctx, cfg.MongodbURL)
	if err != nil {
		log.Println("Failed to connect to MongoDB")
		panic(err)
	}

	elasticClient, err := pkg.InitElasticSearchClient(ctx, cfg.ElasticSearchURL)
	if err != nil {
		log.Println("Cant connect to MongoDB")
		panic(err)
	}

	// running metrics collector for server
	err = metrics.RunCollector(&metrics.Config{
		Host:     cfg.InfluxURL,
		Database: cfg.InfluxDB,
		Username: cfg.InfluxDBUsername,
		Password: cfg.InfluxDBPassword,
	})
	if err != nil {
		panic(err)
	}

	router := NewRouter(mongoDB, elasticClient)
	server := http.NewServer(cfg.Addr, router)

	log.Printf("Starting up the API server...")
	if err := server.ListenAndServe(ctx); err != nil {
		log.Printf("Failed to serve api %s", err.Error())
	}
}
