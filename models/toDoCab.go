package models

type ToDoCab struct {
	Descricao string `db:"descricao" json:"descricao"`
	DataHora  string `db:"data_hora" json:"data_hora"`
}
