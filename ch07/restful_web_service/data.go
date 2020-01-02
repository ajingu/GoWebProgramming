package main

import(
	"database/sql"
	_ "github.com/lib/pq"
)

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

func Retrieve(id int)(post Post, err error){
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	return
}

func (post *Post) Update() (err error){
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error){
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
