package routes

import (
	"miltonjacomini/alura-loja/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
