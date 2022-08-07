package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gogo-graphql/graph/generated"
	"gogo-graphql/graph/model"
)

// ApproachCreate is the resolver for the approachCreate field.
func (r *mutationResolver) ApproachCreate(ctx context.Context, taskID string, input model.ApproachInput) (*model.ApproachPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ApproachVote is the resolver for the approachVote field.
func (r *mutationResolver) ApproachVote(ctx context.Context, approachID string, input model.ApproachVoteInput) (*model.ApproachVotePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskCreate is the resolver for the taskCreate field.
func (r *mutationResolver) TaskCreate(ctx context.Context, input model.TaskInput) (*model.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserInput) (*model.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserLogin is the resolver for the userLogin field.
func (r *mutationResolver) UserLogin(ctx context.Context, input model.AuthInput) (*model.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CurrentTime is the resolver for the currentTime field.
func (r *queryResolver) CurrentTime(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskInfo is the resolver for the taskInfo field.
func (r *queryResolver) TaskInfo(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, term string) ([]model.SearchResultItem, error) {
	panic(fmt.Errorf("not implemented"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	panic(fmt.Errorf("not implemented"))
}

// VoteChanged is the resolver for the voteChanged field.
func (r *subscriptionResolver) VoteChanged(ctx context.Context, taskID string) (<-chan *model.Approach, error) {
	panic(fmt.Errorf("not implemented"))
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
