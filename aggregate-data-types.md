## Aggregate Data types

### Arrays
- Fixed size collection of a single data type
- Array can never change size
- Use the index to access the elements and arrays are 0 indexed
- static data structure

```
var arr [3]int    // array of 3 ints
fmt.Println(arr)   -> this will output [0,0,0]
why will it output all zeros?
it is because in GO any variable declared will be initialized to its zero value


arr = [3]int{1, 2, 3}    //array literal
fmt.Println(arr[1])

arr[1] = 100    //update value

```