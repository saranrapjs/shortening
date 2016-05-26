package links

import (
	"database/sql"
)

type linkService struct {
	db *sql.DB
}

func NewLinkService(db *sql.DB) *linkService {
	return &linkService{
		db: db,
	}
}

func (l linkService) Get(slug string) (*Link, error) {
	var li Link
	row := l.db.QueryRow("SELECT url, slug FROM links WHERE slug = ?", slug)
	if err := row.Scan(&li.URL, &li.Slug); err != nil {
		return nil, err
	}
	return &li, nil
}

func (l linkService) Update(slug string, url string) (*Link, error) {
	if _, err := l.db.Exec("INSERT INTO links (url, slug) VALUES(?, ?) ON DUPLICATE KEY UPDATE url=?", url, slug, url); err != nil {
		return nil, err
	}

	return l.Get(slug)
}

func (l linkService) Delete(slug string) error {
	if _, err := l.db.Exec("DELETE FROM links WHERE slug = ?", slug); err != nil {
		return err
	}
	return nil
}

func (l linkService) List() ([]*Link, error) {
	var lks []*Link
	rows, err := l.db.Query("SELECT url, slug FROM links LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var lk Link
		if err := rows.Scan(&lk.URL, &lk.Slug); err != nil {
			return nil, err
		}
		lks = append(lks, &lk)
	}

	return lks, nil
}
