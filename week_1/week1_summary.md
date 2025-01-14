# Week 1: Learning Golang

This README outlines the daily learning goals and accomplishments for the first week of exploring Golang. Each day builds upon the previous day’s concepts to establish a strong foundation for working with Go.

---

## Day 1: Getting Started with Golang

**Goals:**
- Understand Golang’s syntax and basic program structure.
- Set up the Go development environment.

**Accomplishments:**
- Installed Go and configured the workspace.
- Wrote a simple "Hello, World!" program.
- Learned about `package main`, `import`, and the `func main` structure.
- Explored basic data types (`int`, `float`, `string`, `bool`).
- Experimented with basic control structures: `if`, `for`, and `switch`.

---

## Day 2: Working with Functions and Arrays

**Goals:**
- Learn about functions, parameters, and return values.
- Understand arrays and slices, their differences, and use cases.

**Accomplishments:**
- Defined and called custom functions with parameters and return values.
- Explored arrays: fixed size and type constraints.
- Worked with slices: dynamic resizing, appending, and slicing operations.
- Gained an understanding of slice capacity and backing arrays.

---

## Day 3: Exploring Sorting and Filtering

**Goals:**
- Learn to sort and filter slices.
- Understand space optimization when working with slices.
- Deepen familiarity with loops and built-in Go packages.

**Accomplishments:**
- Used the `sort` package to sort slices of integers and strings.
- Implemented custom sorting with `sort.Slice` and `sort.Sort`.
- Created filters to remove specific elements from slices.
- Explored space optimization by initializing slices with `var` and `make`.

---

## Day 4: Advanced Slice Manipulation

**Goals:**
- Learn to work with multidimensional slices.
- Practice creating and iterating over nested slices.
- Strengthen debugging skills by identifying and resolving common errors.

**Accomplishments:**
- Created and manipulated 2D slices.
- Wrote nested loops to process multidimensional data.
- Debugged index out-of-range errors when iterating over slices.
- Practiced using `fmt` for structured debugging and output.

---

## Day 5: Errors and File Handling

**Goals:**
- Learn to work with errors, creating custom ones and handling errors returned by functions
- Learn to work with files: reading, writing, copying

**Accomplishments:**
- Created custom Error 
- Read and wrote to a file
- Debugged common file errors (eg. FileNotFound)
- Practicied using `fmt` for custom error messages
---

## Day 6: Interfaces and Polymorphism in Go

**Goals:**
- Learn how to define and use interfaces
- Understand how interfaces enable polymorphism in Go
- Practice implmenting interfaces with concrete types
- Explore the `empty interface` and type assertions

**Accomplishments:**

1. **Defined and Implemented Interfaces:**
   - Created a `Shape` interface with `Area` and `Perimeter` methods.
   - Implemented `Shape` for `Circle` and `Rectangle`.

2. **Used Polymorphism:**
   - Wrote functions like `printShapeDetails` and `findLargestArea` to work with any `Shape`.

3. **Explored Empty Interface:**
   - Created a `printAny` function to accept values of any type.
   - Used `reflect.TypeOf` to determine the type of input dynamically.

4. **Practiced Type Assertions:**
   - Enhanced `printAny` to detect and print additional details for `Shape` types.

5. **Composed Interfaces:**
   - Created `Flyer` and `Swimmer` interfaces.
   - Implemented both in a `Duck` type.
   - Wrote a `describeAnimals` function to check and describe capabilities dynamically.

---

