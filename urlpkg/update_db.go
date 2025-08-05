package urlpkg

import "database/sql"

func updateDb(db *sql.DB, short_url, long_url string, userID int) error {
	_, err := db.Exec("UPDATE urls SET short_url = ? WHERE long_url = ? AND userID = ?", short_url, long_url, userID)
	if err != nil {
		return err
	}
	return nil
}
