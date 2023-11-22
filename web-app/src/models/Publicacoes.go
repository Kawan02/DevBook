package models

import "time"

type Publicacao struct {
	ID         uint64    `json:"id,omitempty"`
	Titulo     string    `json:"titulo,omitempty"`
	Conteudo   string    `json:"conteudo,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick uint64    `json:"authorNick,omitempty"`
	Curtidas   uint64    `json:"curtidas"`
	CriadaEm   time.Time `json:"criadaEm,omitempty"`
}
