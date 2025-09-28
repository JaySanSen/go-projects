## Pointers and Values

```
a := 42
b := a

So what happens here is a contains the value 42 and when b = a is assigned the compiler go and looks at the value and takes a copy and puts it to b.

a and b are independent of each other
changes to value of a don't affect b



a := 42
b := &a

& -> is called the address of operator.

here b is no longer holding a copy of value of a.
it is holding the location in memory that a is pointing to 
it is holding memory address a is pointing to
so b is pointing to a

```

### De referencing a pointer
```
a := 42
b := &a


To de reference a point use *

*b will give 42

if the value of a changes to 27 then *b will give 27 since b hold a reference or pointer to a.
```

- Values hold copies
- Pointers don't hold their own copies but rather point to data held by another variable.