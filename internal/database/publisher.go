package database

import (
	"database/sql"
	"log"

	"github.com/jwDevOps/atlas-backend/internal/models"
)

func (m *DbManager) QueryPublishers() ([]*models.Publisher, error) {
	rows, err := m.Db.Query("SELECT * FROM publishers")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := make([]*models.Publisher, 0)

	for rows.Next() {
		p := new(models.Publisher)
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName)

		if err != nil {
			return nil, err
		}
		data = append(data, p)
	}
	return data, nil
}

func (m *DbManager) QuerySinglePublisher(id int) (*models.Publisher, error) {
	row := m.Db.QueryRow("SELECT * FROM publishers WHERE id = ?", id)

	pub := new(models.Publisher)
	if err := row.Scan(&pub.Id, &pub.FirstName, &pub.LastName); err == sql.ErrNoRows {
		return nil, err
	}

	return pub, nil
}

func (m *DbManager) InsertPublisher(firstName string, lastName string) (int64, error) {
	result, err := m.Db.Exec("INSERT INTO publishers(first_name, last_name) values (?, ?)", firstName, lastName)

	if err != nil {
		log.Printf("cannot add publisher to database (%s)\n", err)
		return 0, err
	}

	return result.LastInsertId()
}
func (m *DbManager) UpdatePublisher(id int, firstName string, lastName string) (int64, error) {
	result, err := m.Db.Exec("UPDATE publishers set first_name = ?, last_name = ? WHERE id = ?", firstName, lastName, id)

	if err != nil {
		log.Printf("cannot modify publisher with id %d in database (%s)\n", id, err)
		return 0, err
	}

	return result.RowsAffected()
}
func (m *DbManager) DeletePublisher(id int) (int64, error) {
	result, err := m.Db.Exec("DELETE FROM publishers WHERE id = ?", id)

	if err != nil {
		log.Printf("cannot delete publisher from database (%s)\n", err)
		return 0, err
	}

	return result.RowsAffected()
}
