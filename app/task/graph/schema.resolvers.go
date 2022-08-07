package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gogo-graphql/app/task/graph/generated"
	"gogo-graphql/app/task/graph/model"
	"gogo-graphql/app/task/graph/models"
)

// Author is the resolver for the author field.
func (r *approachResolver) Author(ctx context.Context, obj *models.Approach) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// DetailList is the resolver for the detailList field.
func (r *approachResolver) DetailList(ctx context.Context, obj *models.Approach) ([]model.ApproachDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task is the resolver for the task field.
func (r *approachResolver) Task(ctx context.Context, obj *models.Approach) (*models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskList is the resolver for the taskList field.
func (r *meResolver) TaskList(ctx context.Context, obj *models.Me) ([]models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

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
func (r *queryResolver) TaskInfo(ctx context.Context, id string) (*models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, term string) ([]model.SearchResultItem, error) {
	panic(fmt.Errorf("not implemented"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*models.Me, error) {
	panic(fmt.Errorf("not implemented"))
}

// VoteChanged is the resolver for the voteChanged field.
func (r *subscriptionResolver) VoteChanged(ctx context.Context, taskID string) (<-chan *models.Approach, error) {
	panic(fmt.Errorf("not implemented"))
}

// ApproachList is the resolver for the approachList field.
func (r *taskResolver) ApproachList(ctx context.Context, obj *models.Task) ([]models.Approach, error) {
	panic(fmt.Errorf("not implemented"))
}

// Author is the resolver for the author field.
func (r *taskResolver) Author(ctx context.Context, obj *models.Task) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Approach returns generated.ApproachResolver implementation.
func (r *Resolver) Approach() generated.ApproachResolver { return &approachResolver{r} }

// Me returns generated.MeResolver implementation.
func (r *Resolver) Me() generated.MeResolver { return &meResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type approachResolver struct{ *Resolver }
type meResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
