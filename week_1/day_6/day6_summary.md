# Goals

- Understand and implement interfaces in Go.
- Practice polymorphism by writing functions that operate on interfaces.
- Explore the flexibility of the empty interface (`interface{}`) and type assertions.
- Learn how to compose multiple interfaces for reusable behavior.

---

# Notes

## Interfaces:
- Interfaces define a set of methods that a type must implement.
- Example:
  ```go
  type Shape interface {
      Area() float64
      Perimeter() float64
  }
  ```
  - Any type that implements these methods satisfies the `Shape` interface.

## Polymorphism:
- Functions can operate on interfaces, allowing different types to share a common behavior.
- Example:
  ```go
  func printShapeDetails(s Shape) {
      fmt.Println("Area:", s.Area())
      fmt.Println("Perimeter:", s.Perimeter())
  }
  ```
  - Both `Circle` and `Rectangle` implement `Shape`, so they can be passed to `printShapeDetails`.

## Empty Interface (`interface{}`):
- The empty interface can hold values of any type, making it a flexible container for heterogeneous data.
- Example:
  ```go
  func printAny(value interface{}) {
      fmt.Println("Value:", value, ", Type:", reflect.TypeOf(value))
  }
  ```

## Type Assertions:
- Use type assertions to extract specific types or check if a value implements an interface.
- Syntax:
  ```go
  if shape, ok := value.(Shape); ok {
      fmt.Println("Area:", shape.Area())
  }
  ```

## Interface Composition:
- Interfaces can be composed to define more complex behaviors.
- Example:
  ```go
  type Flyer interface {
      Fly() string
  }
  type Swimmer interface {
      Swim() string
  }
  type Duck struct {
      Name string
  }
  ```

---
