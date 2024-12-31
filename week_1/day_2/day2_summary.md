# Goals

- Learn about arrays and slices
- Learn common operations for arrays and slices
- More practice with loops and parsing standard input
- Use more packages for working with arrays and slices (e.g., sorting)

## Notes

- **Arrays vs Slices**:
  - Arrays have a fixed size, while slices can grow beyond their original capacity. 
  - Once the capacity is exceeded, the slice loses connection with its underlying array.
  - To maintain the connection to the backing array, check the slice's capacity before appending.
  
- **Passing Arrays to Functions**:
  - When passing arrays, a copy of the array is made.
  - To pass the array by reference, use a pointer (`&`).
  
- **`append` Function**:
  - Built-in function in Go, different from Python's `append` (which is specific to iterable data structures).
  - Can append a slice to another slice but use `...` to expand the slice before appending.
  - Always check the slice's capacity to avoid errors or breaking the connection to the underlying array.
  
- **Sorting with `sort` Package**:
  - Go provides various sorting functions. For example:
    - `sort.Ints` sorts integers in ascending order.
    - `sort.Strings` sorts strings alphabetically.
  - If you prefer not to use an anonymous function, functions like `sort.Ints` or `sort.Strings` can be used directly.
  
- **Looping through Arrays/Slices**:
  - Use `for idx, v := range arr` to iterate, where `idx` is the index and `v` is the value of each element.
  - Note: `range` returns both the index and the value.

- **Input Parsing with `bufio.NewScanner`**:
  - Use `bufio.NewScanner` and `os.Stdin` to read user input.