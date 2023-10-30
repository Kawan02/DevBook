package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

func mensagem(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Criando usuário!"))
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Criando usuário!"))
	mensagem(w, "Buscando todos os usuários!")
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Criando usuário!"))
	mensagem(w, "Buscando um usuário!")
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Criando usuário!"))
	mensagem(w, "Atualizando o usuário!")
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Criando usuário!"))
	mensagem(w, "Deletando um usuário!")
}
