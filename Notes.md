# Go
- It is not an object oriented language.
- GO is a pass by value language.

# Conventions 
- Use of **`this` or `self` is not a good** practice. It is technically correct though.
- Using **Pointers for receiver functions** is a good practice in Go.

# New Leanings 
## Variables
- Variable declaration outside of function is ok but **definition** is not.
- It is required to be all the defined variables. If a variable has to be defined but won't be used then **used `_` to avoid the compilation error.**
- `var` declaration does not do initialization for the variables (checked for `map`.) Instead use `make` for declaration and initialization.
- Anything that starts with CAPITAL letter will be available outside of the package. It can be single variable, type, struct props etc.

## References
- The slices splitting creates **two new references that still points** to the main slice.

## Value and Reference data types
- Value types: int, float, string, bool, struct (Would require pointers for references)
- slices, maps, channels, pointers, functions
- When we create a slice, Go will automatically create two data structures
  - An array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array

# Interfaces
- In Go the types "satisfies" an interface, instead of the types "implements" an interface. Subtle difference.
- The above means a receiver method "() xyz()" in type "A" and "() xyz()" in type "B" doesn't have to be implement an interface "I" which required "xyz()". But the types still satisfies the "requirement" of interface and hence it is valid.