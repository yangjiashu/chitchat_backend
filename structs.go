package data

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Auth     Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Auth    Author `json:"author"`
	Content string `json:"content"`
}
