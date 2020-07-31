package data

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=646233 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// 返回posts表中的所有数据,必须以特定的Post格式返回
// 测试通过
func ListPosts() (posts []Post) {
	// 查询所有的post和对应的author
	rows, err := Db.Query(
		`select p.id, p.content, a.id, a.name 
				from posts as p, authors as a 
				where a.id = p.author_id`)
	if err != nil {
		panic(err)
	}

	// 录入posts
	posts = []Post{}
	for rows.Next() {
		post := Post{}
		auth := Author{}
		err := rows.Scan(&post.Id, &post.Content, &auth.Id, &auth.Name)
		if err != nil {
			panic(err)
		}
		post.Auth = auth
		posts = append(posts, post)
	}

	// 查询comments
	for i := range posts {
		posts[i].Comments = FindCommentsByPostId(posts[i].Id)
	}

	return
}

// 根据postId查询comments
// 测试通过
func FindCommentsByPostId(id int) (comments []Comment) {
	rows, err := Db.Query(
		`select c.id, c.content, a.id, a.name
				from comments c inner join authors a 
				on c.author_id = a.id
				where c.post_id = $1`, id)
	if err != nil {
		panic(err)
	}

	comments = []Comment{}
	for rows.Next() {
		comment := Comment{}
		auth := Author{}
		err := rows.Scan(&comment.Id, &comment.Content, &auth.Id, &auth.Name)
		if err != nil {
			panic(err)
		}
		comment.Auth = auth
		comments = append(comments, comment)
	}

	return
}
