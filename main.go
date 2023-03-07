package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func main() {
	// load .env file
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// DB connection
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	// get a db handle
	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	// ping the db
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to DB")

	// get albums by artist
	albums, err := albumByArtist("Pink Floyd")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Albums by Pink Floyd: ", albums)

	// get album by id
	alb, err := albumById(4)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Album by id: ", alb)

	// add album
	id, err := addAlbum(Album{
		Title:  "More",
		Artist: "Pink Floyd",
		Price:  9.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added album with id: ", id)

	// update album
	msg, err := updateAlbum(Album{
		ID:     5,
		Title:  "The Endless River",
		Artist: "Pink Floyd",
		Price:  9.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)

	// delete album
	msg, err = deleteAlbum(5)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}

func albumByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query(
		"SELECT * FROM albums WHERE artist = ?", name,
	)

	if err != nil {
		return nil, fmt.Errorf("error querying db: %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return albums, nil
}

func albumById(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow(
		"SELECT * FROM albums WHERE id = ?", id,
	)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		return alb, fmt.Errorf("error scanning row: %v", err)
	}

	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)",
		alb.Title, alb.Artist, alb.Price,
	)

	if err != nil {
		return 0, fmt.Errorf("error inserting album: %v", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error getting last inserted id: %v", err)
	}

	return id, nil
}

func updateAlbum(alb Album) (string, error) {
	_, err := db.Exec(
		"UPDATE albums SET title = ?, artist = ?, price = ? WHERE id = ?",
		alb.Title, alb.Artist, alb.Price, alb.ID,
	)

	if err != nil {
		return "", fmt.Errorf("error updating album: %v", err)
	}

	return "Album updated", nil
}

func deleteAlbum(id int64) (string, error) {
	_, err := db.Exec(
		"DELETE FROM albums WHERE id = ?", id,
	)

	if err != nil {
		return "", fmt.Errorf("error deleting album: %v", err)
	}

	return "Album deleted", nil
}
