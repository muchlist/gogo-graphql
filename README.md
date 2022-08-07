# gogo-graphql
graphql learning from graphql in action in golang

# init
```bash
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go

go mod tidy

go run github.com/99designs/gqlgen init

go run server.go

go run github.com/99designs/gqlgen generate
```