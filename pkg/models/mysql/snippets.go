package mysql

import (
	"database/sql"

	"github.com/abhijitherekar/websvc_backend/pkg/models"
)

type SnippetModel struct {
	Db *sql.DB
}

func (s *SnippetModel) Insert(title, content, expTime string) (int, error) {

	sqlStmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := s.Db.Exec(sqlStmt, title, content, expTime)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := s.Db.QueryRow(stmt, id)
	temp := &models.Snippet{}
	err := row.Scan(&temp.ID, &temp.Title, &temp.Content, &temp.CreateTime, &temp.ExpTime)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

func (s *SnippetModel) Show() ([]*models.Snippet, error) {
	return nil, nil
}
