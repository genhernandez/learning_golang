# Goals

- Learn to work with multidimensional slices
- Practice creating and iterating over nested slices
- Strengthen debugging skills by identifying and resolving commmon errors
- Explore additional slice manipulation techniques, including reversing and reshaping slices
- Explore creating custom data structures

## Notes 
- We can use `fmt` package to create strings that reference other variables
    - was mostly using `fmt` for printing to console
- Since we are often modifying the slice that is passed in, it can be redundant to return that same slice
    - do not return anything and let caller use updated data structure OR explicitly create a copy and return that
- `make` is useful when you know the max capacity/length of a slice; make sure to use it to avoid unnecessary memory reallocations
