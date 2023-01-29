package controllers

import (
	"log"
	"miltonjacomini/alura-loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.BuscaTodosProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvetido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preço: ", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade: ", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvetido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	temp.ExecuteTemplate(w, "Edit", models.BuscaProdutoPorId(idDoProduto))
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Não foi possível converter ID, ", err)
		}
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Não foi possível converter a quantidade, ", err)
		}
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Não foi possível converter o Preço, ", err)
		}

		models.AtualizaProduto(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}
