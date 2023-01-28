package main

import (
	"net/http"

	"miltonjacomini/alura-loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
