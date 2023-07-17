package database

import "github.com/procode2/accunotes/database/postgres"

var Db Storer = postgres.NewPostgresStore()
