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
- [A curated list of awesome Go frameworks](https://awesome-go.com/)

### Moving from Go way of API Dev to Python (Django) way of API Dev

I think a lot has changed since then, coming to the frameworks beego, gingonic,echo seem to be more adopted now (Im just saying by the github stars I guess thats shouldn't be the only metric)

And Im from Django background so curios how to handle

- Serialization - I used DjangoRestFramework, I see mostly using struct pointer
- Apps - In Django everything is organized into group of apps - I dont see any such structure suggestion here, (Could be a go way of doing things) so how to organize different independent parts of same project.
- ORM - Coming from a Django background I was very happy with the django ORM, but here mostly I see developers encourage writing SQL queries. And majorly `Migrations` , I miss using migrations in Go lang.
-

---

# Why Go?

- How to see Golang and What to compare? : Get right glasses
-

## Simplicity is Complicated

- Adding features doesn't make go better? (take js as example)
  - What is does this mean to me?
    - Only one way of doing things is more likely ensured (as you have limited things in hand)
    - Refactoring Effort
    - Harder to understand (Readability)
    - We want simplicity (not complexity)
- Only right features are added

- Long term maintenance in mind

Examples:

No to map and filter -as they might be more expensive in terms of computer

- Vector space examples - Ensures one way of doing things

## Concurrency

- 3 key strokes

## Constants

- int vs float

## Interfaces

- Many of the best practices are enforced by the language by default
- Best practices are enforced by the language itself

  - Encapsulation
  - Coding to interface than implementations
  - Do not depend on the internals
  - Proper sharing of logic - If there is a logic that is being shared your properly
  - You aren't loosing anything from this way of writing than properly - Need to write properly

- Don't

## Packages

-

## Code Resuability

- Composition or Embedding over

## Copy by reference vs Copy by value mess

- There is only pass by reference

## Clean Code is encouraged

- Proper braces is mandatory for the code to run

---

- Cockroach DB Team - https://www.youtube.com/watch?v=hWNwI5q01gI

  - Performance of GC - More Control

- Simplicity - https://www.youtube.com/watch?v=k9Zbuuo51go&t=0s

TODO:

#### Code Reuse In Golang

- Google Understanding Go Interfaces - https://www.youtube.com/watch?v=F4wUrj6pmSI
- The Challenges of Writing a Massive and Complex Go Application - https://www.youtube.com/watch?v=hWNwI5q01gI&t=1024s

- There is no pass by reference - https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
- Pass By Reference vs Pass By Value - https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/

- Using Generics - http://tinystruggles.com/2015/08/29/golang-code-reuse.html
- OOP without inheritance - https://yourbasic.org/golang/inheritance-object-oriented/

- https://stackoverflow.com/questions/31997753/reuse-a-function-across-multiple-structs-to-satisfy-an-interface
- https://forum.golangbridge.org/t/trying-not-to-reuse-code/12195
- http://jmoiron.net/blog/idiomatic-code-reuse-in-go/

## Design in Go

- https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html
- https://www.reddit.com/r/programming/comments/1mue70/less_is_exponentially_more_rob_pike_on_why_c_devs/
- Less is more Complicated - https://www.youtube.com/watch?v=k9Zbuuo51go&t=0s

- What makes Go Different - https://www.youtube.com/watch?v=FEFXjRoac_U&t=0s
- The Design of Go Assembler - https://www.youtube.com/watch?v=KINIAgRpkDA&t=117s
- Concurrency is not Parallelism - https://www.youtube.com/watch?v=cN_DpYBzKso&t=69s
- Go Proverbs - https://www.youtube.com/watch?v=PAAkCSZUG1c&t=287s
- Meet the Go Team - https://www.youtube.com/watch?v=sln-gJaURzk&t=154s
- Simplicity is Complicated - https://talks.golang.org/2015/simplicity-is-complicated.slide#1


## Package Management

- https://blog.golang.org/using-go-modules
- https://www.activestate.com/blog/golang-module-vs-dep-pros-cons/

## Clean Code

- https://medium.com/@teivah/good-code-vs-bad-code-in-golang-84cb3c5da49d

#### SOLID

- SOLID Go Design - https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=1s
- Idiomatic Go - https://www.youtube.com/watch?v=yeetIgNeIkc
- Go For Python Dev - https://www.youtube.com/watch?v=AVxosSFzq5s
- Go Best Practices - https://www.youtube.com/watch?v=8D3Vmm1BGoY

#### Extras

- Things in Go I Never Use - https://www.youtube.com/watch?v=5DVV36uqQ4E&t=3s
- GoLang for DevOps
  - https://qarea.com/blog/golang-for-devopss
  - https://qarea.com/blog/golang-web-development-better-than-python
  - https://qarea.com/blog/8-reasons-you-need-to-go-golang#:~:text=Golang%20is%20fast&text=It%20is%20devoid%20of%20classes,development%20is%20faster%20and%20cheaper.
- Interfaces
  - https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c
  - How structures are better than classes - https://medium.com/@simplyianm/why-gos-structs-are-superior-to-class-based-inheritance-b661ba897c67
  - Structs Instead of Classes - OOP in Go - https://golangbot.com/structs-instead-of-classes/
  - Is Go object oriented - https://flaviocopes.com/golang-is-go-object-oriented/

## Go vs Python

- https://analyticsindiamag.com/the-war-between-python-and-go-explained-in-8-points/
-
