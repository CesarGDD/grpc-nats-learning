package controller

import (
	"cesargdd/service_one/blogpb"
	"cesargdd/service_one/postgres"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

type Server struct {
	blogpb.BlogServiceServer
}

var conn = postgres.Connect()
var db = postgres.New(conn)

func (*Server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create Blog grpc")
	res, err := db.CreateBlog(ctx, postgres.CreateBlogParams{
		Username: req.Username,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       res.Id,
			Username: res.Username,
			Content:  res.Content,
		},
	}, nil
}
func (*Server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Delete blog gRPC")
	res, err := db.DeleteBlog(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &blogpb.DeleteBlogResponse{
		Blog: &blogpb.Blog{
			Id:       res.Id,
			Username: res.Username,
			Content:  res.Content,
		},
	}, nil
}
func (*Server) GetBlog(ctx context.Context, req *blogpb.GetBlogRequest) (*blogpb.GetBlogResponse, error) {
	fmt.Println("Get Blog gRPC")
	res, err := db.GetBlog(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &blogpb.GetBlogResponse{
		Blog: &blogpb.Blog{
			Id:       res.Id,
			Username: res.Username,
			Content:  res.Content,
		},
	}, nil
}
func (*Server) GetBlogs(ctx context.Context, req *blogpb.GetBlogsRequest) (*blogpb.GetBlogsResponse, error) {
	fmt.Println("Get Blogs gRPC")
	res, err := db.GetBlogs(ctx)
	if err != nil {
		return nil, err
	}
	data := &blogpb.GetBlogsResponse{}
	copier.Copy(&data.Blog, &res)

	return &blogpb.GetBlogsResponse{
		Blog: data.Blog,
	}, nil

}
