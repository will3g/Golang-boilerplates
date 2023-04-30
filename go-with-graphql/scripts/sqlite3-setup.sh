#!/bin/bash

# Set the filename for the SQLite database
DB_FILE="data.db"

# Create the SQLite database file if it doesn't exist
if [ ! -f "$DB_FILE" ]; then
  touch "$DB_FILE"
fi

# Define the SQL statements for creating the tables
AUTHORS_TABLE_SQL="create table if not exists authors (id string, name string, description string, email string primary key);"

ARTICLES_TABLE_SQL="create table if not exists articles (id string primary key, title string, subtitle string, body string, author_id string, foreign key (author_id) references authors(id));"

# Execute the SQL statements to create the tables
sqlite3 "$DB_FILE" "$AUTHORS_TABLE_SQL"
sqlite3 "$DB_FILE" "$ARTICLES_TABLE_SQL"

# Display a message indicating that the tables were created
echo "Tables created successfully in $DB_FILE"
