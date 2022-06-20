package entities

import "github.com/ramiro/API/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close() // vai ser chamada depois da funcao Insert

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id` //comando sql para inserir os dados  RETURNING para retornar o id
	
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
