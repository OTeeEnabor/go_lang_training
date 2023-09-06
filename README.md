# Go Lang Training for Demysifying Blockchain Bootcamp

Go has one entry point into the application - main. 
Once in the entry point, go will run all other functions/applications it sees. 

two methods tp specify entry point

1* package - main
A folder in Go is considered as a package. 
2* func main
A function named main. 

To execute a file in Go - use go run filename.go

### variable declaration
A variable is anything that has the ability to contain information. It is a placeholder in memory used to reference information. 

long form declaration *(var var_name variable type = value)*

Short form declaration  *(variable := value)* only works within a function. 

### Go Mod Init
go mod init github repository
go mod 

## Day 2
Topics covered
const
functions introduction a

What does go mod tidy do?

## Day 3
functions continued
package imports

## Day 4
Structs
## Day 5
if is used to control the program flow. considered a control structure. 

What is a buffer? computer science.

*VS Code shortcut - ctrl+D - select words

Error handling in Go

make 
new
### Interface
*what is an interface?*

Go interfaces provide method signatures for similar types of objects. 

Interface values can be thought of as a tuple of a value and a concrete type: (value, type). An interface value holds a value of a specific underlying concrete type.

A method call on an interface value executes the method of the same name on its underlying type. 

*How to create an interface?*

Use type keyword, followed by the interface name and interface keyword. Method signatures can be specified inside curly braces. 


type interface_name interface {
    // method signature
}

### Empty Interfaces
- Interfaces with no methods. 
- may hold values of any type.(Every type implements at least zero methods.)
- used by code that handles values of unknown type.

- Go interfaces are implemented implicitly.

### Type assertion 
This provides access to an interface value's underlying concrete value. 
*t := i.(T)*

go routines


# Week 2
Frameworks for GoLang
## Echo
Installed Go lang Air - which is a live reloading development tool.
Air is yet another live-reloading command line utility for developing Go applications.
**PORT** HTTP
Annonymous functions in golang

They do not have a name and run automatically.
