# Go
- It is not an object oriented language.
- GO is a pass by value language.

# Conventions 
- Use of **`this` or `self` is not a good** practice. It is technically correct though.

# New Leanings 
## Variables
- Variable declaration outside of function is ok but **definition** is not.
- It is required to be all the defined variables. If a variable has to be defined but won't be used then **used `_` to avoid the compilation error.**

## References
- The slices splitting creates **two new references that still points** to the main slice.
- 