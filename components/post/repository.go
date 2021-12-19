package post

import (
	"context"
)


func (d *Dependency) FindAll(ctx context.Context) ([]Post, error) {
	query := `SELECT id, title, 'desc' FROM posts`
	rows, err := d.DB.QueryContext(ctx, query)
	if err != nil {
		return []Post{}, nil
	}
	defer rows.Close()
	var posts []Post
	for rows.Next(){
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Desc)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (d *Dependency) FindId(ctx context.Context, postId int) ([]Post, error) {
	query := `SELECT id, title, 'desc' FROM posts where id=?`
	rows, err := d.DB.QueryContext(ctx, query, postId)
	if err != nil {
		return []Post{}, err
	}
	defer rows.Close()
	var posts []Post
	if rows.Next(){
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Desc)
		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}