# TODO APP in Golang 🙊

This is my very first attempt to creating a TODO App from zero
with `Golang` while learning the programming languge 🔥.
This project use only one third party dependency, this is a package
to handle the sqlite3 connection.  

All the `http server` and routes and all, is managed by `Golang standard library` 🧰 - https://pkg.go.dev/std.

1. Use of `net/http` standard library
2. Use of `Structs`.
3. Use of `pointers`.
4. Use of `Struct Methods`.
5. Use of the `database/sql` standard library.
6. Use of `flags` for args on execute.


## TODO 📝
- [x] Create HTTP Server with standard library.
- [x] Add one HTTP Route.
- [x] Check for HTTP Method on route validation.
- [x] Use prepared statements to avoid overhead.
- [ ] Better / improve error handler
- [ ] Improve the router with Mux.
- [ ] Use singleton pattern for database connection.
- [ ] Add tests
- [ ] Complete the README.

## Getting Started ⚒️

To run the project, you must installed or clone it in your computer.

```bash 
$ git clone git@github.com:migantoju/go-todo-app.git
```

then install the only one dependency.
```sh
$ go get github.com/mattn/go-sqlite3
```

and finally you can build the project or, run.

* use this for the first time to create the `sqlite3` database with the table. 
```bash 
$ go build *.go

$ ./main --migrate
```

When you make the first time migration, then you don't need to make it anymore, only run the server executing the binary file.
```bash
$ ./main
```

Maded with ❤️ by Miguel Toledano.