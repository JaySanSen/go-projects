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