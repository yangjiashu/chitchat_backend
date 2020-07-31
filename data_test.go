package data

import (
	"fmt"
	"testing"
)

func TestFindCommentsByPostId(t *testing.T) {
	t.Skip()
	comments := FindCommentsByPostId(3)
	if len(comments) != 2 {
		t.Error("查询postid为3的评论，结果错误，期望的结果是2条，获得了", len(comments), "条")
	}
	if comments[0].Auth.Name != "yangjiashu" || comments[1].Auth.Name != "bean_man" {
		t.Error("查询作者错误")
	}
}

func TestListPosts(t *testing.T) {
	fmt.Println(ListPosts())
}
