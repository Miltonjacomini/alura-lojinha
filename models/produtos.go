package models

import (
	"miltonjacomini/alura-loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {

	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletaProdutoPorId, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletaProdutoPorId.Exec(id)
	defer db.Close()
}

func BuscaProdutoPorId(id string) Produto {
	db := db.ConectaComBancoDeDados()

	buscaProdutoPorId, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoFinal := Produto{}
	for buscaProdutoPorId.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = buscaProdutoPorId.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoFinal.Id = id
		produtoFinal.Nome = nome
		produtoFinal.Descricao = descricao
		produtoFinal.Preco = preco
		produtoFinal.Quantidade = quantidade
	}
	defer db.Close()
	return produtoFinal
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	produtoParaAtualizar := Produto{}

	produtoParaAtualizar.Id = id
	produtoParaAtualizar.Nome = nome
	produtoParaAtualizar.Descricao = descricao
	produtoParaAtualizar.Preco = preco
	produtoParaAtualizar.Quantidade = quantidade

	db := db.ConectaComBancoDeDados()

	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
