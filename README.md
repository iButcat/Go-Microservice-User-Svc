# Go-Microservivce-User-Svc

A Goland microservice using Go-kit

<h1>About</h1>

A CRUD microservice in Go and Go-kit.Go kit is a programming toolkit for building microservices (or elegant monoliths) in Go. We solve common problems in distributed systems and application architecture so you can focus on delivering business value. 

<h1>Goal</h1>

I did this application while reading about Uncle's bob clean Architecture book. I got really curious for software design and architecture in general. 
I learned how Microservice works and how to make the code follow good design patterns.

<h1>Clean Architecture</h1>

![Clean Architecture](https://github.com/iButcat/Go-Microservice-User-Svc/blob/master/pic/onion.png)

<h1>Install</h1>

Download the source code and then add your database credentials in main.go:

```
dsn := "user=user password=pass dbname=dbname port=5432 sslmode=disable"
```

You should be able to start the application by running:
```
go run * http.addr 8080
```

<h1>Endpoints</h1>

- Create
```
127.0.0.1:8080/user
```
- Get 
```
127.0.0.1:8080/user/{id}
```
- Update 
```
127.0.0.1:8080/user/update
```
- Delete 
```
127.0.0.1:8080/user/{id}
```
