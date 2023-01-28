package controllers

import (
	"miltonjacomini/alura-loja/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.BuscaTodosProdutos())
}
