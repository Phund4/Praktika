package main

import (
	"log"
	"os"
	"os/signal"
	"context"
	"time"

	"github.com/Phund4/Praktika/internal"
	"github.com/Phund4/Praktika/api"
)

func main() {
	err := internal.LoadEnv()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db, err := internal.NewDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.CloseDB()

	cch := internal.NewCache(db)

	err = db.InsertVacancies()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.RemoveVacancies()

	srv := api.NewAPI(cch);
	srv.StartServer();

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = srv.FinishServer(ctx);
	if err != nil {
		log.Fatal(err);
	}
	log.Println("Server successfully closed")
}
