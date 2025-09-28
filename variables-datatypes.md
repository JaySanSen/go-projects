## Data Types


### Simple Data Types

- There are 4 simple data types
  - Strings
  - Numbers
  - Booleans
  - Errors  -> go is unique in that the error type is built in


#### Strings
- "this is a string" -> interpreted string
- \`this is also a string\` -> uses '`' backticks this is called a raw string


> Difference between interpreted strings and raw strings
- interpreted strings interpret the characters in the string including escape characters
  - "this is a string \n this is a new line string" this will output
  ```
  this is a string
  this is a new line string
  ```
  it reads the escape character and formats string accordingly

- raw strings print out the exact string contents
  - \`this is a string \n this is a new line string\` -> will output
  ```
  this is a string \n this is a new line string
  ```
  - whatever is the contents between the \`\` it will print that so it can be useful to print things if format should be maintained 

  ```
  'line 1
  line 2
  line 3`
  ```


  #### Numbers
  - Integers - int 99,0,-937
  - Unsigned Integers - uint - unlike integers can't encode negative numbers
  - Floating point numbers - float32 and float64 - 6.02e23, 3.14
  - Complex Numbers - complex64 and complex128 - 1+2i, 0.833i, 6.02e23 + 3.1415i


  #### Booleans
  - true
  - false
  > Booleans are first class citizens in GO which mean unlike other programming languages where numbers can sometimes be used as true or false for example 0 is considered false, will not go through in GO. Booleans must be used where booleans are required. There is no truthy and falsy values like in javascript


  #### Errors
  - The error built in interface type is the conventional interface for representing an error condition, with nil value (which is the null in GO) representing no error.
  - Error is a first class citizen
  - Unlike C# or Java where we can throw or raise exceptions, there is no such mechanism in GO. Simply return an error when things go wrong.
  ```
  type error interface{
    Error() string
  }
  ```


  #### Declaring Variables
  > var myVariable string
  - syntax might be awkward if coming from other languages like C# or Java
  - this is actually an optimization made for the compiler
  - anytime in go if we don't define a value for the variable then it is going to assign what is called a zero value
    - the zero value of a string is an empty string.
    - the zero value of a number is 0
    - the zero value of a boolean is false
  > var myName string = "GO User"
  > var myName = "GO User" // initialize with inferred type
  - specifying the type is not required as the compiler will infer the type

  #### Alternate variable declaration syntax
  - this is very common and is the most used way of declaring variables
  > myName := "Go User"   // short declaration syntax



  #### Constants
  > const a = 42   -> implicitly typed constant

- another thing to note here is that the data type is not mentioned.
  so the compiler will take it as what it needs it to be at that particular execution which may be an integer or a floating point or a unsigned int.


  > const b string = "Hello world"   -> explicitly typed constant
  > const c = a  -> assign one constant to another

- if you want to declare multiple constants at one time
```
const (
  d = true
  e = 3.14
)

```

- expressions are also allowed
> const c = 2 * 5
> const d = "Hello" + " world"
- expressions are allowed as long as it is calculable/must be able to be evaluated at compile time.

> const e = someFunction()   -> THIS IS NOT ALLOWED



#### some additional information on short hand declaration
-  Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
- Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available. 

```
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

```