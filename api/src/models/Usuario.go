package models

import (
	"errors"
	"strings"
	"time"
)

// omitempty => se o campo ID tiver em branco, ele não vai passar para o JSON. Ele não vai passar o ID 0, ele vai tirar o ID do JSON

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("Nome é obrigátorio!")
	}
	if usuario.Nick == "" {
		return errors.New("Nick é obrigátorio!")
	}
	if usuario.Email == "" {
		return errors.New("E-mail é obrigátorio!")
	}
	if usuario.Senha == "" {
		return errors.New("Senha é obrigátoria!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
}
