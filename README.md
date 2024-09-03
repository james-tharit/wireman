## Wireman
- Graphql server, BFF for Anonymity Project

### For generate graphQL, using the command
```sh
go run github.com/99designs/gqlgen generate
```

### ENV required
```
PORT=5423
```
### Build Binary and Run locally
```sh
 go build -o ./out/dist && ./out/dist
```

### Build Docker container
```sh
docker build -t wireman .
```

### Run Docker container
```sh
docker run -p 5423:5423 wireman
```

### Add new mod
1. add new mod to `./tools.go`
2. run `go mod tidy`