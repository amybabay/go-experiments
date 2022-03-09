# go-experiments

Starting from the tutorial
[https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started).


## code organization

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```
go mod init github.com/amybabay/go-experiments
```

To get dependencies (e.g. quote package from hello example):
```
go mod tidy
```

Notes to remember:
- function name that starts with a capital letter can be called by a function
  not in the same package ([exported names](https://go.dev/tour/basics/3))

- ``:=`` operator is a shortcut for declaring and initializing a variable
