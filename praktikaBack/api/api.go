package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/Phund4/Praktika/internal"
)

type API interface {
	StartServer()
	FinishServer(context.Context) error
}

type api struct {
	router *chi.Mux
	srv    *http.Server
	ctx    context.Context
	cch    internal.ICache
}

func NewAPI(cache internal.ICache) API {
	r := chi.NewRouter()
	cont := context.Background()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	api := &api{
		router: r,
		ctx:    cont,
		cch:    cache,
	}

	r.Get("/vacancies", api.getVacancies)

	return api
}

func (a *api) StartServer() {
	connStr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	log.Printf("Server must be start on %v", connStr)
	server := &http.Server{Addr: connStr, Handler: a.router}
	a.srv = server

	go func() {
		log.Println("Server successfully started")
		err := a.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error in start server: %v", err.Error())
		}
	}()
}

func (a *api) FinishServer(ctx context.Context) error {
	err := a.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error in closing server: %v", err)
	}
	return nil
}

func (a *api) getVacancies(w http.ResponseWriter, r *http.Request) {
	vacancies, err := a.cch.GetVacancies();
	if err != nil {
		log.Fatalf("error in getting vacancies: %v", err);
		w.WriteHeader(500)
		return;
	}

	bytes, err := json.Marshal(vacancies)
	if err != nil {
		log.Fatalf("error in marshalling data: %v", err);
		w.WriteHeader(500)
		return;
	}
	w.Write(bytes);
}
