# Beego Project

- Run Project

```shell script
$ bee run
...
```

## Tasks

- [x] Writing a Working API
- [x] Accessing Database
- [ ] Using OAuth / Any API Security
- [ ] Writing Tests
- [ ] Mocking
- [ ] Clean Architecture
- [ ] Deploying to AWS Lambda
- [ ] Building APIs

### One hour Tasks

Above tasks are made into smaller and specific chunks.

#### Set One

- [x] Beego ORM
- [x] Gorm

## Concepts

### Golang Concepts

- Structs
  - Inheritance
  - References:
    - [Structure Basics](https://medium.com/rungo/structures-in-go-76377cc106a2)
  - Anonymous Structures: [Refer](https://talks.golang.org/2012/10things.slide#1)
  
- Interfaces
  - [Reference](https://medium.com/rungo/interfaces-in-go-ab1601159b3a)
- Pointers/references
- Exported Identifiers
  - References:
    - [Intro](https://medium.com/golangspec/exported-identifiers-in-go-518e93cc98af)
- Json Parsing
  - [Reference I](https://yourbasic.org/golang/json-example/)
  - [Reference II](https://eager.io/blog/go-and-json/)
- Exceptions
  - [Reference](https://medium.com/rungo/error-handling-in-go-f0125de052f0)
- Debugging

### Beego Concepts

- Request Body Handling
  - [Accessing Request Body](https://stackoverflow.com/questions/30982891/beego-post-request-body-always-empty)

### Useful links

- [Go Web Examples](https://gowebexamples.com/mysql-database/)

### Moving from Go way of API Dev to Python (Django) way of API Dev

I think a lot has changed since then, coming to the frameworks beego, gingonic,echo seem to be more adopted now (Im just saying by the github stars I guess thats shouldn't be the only metric)

And Im from Django background so curios how to handle

- Serialization - I used DjangoRestFramework, I see mostly using struct pointer
- Apps - In Django everything is organized into group of apps - I dont see any such structure suggestion here, (Could be a go way of doing things) so how to organize different independent parts of same project.
- ORM - Coming from a Django background I was very happy with the django ORM, but here mostly I see developers encourage writing SQL queries. And majorly `Migrations` , I miss using migrations in Go lang.
-
