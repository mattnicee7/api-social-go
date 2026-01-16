package controllers

import (
	"api-social-go/src/banco"
	"api-social-go/src/modelos"
	"api-social-go/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario cria um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
}

// BuscarUsuarios retorna todos usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios"))
}

// BuscarUsuario retorna um usuario especifico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

// AtualizarUsuario atualiza um usuario especifico
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

// DeletarUsuario deleta um usuario especifico
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
