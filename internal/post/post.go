package post

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Post struct {
	Author    string
	Content   string
	CreatedAt string
}

func GetPosts() (posts []Post, err error) {
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	database, err := sql.Open("postgres", os.Getenv("credentials"))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		fmt.Printf("%v", err)
		return
	}

	rows, err := database.Query("SELECT author, content, created_at FROM users ORDER BY created_at DESC;")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		rows.Scan(&post.Author, &post.Content, &post.CreatedAt)
		post.CreatedAt = post.CreatedAt[11:16] + " " +  post.CreatedAt[:10]
		posts = append(posts, post)
	}
	return
}

func SavePost(author, content string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	database, err := sql.Open("postgres", os.Getenv("credentials"))
	if err != nil {
		return err
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		return err
	}

	querry := fmt.Sprintf("INSERT INTO users (author, content) VALUES ('%s', '%s');", author, content)
	_, err = database.Exec(querry)
	if err != nil {
		log.Println(err)
	}
	return nil
}
