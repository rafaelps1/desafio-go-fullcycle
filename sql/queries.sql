-- name: CreateRoutes :exec
INSERT INTO routes (id, name, source_lat, source_lng, destination_lat, destination_lng) VALUES(?,?,?,?,?,?);

-- name: GetRoutes :many
SELECT * FROM routes;
