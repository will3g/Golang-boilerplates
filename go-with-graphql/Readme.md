# Golang with GraphQL

This is a simple **GraphQL API** boilerplate built using **gqlgen** and **Golang**. It uses **sqlite3** as the database engine and the uuid library for generating unique IDs.

# Demo

![Golang with GraphQL demonstration](https://i.imgur.com/uG3tyfG.gif)

# Getting Started
To get started with this API, clone the repository and run the following commands:

```sh
# Install dependencies
go mod download

# Generate GraphQL boilerplate code
make graphql-setup

# Set up the sqlite3 database
make sqlite3-setup

# Start the API server
make compose-up
```

**This should start the API server and make it available at `http://localhost:8080/graphql`**.

# API Endpoints
Here are the **available endpoints for the API**:

- **/graphql**: This is the main GraphQL endpoint that allows you to query and mutate data using GraphQL.

# Example Queries
Here are some example GraphQL queries that you can use to query the API:

```graphql
# Query all authors
query {
  authors {
    id
    name
    description
  }
}
```

```graphql
# Query all articles and their authors
query {
  articles {
    id
    title
    subtitle
    body
    authorId {
      id
      name
      description
      email
    }
  }
}
```

```graphql
# Create a new author
mutation {
  createAuthor(input: {
    name: "John Doe",
    description: "A new author",
    email: "john.doe@example.com"
  }) {
    id
    name
    description
    email
  }
}
```

```graphql
# Create a new article
mutation {
  createArticle(input: {
    title: "My Article",
    subtitle: "A subtitle",
    body: "Some article content",
    authorId: "f4b4e83c-f964-11eb-ba80-0242ac130004"
  }) {
    id
    title
    subtitle
    body
    authorId {
      id
      name
      description
      email
    }
  }
}
```

```graphql
# Query authors and their articles
query {
  authors {
    id
    name
    description
    email
    articles {
      id
      title
      subtitle
      body
    }
  }
}
```

```graphql
# Query articles and their authors
query {
  articles {
    id
    title
    subtitle
    body
    authorId {
      id
      name
      description
      email
    }
  }
}
```

# Acknowledgments
This project was built using **gqlgen**, **sqlite3**, and **uuid**.
