package dao

import (
	"database/sql"
	"fmt"
	"treinamento-api/database"
	"treinamento-api/models"
)

func GetToDoByID(pID int) (ret models.ToDoList, err error) {
	var toDoCab models.ToDo
	var toDoIt models.ToDoIt

	err = database.DB.Get(&toDoCab, `select descricao, 
	                                        datahora 
								     from to_do_cab
								     where id_do = $1`, pID)

	if err != nil {
		return
	}

	ret.Cab = toDoCab

	rows, err := database.DB.Queryx(`select descricao,
	                                        prioridade
				  				  	 from to_do_it
									 where id_do = $1
									 order by prioridade`, pID)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&toDoIt.Descricao, &toDoIt.Prioridade)
		ret.It = append(ret.It, toDoIt)
	}

	return
}

func GetAllToDo(pQtd int) (ret models.AllToDoList, err error) {
	limit := ""
	if pQtd > 0 {
		limit = fmt.Sprintf(`limit %d`, pQtd)
	}

	rows, err := database.DB.Queryx(fmt.Sprintf(`select id_do
				 	 			                 from to_do_cab
							                     order by id_do desc
							   			         %s`, limit))

	if err != nil {
		return
	}

	var idDo int
	var it models.ToDoList
	for rows.Next() {
		rows.Scan(&idDo)

		it, err = GetToDoByID(idDo)
		if (err != nil) && (err != sql.ErrNoRows) {
			return
		}

		ret.All = append(ret.All, it)
	}

	return

}

func CadToDo(toDo models.ToDoList) (err error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return
	}

	var idToDo int
	err = tx.QueryRow(`insert into to_do_cab(id_do, descricao, datahora)
	                   values(nextval('sequence_to_do_id'), $1, $2)
					   returning id_do`, toDo.Cab.Descricao, toDo.Cab.DataHora).Scan(&idToDo)
	if err != nil {
		tx.Rollback()
		return
	}

	for _, it := range toDo.It {
		_, err = tx.Exec(`insert into to_do_it(descricao, prioridade, id_do)
		                  values($1, $2, $3)`, it.Descricao, it.Prioridade, idToDo)

		if err != nil {
			tx.Rollback()
			return
		}

	}

	tx.Commit()

	return nil

}
