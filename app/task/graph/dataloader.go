package graph

import (
	"context"
	"gogo-graphql/app/task/graph/models"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

type loaderKey string

const userLoaderKey = loaderKey("userLoader")

func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userLoader := NewUserLoader(
				UserLoaderConfig{
					Wait:     1 * time.Millisecond,
					MaxBatch: 100,
					Fetch: func(ids []int64) ([]*models.User, []error) {
						var users []*models.User
						err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
						if err != nil {
							return nil, []error{err}
						}

						u := make(map[int64]*models.User, len(users))

						for _, user := range users {
							u[user.ID] = user
						}

						result := make([]*models.User, 0, len(ids))
						for _, id := range ids {
							result = append(result, u[id])
						}

						return result, nil
					},
				},
			)

			ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
