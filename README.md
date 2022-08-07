# gogo-graphql
graphql learning from graphql in action in golang

# init
```bash
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go

go mod tidy

go run github.com/99designs/gqlgen init

go run server.go

go run github.com/99designs/gqlgen generate

cd app/task/graph
#if failed, try to generate file gen.go and write package generated inside
go run github.com/vektah/dataloaden UserLoader int64 '*gogo-graphql/app/task/graph/models.User'
```

# Youtube
[youtube]("https://youtube.com/playlist?list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn")