module github.com/is0405/docker

go 1.16

// replace (
// 	github.com/is0405/docker/db => ./db
// 	github.com/is0405/docker/controller => ./controller
// 	github.com/is0405/docker/repository => ./repository
// 	github.com/is0405/docker/service => ./service
// )

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/justinas/alice v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.7.0
)
