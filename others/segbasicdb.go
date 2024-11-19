package others

import (
	"database/sql"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB // Variabel db dideklarasikan secara global

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func AddAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func AlbumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func AlbumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// func main() {
// 	fmt.Println("Adnan Abdullah Juan | 5025221155")
// 	fmt.Println("Golang Web Application")

// 	// --- SESI DATABASE

// 	db = config.ConnectDb()
// 	defer db.Close()

// 	// --- CREATE ALBUM
// 	// newAlbum := Album{
// 	// 	Title:  "Go Lang Greatest Hits",
// 	// 	Artist: "John Doe",
// 	// 	Price:  29.99,
// 	// }

// 	// albumID, err := addAlbum(newAlbum)
// 	// if err != nil {
// 	// 	log.Fatalf("Gagal menambahkan album: %v", err)
// 	// }

// 	// fmt.Printf("Album baru berhasil ditambahkan dengan ID: %d\n", albumID)

// 	// --- READ ALBUM
// 	albums, err := albumsByArtist("John Doe")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Albums found: %v\n", albums)

// 	// --- READ ALBUM BY ID
// 	alb, err := albumByID(1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Album found: %v\n", alb)

// }
