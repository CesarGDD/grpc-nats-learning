package controllers

import (
	"cesargdd/graphql_test/blogpb"
	"cesargdd/graphql_test/graph/model"
	"cesargdd/graphql_test/server"
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type Server struct {
	blogpb.BlogServiceServer
}

var blogSvc = server.BlogSrv()
var nc = server.NatsSvc()

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

func BlogSubscription(ctx context.Context) (<-chan *blogpb.Blog, error) {
	blog := make(chan *blogpb.Blog)
	res := &blogpb.Blog{}
	unsub, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		go func() {
			fmt.Printf("Received a message: %s\n", string(msg.Data))
			if err := json.Unmarshal(msg.Data, &res); err != nil {
				fmt.Println(err)
			}
			blog <- res
		}()
	})
	if err != nil {
		return nil, err
	}
	go func() {
		<-ctx.Done()
		if err := unsub.Unsubscribe(); err != nil {
			_ = err
		}
		close(blog)
	}()
	return blog, nil
}
