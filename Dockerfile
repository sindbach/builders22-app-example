FROM golang:1.16-alpine 
WORKDIR /app
COPY app.go go.mod ./
RUN go build -o /builders22-app
EXPOSE 8080
CMD ["/builders22-app"]