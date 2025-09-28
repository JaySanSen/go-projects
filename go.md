## Go steps

- Help Docs
  - To access the command line help docs use go help
```
go help mod   -> this will show the docs for mod command
go help mod init  -> this will show docs for mod init command
```

### Initialize a module
- The module name is optional
- Running go mod init demo will create a file called go.mod

```
go mod init <module-name>
```

- The contents of the file will look something like
```
module demo

go 1.25.1
```
- This is something similar to a namespace

### Creating a program file
- All go files end with the .go extension
- Go is a package based language so every file must contain a package name
```
package main
package helloworld
```
- If the package is marked as main like `package main` that means this package contain executable program.
- The main package usually contains the main method which acts as the starting point
- For libraries and other types of packages any name is fine `package utils`

### Run a go program
```
go run .   -> to run the go project in the current directory

go run project/main.go    -> to run the project in a specific directory
```