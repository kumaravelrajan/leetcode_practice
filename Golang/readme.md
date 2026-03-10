# Notes

## Rune vs Strings and their respective lengths
Strings in Go are UTF-8 encoded. This means each character in a string can occupy anywhere between 1 and 4 bytes. 

This is why the below code says that s is of length 3 because technically the character 京 takes 3 bytes in UTF-8 encoding and string in Go is implemented as a read-only array of byte (which is an alias for uint8). Hence, len(s) -> 3 (bytes).

On the other hand, rune is an alias for int32 (4 bytes). A slice of rune is hence just an array of 32 bit elements stacked one after the other. Rune slice does not care about UTF-8 encoding. 

This is why if we have len([]rune{'京'}), this would give length one. 

Hence, a common strategy in Go to find the human-seeable length of a string is to convert the string to a rune slice and then measure the length of the rune slice. 

```go
s := "京"
len(s) // 3
r := []rune(s)
len(r) // 1
```

Another common way to find the rune count in strings and not the underlying byte count is by using the utf8.RuneCountInString(s) function in the utf8 package. 

```go
s := "京"
len(s) //3
len(utf8.RuneCountInString(s)) //1
```

### Why this distinction is important?
Imagine you are writing a "Twitter-style" character counter where you want to limit a user to 280 characters, using len(myString) would be a bug. A user typing in Kanji or Emojis would hit the limit much faster than someone typing in English.

In that scenario, you would use utf8.RuneCountInString(s) to get the "human-perceived" length without the memory cost of converting to []rune.

## Arguments in Go are strictly call by value (with some nuances ofc)

1. Simple values

    Values such as int and rune when passed in as arguments are copied into a brand new int / rune in the called function.

1. Pointers
    
    When I pass in a pointer as argument, the pointer address is copied into a brand new pointer variable in the callee but the new pointer points to the same address as the original pointer. This means, any changed made to the underlying structure the pointer points to will be visible even outside the callee.

2. Slices

    When I pass in slices as arguments, the slices is not copied element by element in the callee. Instead the slice template is copied. This template include a pointer pointing to the underlying array, length and capacity. Hence, when in the callee the copied slice template and the pointer within it (that points to the same location as the original pointer in the caller) is used to change data, this change is also visible in the caller function.

1. Structs

    1. Passing a Struct Directly

        If you pass a struct (e.g., func process(u User)) as an argument, Go copies the entire data structure.

        The Copy: Every field within that struct is copied into a brand-new instance in the callee's stack.

        The Impact: If your struct has 2 fields, it's cheap. If it has 50 fields, it’s a lot of data being copied onto the stack.

        The Scope: Any changes made to the fields inside the function are local only. The original struct in the caller remains untouched.

    1. Passing a Pointer to a Struct

        This is the idiomatic way to handle "large" structs or when you need to "mutate" (change) the original data.

        The Copy: Only the memory address (usually 8 bytes on a 64-bit system) is copied.

        The Impact: It is highly efficient regardless of how many fields the struct has.

        The Scope: Changes made to the fields (e.g., user.Name = "Alice") modify the original data because the pointer points to the original memory address.

    1. A Note on "Embedded" Types

        This is where it gets interesting. If your struct contains a slice or a map, and you pass that struct by value:

        The struct header is copied (pass-by-value).

        The slice/map header inside that struct is also copied.

        However, since the copied slice header still contains a pointer to the same underlying array, changing an element in that slice will affect the original struct, even though you passed the struct by value!