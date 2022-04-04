package album

import (
	"database/sql"
	"fmt"

	"github.com/jluo926/GoProjects/web-service/internal/db"
)

func FindAll() ([]Album, error) {
	rows, error := db.Pool.Query("select * from album")
	if error != nil {
		return nil, fmt.Errorf("error find all %v", error)
	}
	defer rows.Close()
	var res []Album
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("find all %v", err)
		}
		res = append(res, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("find all %v", err)
	}
	return res, nil
}

// albumByID queries for the album with the specified ID.
func FindById(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.Pool.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, err
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func Add(alb Album) (int64, error) {
	result, err := db.Pool.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
