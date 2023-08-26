package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rafaelps1/imersao14/go/internal/db"
)

func main() {
	logger := httplog.NewLogger("route", httplog.Options{
		JSON: true,
	})

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "mysql", "3306", "routes_db")
	dbCon, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Panic().Msg(dataSourceName)
	}
	defer dbCon.Close()

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi!"))
	})

	r.Post("/routes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	r.Get("/routes", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Routes list"))
		dt := db.New(dbCon)
		ctx := context.Background()
		routes, err := dt.GetRoutes(ctx)
		if err != nil {
			panic("fail")
		}
		json.NewEncoder(w).Encode(routes)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
