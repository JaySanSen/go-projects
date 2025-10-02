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



	employeeIds := [3]int{1, 2, 3}
	backUpEmployeeIds := employeeIds
	employeeIds[1] = 2000
	fmt.Println(employeeIds)
	fmt.Println(backUpEmployeeIds)
  [1 2000 3]
  [1 2 3]
```


### Slices
- slices are a subset of an array
- to create a slice we need to specify a starting point and a length
- slices don't contain their own data
- slices are reference types and they refer to data stored somewhere else
- so if the parent array from which the slice is created is updated then the slice will reflect that changed value
- vice versa is also true if the slice is updated then the array will reflect it
- a benefit of this referential nature of the slice is that it allows to make the slice a dynamic data type whic means the size of the slice can grow over time
- slices are copied by reference, to clone the values use slices.Clone


```
var sliceExample []int
signature is similar to array but just that we don't specify the size

fmt.Println(sliceExample)    -> output: [] (nil)

Why is the output nil

So nil is the zero value for un-initialized pointers and reference types like how 0 is the zero value for int.
since slice is a reference type it prints nil since it is currently uninitialized


var s []int{1, 2, 3}   //slice literal
s[1] = 99
fmt.Println(s)     // [1, 99, 3]



```


#### Adding data to the slice
- append is a built in method which is used to add elements to a slice by providing comma separated values.
```
s = append(s, 5, 10, 15)      // add elements to the slice
fmt.Println(s)     // [1, 99, 3, 5, 10, 15]
```
- so why do we do s = append(s ...)
- so append does not modify the slice it returns a new slice so we assign it back to original slice


#### Delete data from a slice
- There is a built in delete method BUT THIS DOES NOT work on slices
- Use standard library instead
```
s = slices.delete(s, startingIndex, endingIndex)
s = slices.delete(s, 1, 3)  // will remove 99(index 1) and 3(index 2 and stop before index 3) from [1, 99, 3, 5, 10, 15]   final output [1, 5, 10, 15]
startingIndex -> initial index we want to remove -> which mean index to start removing elements from
endingIndex -> index up to which it should delete but not including this end index
```
- slices are copied by reference
```
s := []string{"foo", "bar", "baz"}
s2 := s

s[0], s2[2] = "qux", "fred"
fmt.Println(s, s2)
["qux", "bar", "fred"] ["qux", "bar", "fred"]

s == s2    -> this is will give a compile time error - slices are not comparable because they are reference types
if there is a requirement to compare use the standard library


to clone as a copy use the standard library and clone using slices.Clone



	employeeIds := []int{1, 2, 3}
	backUpEmployeeIds := employeeIds
	employeeIds[1] = 2000
	fmt.Println(employeeIds)
	fmt.Println(backUpEmployeeIds)
  [1 2000 3]
  [1 2000 3]
```




### Maps
- Data structure that stores a collection of unordered key-value pairs
#### Key Characteristics of Go Maps:
- Key-Value Pairs: Maps store data as pairs, where each unique key is associated with a specific value.
- Unique Keys: Keys within a map must be unique and of a comparable type (e.g., int, string, float64, structs where all fields are comparable). Slices and non-comparable arrays/structs cannot be used as keys. 
- Flexible Values: Values can be of any type, including other maps, pointers, or reference types. 
- Unordered: The order of elements when iterating over a map is not guaranteed and can vary between iterations.
- References to Hash Tables: Maps are reference types; passing them around is efficient as it only passes a reference to the underlying hash table.
- map is a reference type
- unlike arrays and slice where order is preserved maps don't preserve the order
- maps are copied by reference similar to slices, to copy by value use maps.Clone to clone
```
var m map[string]int      // declare map with key of type string an value of int
fmt.Println(m)   // will return map[] which is nil because map is a reference type


m = map[string]int{"foo":1, "bar":2}


fmt.Println(m["foo"])     -> lookup values in a map
m["bar"] = 100            -> update values in a map

delete(m,"foo")           -> remove entry in a map


m["new"] = 500            -> add values to the map using keys which are not in the map yet, similar to update syntax

```
#### What happens when you try to fetch a key value pair that is not in the map
- the query always returns a result
	- this means that if the key is not present in the map the zero value for that particular value which they key was pointing to is returned

```
	sampleMap := map[string]int{"foo": 1, "bar": 2}
	delete(sampleMap, "foo")
	fmt.Println(sampleMap)
	fmt.Println(sampleMap["foo"])


  map[bar:2]
  0        -> 0 is being printed here since the value is of type integer for this map and the zero value of integer is 0




	sampleMap := map[string]bool{"foo": true, "bar": false}
	delete(sampleMap, "foo")
	fmt.Println(sampleMap)
	fmt.Println(sampleMap["foo"])
	 

  map[bar:false]
  false     -> here false is being printed since th value is of type bool and false is the zero value of bool
```
#### The comma ok syntax
- another way to query the map where we ask for a boolean along with the map query
  - boolean will be true if the value is returned from the map
  - boolean will be false if the map does not contain the key value pair

```
	sampleMap := map[string]bool{"foo": true, "bar": false}
	delete(sampleMap, "foo")
	fmt.Println(sampleMap)
	v, ok := sampleMap["foo"]
	fmt.Println(v)
	fmt.Println(ok)
	map[bar:false]
	false
	false



	sampleMap := map[string]int{"foo": 1, "bar": 200}
	delete(sampleMap, "foo")
	fmt.Println(sampleMap)
	v,ok := sampleMap["foo"]
	fmt.Println(v);
	fmt.Println(ok);

	map[bar:200]
	0
	false

```