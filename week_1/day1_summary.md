# Day 1 Goals

- Set up developer environment
- Learn basic syntax for declaring various types of variables (strings, booleans, integers, floats)
- Use common packages (fmt, math)
- Explore flow control (`if`, `switch`, `for`)
- Learn how to pass in arguments from CLI

## Day 1 Notes

- Golang (like Java) is more picky about types (compared to Python)
    - can't use a function if arg type doesn't match expected type in function signature
    - integer division resolves to 0, must use floats if you want a precise answer
    - interpolation requires specific types, no automatic casting to string
- Complains more about unused packages/variables (compared to Python)
- Golang style guide follows CamelCase conventions and I am used to snake_case
- DateTime formatting does not follow using `%Y` or `YYYY`, instead uses specific values to represent a date time format.
    - Eg: `YYYY-MM-DD` is `2006-01-02`
- Function signatures expect parameter type and return type
- Naked returns will return all named variables (wild!)
- started to explore flow control, but still have to work on using `switch`
    - gotta remember to use `{}` instead of `:` 
