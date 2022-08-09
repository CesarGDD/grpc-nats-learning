package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/graphql_test/blogpb"
	"cesargdd/graphql_test/controllers"
	"cesargdd/graphql_test/graph/generated"
	"cesargdd/graphql_test/graph/model"
	"context"
)

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

// Blog is the resolver for the blog field.
func (r *subscriptionResolver) Blog(ctx context.Context) (<-chan *blogpb.Blog, error) {
	return controllers.BlogSubscription(ctx)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
