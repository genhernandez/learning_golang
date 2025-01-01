# Goals

- Continue learning about arrays and slicing
    - understanding more `Sort` use cases and filtering

## Notes

- `sort.Reverse` and `sort.StringSlice` return a new `sort.Interface` that has a sorting operation associated with it, in the case of `sort.Reverse` the operation indicates to sort the data structure in reverse order
    - in order to get back the data structure in the desired format, you pass this interface to `sort.Sort` and it will apply the operation and return the sorted data structure
- like `sort.Int` and `sort.String`, `sort.StringSlice` will default to sort in alphabetical/ascending order
- can optimize space when initializing an empty slice
    - eg: `var result [] int`:
        - a zero values slice is created with no allocated array
        - memory allocation happens only when elements are appended
        - avoids creating unused underlying arrays
    - vs: `slice := []int{}`:
        - a slice is explicitly allocated with a backing array
        - the capacity might be unnecessarily large at times
        - appending works fine but might allocate more memory initially than necessary
- can interpolate values in string using `fmt.Println`
    - eg: `fmt.Println("Value of a is", a," and value of b is", b, ".")
        > Value of a is 6 and value of b is 3.
    - don't need to have space at end of strings before variable is interpolated, happens automatically
- can use `make` to create slices, maps, and channels that have a specified length and capacity
    - this is useful for when you know the max capacity ahead of time or want to set a limit
    - using this syntax means that enough memory is pre-allocated upfront, helps avoid memory reallocations that can happen when we append and exceed capacity
    - lack of reallocations means better performance 
- like python, golang also supports working with 2d slices/arrays
    - to create: `nestedSlice := [][]int{{1,2,3}, {1,2,3}, {1,2,3}}`
    - use nested loops to iterate through them
    ```
    for i, row := range nestedSlice {
        for j, val := range row {
            fmt.Printf("Element [%d][%d] = %d\n", i, j, val)
        }
    }
    ```
    - access elements using `nestedSlice[i][j]`
    - appending to a sub slice
        ```
        nestedSlice[0] = append(nestedSlice[0], 10)
        ```
