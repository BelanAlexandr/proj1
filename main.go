package main

import (
	"proj1/repository"
	rout "proj1/router"

	_ "github.com/lib/pq"
)

func main() {
	repository.Loading()
	rout.Rout()
}
