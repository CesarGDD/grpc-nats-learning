package controllers

import (
	"cesargdd/graphql_test/blogpb"
	"cesargdd/graphql_test/graph/model"
	"cesargdd/graphql_test/server"
	"context"
	"fmt"
)

type Server struct {
	blogpb.BlogServiceServer
}

var blogSvc = server.BlogSrv()

func CreateBlog(ctx context.Context, input model.NewBlog) (*blogpb.Blog, error) {
	fmt.Println("create blog graphql")
	res, err := blogSvc.CreateBlog(ctx, &blogpb.CreateBlogRequest{
		Username: input.Username,
		Content:  input.Content,
	})
	if err != nil {
		return nil, err
	}
	return res.Blog, nil
}

// DeleteBlog is the resolver for the deleteBlog field.
func DeleteBlog(ctx context.Context, id int) (*blogpb.Blog, error) {
	fmt.Println("Delete blog graphql")
	res, err := blogSvc.DeleteBlog(ctx, &blogpb.DeleteBlogRequest{Id: int32(id)})
	if err != nil {
		return nil, err
	}
	return res.Blog, nil
}

// Blogs is the resolver for the blogs field.
func Blogs(ctx context.Context) ([]*blogpb.Blog, error) {
	fmt.Println("get blogs graphql")
	res, err := blogSvc.GetBlogs(ctx, &blogpb.GetBlogsRequest{})
	if err != nil {
		return nil, err
	}
	return res.Blog, nil
}

// Blog is the resolver for the blog field.
func Blog(ctx context.Context, id int) (*blogpb.Blog, error) {
	fmt.Print("get blog")
	res, err := blogSvc.GetBlog(ctx, &blogpb.GetBlogRequest{Id: int32(id)})
	if err != nil {
		return nil, err
	}
	return res.Blog, nil
}
