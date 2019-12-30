package main

import(
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct{
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct{
	Id int
	Content string
	Author string
	Post *Post
}

var Db *sql.DB

// postgresql driver initialization(automatically called)
func init(){
	var err error
	// set sql.DB instance(not creating connection)
	Db, err = sql.Open(
		"postgres",
		"user=gwp dbname=gwp password=gwp sslmode=disable")

	if err != nil{
		panic(err)
	}
}

func (comment *Comment) Create() (err error){
	if comment.Post == nil{
		err = errors.New("comments not found")
		return
	}

	err = Db.QueryRow(
		"insert into comments (content, author, post_id) values ($1, $2, $3) returning id",
		comment.Content,
		comment.Author,
		comment.Post.Id).Scan(&comment.Id)

	return
}

func (post *Post) Create() (err error){
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// create prepared statement
	stmt, err := Db.Prepare(statement)
    if err != nil {
    	return
	}
	defer stmt.Close()

	// execute prepared statement
    err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
    return
}

func GetPost(id int)(post Post, err error){
    post = Post{}
    post.Comments = []Comment{}
    err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

    rows, err := Db.Query("select id, content, author from comments")
    if err != nil{
    	return
	}

	for rows.Next(){
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil{
			return
		}

		post.Comments = append(post.Comments, comment)
	}

	rows.Close()
    return
}

/*
func (post *Post) Update() (err error){
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error){
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
*/

func main(){
	post := Post{Content: "Hello World", Author:"Ace"}
	post.Create()

	comment := Comment{Content:"Good Comment!", Author:"Ace", Post:&post}
	comment.Create()

	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}