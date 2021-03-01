# go-intro
Repository for learning basics of Go.

## Learning points
- Package with name main will build an executable.
- Use 'go build' to build executable.
- Use 'go run' to build and execute. Does not save executeable file in location.  
- Every package should have a package comment "Package *packagename*". To be used in godoc.
- Use 'go doc *packagename*' to see documentation on a specific package.
- Unused imports are not allowed and will cause an error.
- Unused variables are not allowed and will cause an error.
- Unassigned variables have default values.
- Decides the type for *default* int types depending on the architecture of the computer it is running on.
- Delay a function call to the end of current scope by using *defer*. Useful for logging and file writing.
- Go is a *pass-by-value* language. The value of an argument is passed to a function rather than the argument itself.
