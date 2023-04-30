package graph

import "github.com/will3g/golang-boilerplates/golang-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleDB *database.Article
	AuthorDB  *database.Author
}
