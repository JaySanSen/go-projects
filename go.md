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
```
In Go, package main is a special declaration that signifies a package containing an executable program. This declaration informs the Go compiler that the code within this package should be compiled into a standalone executable binary, rather than a reusable library.

Key characteristics of package main:

    Executable Entry Point:
    A package main must contain a func main() function. This main() function serves as the entry point for the executable program, meaning the program's execution begins here.

    Standalone Programs:
    Packages declared as main are designed to be run directly. They are the top-level components of a Go application.

    Compilation:
    When you compile a Go program with package main (e.g., using go build), it produces an executable file.

    Contrast with Other Packages:
    Other packages in Go (e.g., package utils, package database) are typically designed to provide reusable functionality to be imported and used by other packages, including main packages. They do not contain a main() function and cannot be directly executed.
```

### Key differences between a Module and a Package
# Go Modules vs Packages

In Go, **Modules** and **Packages** are fundamental concepts for organizing and structuring your code. Below is a detailed comparison, along with examples to help you understand the differences.

| **Feature**                   | **Module**                                           | **Package**                                              |
|-------------------------------|------------------------------------------------------|----------------------------------------------------------|
| **Definition**                 | A **module** is a collection of Go packages that are versioned together and defined in a `go.mod` file. | A **package** is a collection of Go source files grouped together under a common name. |
| **Scope**                      | A module is a higher-level unit that includes multiple packages. It is defined at the root of your project. | A package is a lower-level unit that contains the actual implementation and logic. |
| **go.mod File**                | The `go.mod` file is used to define the module, its name, dependencies, and the Go version. It’s the entry point for dependency management. | A package does not have a separate configuration file; it’s defined by the directory structure and the `package` statement in the source files. |
| **Naming**                     | The **module name** is typically defined when running `go mod init` (e.g., `module demo`). | The **package name** is defined within Go files using the `package` keyword (e.g., `package main`). |
| **Usage**                      | A module is used for dependency management, versioning, and managing the overall structure of the project. It is the root of the project. | A package is used to group related functionality. It is imported into other Go files or packages to provide reusable functionality. |
| **Versioning**                 | The module is versioned using tags (e.g., `v1.0.0`) and stored in version control (e.g., GitHub). | Packages do not have individual versioning; they follow the module's version. |
| **Importing**                  | When importing a module, you reference its module path as defined in the `go.mod` (e.g., `github.com/user/module`). | When importing a package, you reference its path within the module (e.g., `github.com/user/module/utils`). |
| **Main Entry Point**           | The module is the overall project or repository. The module itself does not contain a `main()` function but can contain executable code in a `main` package. | A package (such as `main`) contains the actual logic. A `main` package can be used to define the entry point of a Go application (i.e., `main()` function). |
| **Example**                    | `go mod init github.com/myusername/demo` creates a module. | `package main` or `package utils` defines a package. |
| **Relationship**               | A module can house multiple packages. A module is the root, and it organizes code into packages. | A package is a folder that contains Go files with the same package name. |
| **Compilation**                | A module is not compiled on its own but contains Go packages that are compiled. The module is used for dependency management and version control. | A package is compiled when the Go files within it are built. If the package is part of a `main` package, it will be executed. |
| **Example Structure**          | `demo/` is a module containing several packages. | `utils/` is a subdirectory (package) inside the module. |
| **Best Practice**              | Modules should represent complete projects or reusable components and be versioned. | Packages should be logically separated (e.g., `utils`, `handlers`, `models`) for better code organization. |

---

## Example Project Structure

Here’s how you might structure your Go project with both **modules** and **packages**:



### Run a go program
```
go run .   -> to run the go project in the current directory

go run project/main.go    -> to run the project in a specific directory
```