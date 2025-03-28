package database

import (
	"database/sql"
	"fmt"

	t "github.com/s4bb4t/srest-api/internal/lib/todoConfig"
)

func (s *Storage) Create(t t.TodoRequest) (int64, error) {
	const op = "database.postgres.CreateTodo"

	query := `
		INSERT INTO public.todos (title, is_done, description, executor)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", op, err)
	}
	defer stmt.Close()

	var id int64
	if t.IsDone != nil {
		err = stmt.QueryRow(t.Title, *t.IsDone, t.Description, t.Executor).Scan(&id)
	} else {
		err = stmt.QueryRow(t.Title, false, t.Description, t.Executor).Scan(&id)
	}
	if err != nil {
		return 0, fmt.Errorf("%s: %v", op, err)
	}

	return id, nil
}

func (s *Storage) Update(id int, t t.TodoRequest) (int64, error) {
	const op = "database.postgres.UpdateTodo"

	var stmt *sql.Stmt
	var err error
	var res sql.Result

	if t.IsDone != nil {
		stmt, err = s.db.Prepare(`UPDATE public.todos SET is_done = $1 WHERE id = $2`)
		if err != nil {
			return -1, fmt.Errorf("%s: prepare %v", op, err)
		}
		res, err = stmt.Exec(*t.IsDone, id)
		if err != nil {
			return -1, fmt.Errorf("%s:  exec %v", op, err)
		}

		stmt.Close()
	}

	if t.Title != "" {
		stmt, err = s.db.Prepare(`UPDATE public.todos SET title = $1  WHERE id = $2`)
		if err != nil {
			return -1, fmt.Errorf("%s: %v", op, err)
		}
		res, err = stmt.Exec(t.Title, id)
		if err != nil {
			return -1, fmt.Errorf("%s: %v", op, err)
		}
		stmt.Close()
	}

	if t.Description != "" {
		stmt, err = s.db.Prepare(`UPDATE public.todos SET description = $1 WHERE id = $2`)
		if err != nil {
			return -1, fmt.Errorf("%s: %v", op, err)
		}
		res, err = stmt.Exec(t.Description, id)
		if err != nil {
			return -1, fmt.Errorf("%s: %v", op, err)
		}
		stmt.Close()
	}

	if t.Executor != "" {
		stmt, err = s.db.Prepare(`UPDATE public.todos SET executor = $1 WHERE id = $2`)
		if err != nil {
			return -1, fmt.Errorf("%s: %v", op, err)
		}
		res, err = stmt.Exec(t.Executor, id)
		if err != nil {
			return -1, fmt.Errorf("%s: %v", op, err)
		}
		stmt.Close()
	}

	n, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}

	if n == 0 {
		return n, fmt.Errorf("%s: no task with id: %v", op, id)
	}

	return n, nil
}

func (s *Storage) Delete(id int) (int64, error) {
	const op = "database.postgres.DeleteTodo"

	stmt, err := s.db.Prepare(`
	DELETE FROM public.todos 
		WHERE id = $1
	`)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}

	if n == 0 {
		return n, fmt.Errorf("%s: no task with id: %v", op, id)
	}

	return n, nil
}

func (s *Storage) GetTodo(id int) (t.Todo, error) {
	const op = "database.postgres.GetTodo"

	rows, err := s.db.Query(`SELECT * FROM public.todos WHERE id = $1`, id)
	if err != nil {
		return t.Todo{}, fmt.Errorf("%s: %v", op, err)
	}
	defer rows.Close()

	var todo t.Todo

	if rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Created, &todo.IsDone, &todo.Description, &todo.Executor); err != nil {
			return t.Todo{}, fmt.Errorf("%s: %v", op, err)
		}
	} else {
		return t.Todo{}, fmt.Errorf("%s: no such task", op)
	}

	return todo, nil
}

func (s *Storage) OutputAll(filter string) ([]t.Todo, t.TodoInfo, int, error) {
	const op = "database.postgres.OutputAllTodos"

	query := ``
	switch filter {
	case "all":
		query = `SELECT * FROM public.todos ORDER BY id ASC`
	case "completed":
		query = `SELECT * FROM public.todos WHERE is_done = true ORDER BY id ASC`
	case "inWork":
		query = `SELECT * FROM public.todos WHERE is_done = false ORDER BY id ASC`
	default:
		query = `SELECT * FROM public.todos ORDER BY id ASC`
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, t.TodoInfo{}, 0, fmt.Errorf("%s: %v", op, err)
	}
	defer rows.Close()

	var result []t.Todo
	var todo t.Todo

	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Created, &todo.IsDone, &todo.Description, &todo.Executor); err != nil {
			return nil, t.TodoInfo{}, 0, fmt.Errorf("%s: %v", op, err)
		}

		result = append(result, todo)
	}

	var info t.TodoInfo

	query = `SELECT is_done FROM public.todos`

	rows, err = s.db.Query(query)
	if err != nil {
		return nil, t.TodoInfo{}, 0, fmt.Errorf("%s: %v", op, err)
	}

	var done bool
	for rows.Next() {
		if err := rows.Scan(&done); err != nil {
			return nil, t.TodoInfo{}, 0, fmt.Errorf("%s: %v", op, err)
		}

		if done {
			info.Completed++
		} else {
			info.InWork++
		}
		info.All++
	}

	return result, info, info.All, nil
}
