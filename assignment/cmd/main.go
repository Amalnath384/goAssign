package main

import (
	"project/assignment/pkg/model"
	db "project/assignment/pkg/repository"
	."project/assignment/router"
)

func main() {
	dbHost := "127.0.0.1:27017"
	db.Init(&model.Database{
		Driver:   "mongodb",
		Endpoint: dbHost})
	defer db.Exit()

	Router()

}
