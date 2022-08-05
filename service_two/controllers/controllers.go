package controllers

import (
	"cesargdd/service_two/blogpb"
	"cesargdd/service_two/server"
	"context"
	"fmt"
)

type Server struct {
	blogpb.BlogServiceServer
}

var blogSvc = server.BlogSrv()
var nc = server.NatsSvc()

func (*Server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create blog service two")
	res, err := blogSvc.CreateBlog(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (*Server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Delete blog service two")
	res, err := blogSvc.DeleteBlog(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (*Server) GetBlog(ctx context.Context, req *blogpb.GetBlogRequest) (*blogpb.GetBlogResponse, error) {
	fmt.Println("GetBlog service two")
	res, err := blogSvc.GetBlog(ctx, req)
	if err != nil {
		return nil, err
	}
	nc.Publish("foo", []byte("Hello World"))
	return res, nil
}
func (*Server) GetBlogs(ctx context.Context, req *blogpb.GetBlogsRequest) (*blogpb.GetBlogsResponse, error) {
	fmt.Println("GetBlogs service two")
	res, err := blogSvc.GetBlogs(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
