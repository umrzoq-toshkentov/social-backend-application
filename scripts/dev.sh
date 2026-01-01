#!/bin/bash

# Load environment variables from .envrc
export ADDR=":8080"
export DB_ADDR="postgres://umrzoqtoshkentov:@localhost:5432/social?sslmode=disable"

# Run air
go tool air
