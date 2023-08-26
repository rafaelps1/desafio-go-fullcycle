package repository

import (
	"database/sql"

	"github.com/rafaelps1/imersao14/go/internal/freight/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {
	// sqlSmt := "INSERT INTO routes (id, name, source_lat, source_lng, destination_lat, destination_lng) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	// _, err := r.db.Exec(sqlSmt, route.ID, route.Name, route.Source.Lat, route.Source.Lng, route.Destination.Lat, route.Destination.Lng)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (r *RouteRepositoryMysql) List() ([]*entity.Route, error) {
	// sqlSmt := "SELECT id, name, source_lat, source_lng, destination_lat, destination_lng FROM routes"
	// rows := r.db.Query(sqlSmt, nil)
	// var routes []*entity.Route
	// err := rows.Scan(
	// 	&route.ID,
	// )
	return nil, nil
}
