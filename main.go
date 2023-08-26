package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelps1/imersao14/go/internal/db"
)

func main() {
	logger := httplog.NewLogger("route", httplog.Options{
		JSON: true,
	})

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "db", "3306", "routes_db")
	dbCon, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Panic().Msg(err.Error())
	}
	defer dbCon.Close()

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi!"))
	})

	r.Post("/routes", func(w http.ResponseWriter, r *http.Request) {
		dt := db.New(dbCon)
		ctx := r.Context()
		var route db.CreateRoutesParams
		json.NewDecoder(r.Body).Decode(&route)

		err := dt.CreateRoutes(ctx, route)
		if err != nil {
			logger.Panic().Msg(err.Error())
		}
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
	})

	r.Get("/routes", func(w http.ResponseWriter, r *http.Request) {
		dt := db.New(dbCon)
		ctx := r.Context()
		routes, err := dt.GetRoutes(ctx)
		if err != nil {
			logger.Panic().Msg(err.Error())
		}
		respondwithJSON(w, http.StatusOK, routes)
	})

	http.ListenAndServe(":8080", r)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
