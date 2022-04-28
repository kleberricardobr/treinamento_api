package models

type ToDo struct {
	Descricao string `db:"descricao" json:"descricao"`
	DataHora  string `db:"datahora" json:"datahora"`
}

type ToDoList struct {
	Cab ToDo     `json:"to_do"`
	It  []ToDoIt `json:"it_to_do"`
}

type AllToDoList struct {
	All []ToDoList `json:"lista_to_do"`
}

type ToDoId struct {
	Id        int    `db:"id_do" json:"id_do"`
	Descricao string `db:"descricao" json:"descricao"`
	DataHora  string `db:"datahora" json:"datahora"`
}
