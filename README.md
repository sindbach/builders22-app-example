# An Example App for MongoDB Builder's Fest 2022 

## With The Container

```sh
docker build -t example-app . 
docker run -p 8080:8080 example-app 
```

## Without The Container 

```sh
go run ./app.go 
```
