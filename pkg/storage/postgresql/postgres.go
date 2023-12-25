package postgresql

import (
	"30.80.1/pkg/storage/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func Init(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

func (s *Storage) NewTask(task model.Task) (int, error) {
	var id int
	query := `
		INSERT INTO tasks (opened, author_id, assigned_id, title, content)
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`
	err := s.db.QueryRow(context.Background(), query, task.Opened, task.AuthorID, task.AssignedID, task.Title, task.Content).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("failed to insert task: %w", err)
	}
	return id, nil
}

func (s *Storage) Tasks() (tasks []model.Task, err error) {
	rows, err := s.db.Query(context.Background(),
		`SELECT * FROM tasks ORDER BY id;`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var t model.Task
		if err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		); err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}

		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

func (s *Storage) TaskById(id int) ([]model.Task, error) {
	exists, err := checkExistsId(s, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("task with ID %d does not exist", id)
	}

	rows, err := s.db.Query(context.Background(),
		`SELECT * FROM tasks WHERE id = $1;`, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task by id: %w", err)
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		); err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) TasksByAuthor(id int) ([]model.Task, error) {
	exists, err := checkExistsId(s, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("author with ID %d does not exist", id)
	}

	rows, err := s.db.Query(context.Background(),
		`SELECT * FROM tasks WHERE author_id = $1 ORDER BY id;`,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get task by author: %w", err)
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		); err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) UpdateTask(id int, t model.Task) error {
	exists, err := checkExistsId(s, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("task with ID %d does not exist", id)
	}

	_, err = s.db.Exec(context.Background(),
		`UPDATE tasks SET author_id = $1, assigned_id = $2, title = $3, content = $4 WHERE id = $5;`,
		t.AuthorID,
		t.AssignedID,
		t.Title,
		t.Content,
		id,
	)
	return err
}

func (s *Storage) DeleteTask(id int) error {
	exists, err := checkExistsId(s, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("task with ID %d does not exist", id)
	}

	_, err = s.db.Exec(context.Background(),
		`DELETE FROM tasks WHERE id = $1;`,
		id,
	)
	return err
}

// проверка id в БД, перед выполнением операций
func checkExistsId(s *Storage, id int) (exists bool, err error) {
	err = s.db.QueryRow(context.Background(), `SELECT EXISTS(SELECT * FROM tasks WHERE id = $1)`, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists {
		return true, err
	}
	return false, nil
}
