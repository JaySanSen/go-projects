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

fmt.Println(len(arr))   //3 -> length of array


arr := [3]string{"foo", "bar", "foobar"}
arr2 := arr    -> arrays are copied by value
so they are independent of each other
arr[0] = "baz"
fmt.Println(arr)     -> ["baz", "bar", "foobar"]
fmt.Println(arr2)    -> ["foo", "bar", "foobar"]

arr == arr2 -> false  -> GO checks value by value and returns boolean result.
Arrays are comparable
```


### Slices
- slices are a subset of an array
- to create a slice we need to specify a starting point and a length
- slices don't contain their own data
- slices are reference types and they refer to data stored somewhere else
- so if the parent array from which the slice is created is updated then the slice will reflect that changed value
- vice versa is also true if the slice is updated then the array will reflect it
- a benefit of this referential nature of the slice is that it allows to make the slice a dynamic data type whic means the size of the slice can grow over time