package models

type ToDoIt struct {
	Descricao  string `db:"descricao" json:"descricao"`
	Prioridade string `db:"prioridade" json:"prioridade"`
}
