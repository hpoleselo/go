# Learning Go

Go Perks:

- Compiled language (single binary file)
- Concurrency out of the box
- Load average is low (uses few resources, is efficient), memory usage is less (Remember about PHP Laverel server comparison), mainly because Go is compiled and PHP is interpreted.
- Excellent package manager
- Statically typed (diferente de Python)

Mod file:
Tells the compiler that our application uses Go modules (standard library) and external dependencies, is it our requirements.txt?
We can name it whatever we want (our app name), doesn't have to be related to main or whatever.

`$ go mod init NAME_OF_YOUR_APPLICATION`

When we're making our application modular, i.e. moving all functions and handlers from one program to another and to make the handlers and function available to the main.go, instead of importing them on the main, what we have to do is just create a mod.go, which will make everything visible to the main.go

To run multiple Go files at the same time
Bash:

`$ go run *.go`

Windows:

`$ go run .`

The mod file is important because when we want to add external dependencies, just by using the command of downloading a package:

`$ go get -u github.com/go-chi/chi`

Which is by the way a router package for handlers (HTTP requests), it automatically populates go.mod file with the recent installed package.

If we have unused packages in our programs, just go to root of your app and

`$ go mod tidy`

When we have different programs outside main.go they have package main in the beginning of the program, that is because they should be part of the main package, which is your main.go, but usually we create separate folders (packages) for each program and those are referenced in the main.go.

nil = NULL, represents zero value for pointers.

Functions/variables that start with capital letter are able to be accessed by outside a given package. (Same as in Java as making it as public function)

Variodic Functions: can take one, more or none parameters. Println() is one of them.

## Types

```go
//Declaration:

var s string

// inferes it is a string data type
var s = "seven"
```

Struct:
	Struct is the way to create user-defined concrete types. Struct is a collection of fields or properties. Instead of calling as a Class in other programming languages we call it as a Struct. We don't have to set setters() and getters() for a struct, because when a struct is with capital letter it's already available for other packages.


Defining a behavior for our Struct:
	Interface is the word we're looking for, instead of method as we see in OOP languages.
	In Golang terms: interface types provide contracts to concrete types, basically let us define behaviors for our objects. 
	
	"A method in Go is simply a function that is declared with a receiver. If you want to modify the data ofa receiber from the method, the receiver must be a pointer."
		
Inheritance in Go, I mean, COMPOSITION:
	Go's type system doesn't support inheritance. Composition is preferred. 

## Resources

Initial Tutorials:

About structs and interfaces: https://thenewstack.io/understanding-golang-type-system/

https://cloud.google.com/bigquery/docs/samples/bigquery-query

https://pkg.go.dev/cloud.google.com/go/bigquery

https://github.com/googleapis/google-cloud-go

https://cloud.google.com/bigquery/docs/samples/bigquery-update-table-cmek

https://cloud.google.com/docs/authentication/production#auth-cloud-explicit-python