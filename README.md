# go-teach
teaching wife how to go

## Part 1 Review API spec
  Checkout the branch Part1-apispec `git checkout Part1-apispec`
  1. Review the API specification and understand how to edit it.
    [API Specification](api/colorserver.yaml)
  2. Try it out in the swagger editor.
    [Swagger Editor](https://editor.swagger.io/)
  3. Review basic code structure of the server.
    [Color Server](main.go)

## Part 2 Implement routing
  Checkout the branch Part2-routing `git checkout Part2-request-response`
  1. Implement the routing for the server in the main function.
    [main.go](main.go#L28)

## Part 3 Implement Get
  Checkout the branch Part3-getcolor `git checkout Part3-get`
  1. Implement the GetColor function.
    [color.go](color.go#L41)

## Part 4 Implement Post
  Checkout the branch Part4-postcolor `git checkout Part4-post`
  1. Implement the PostColor function.
    [color.go](color.go#L78)

## Part 5 Manage Concurrent Updates to data store
  Checkout the branch Main `git checkout main` 
  1. Implement the lock for the data store.
    [color.go](color.go#L78)

## Part 6 Go commands
`go run main.go` runs a dev version of the server
`go build` builds and executable file that can be run in a container/ started from a command line


## Part 7 Curl commands

Try these commands with different colors

`curl -X GET http://localhost:8080/color/red`

`curl -X POST http://localhost:8080/color/blue`

`curl -X POST -H "Content-Type: application/json" -d '{"color":"blue"}' http://localhost:8080/color`


## Part 8 Make a container and store it in the container registry

`docker build -t dev/subscripify-new-color-server:latest .`

`docker tag dev/subscripify-new-color-server subscripifycontreg.azurecr.io/dev/subscripify-new-color-server:latest`

`docker push subscripifycontreg.azurecr.io/dev/subscripify-new-color-server:latest`

` docker run -p 8080:8080 dev/subscripify-new-color-server`

