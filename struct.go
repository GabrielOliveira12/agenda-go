package main

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model `swaggerignore:"true"`
	Nome       string `json:"nome"`
	Sobrenome  string `json:"sobrenome"`
	Lembretes  []Lembrete
}

type Lembrete struct {
	gorm.Model `swaggerignore:"true"`
	Nome       string `json:"nome"`
	Descricao  string `json:"descricao"`
	Data       string `json:"data"`
	UsuarioID  uint   `json:"usuarioID"`
}
