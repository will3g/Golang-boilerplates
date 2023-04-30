sqlite3-setup:
	./scripts/sqlite3-setup.sh

graphql-setup: graphqlgen gqlgen-init gqlgen-generate

graphqlgen:
	printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	go mod tidy

gqlgen-init:
	go run github.com/99designs/gqlgen init

gqlgen-generate:
	go run github.com/99designs/gqlgen generate

compose-build:
	docker-compose up --build -d

compose-up:
	docker-compose up