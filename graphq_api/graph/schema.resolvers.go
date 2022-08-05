package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/graphql_test/blogpb"
	"cesargdd/graphql_test/controllers"
	"cesargdd/graphql_test/graph/generated"
	"cesargdd/graphql_test/graph/model"
	"cesargdd/graphql_test/server"
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

var nc = server.NatsSvc()

// CreateBlog is the resolver for the createBlog field.
func (r *mutationResolver) CreateBlog(ctx context.Context, input model.NewBlog) (*blogpb.Blog, error) {
	return controllers.CreateBlog(ctx, input)
}

// DeleteBlog is the resolver for the deleteBlog field.
func (r *mutationResolver) DeleteBlog(ctx context.Context, id int) (*blogpb.Blog, error) {
	return controllers.DeleteBlog(ctx, id)
}

// Blogs is the resolver for the blogs field.
func (r *queryResolver) Blogs(ctx context.Context) ([]*blogpb.Blog, error) {
	return controllers.Blogs(ctx)
}

// Blog is the resolver for the blog field.
func (r *queryResolver) Blog(ctx context.Context, id int) (*blogpb.Blog, error) {
	return controllers.Blog(ctx, id)
}

// Hello is the resolver for the hello field.
func (r *subscriptionResolver) Hello(ctx context.Context) (<-chan string, error) {
	msg := make(chan string)
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		st := string(m.Data)
		go func() { msg <- st }()
	})
	return msg, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
