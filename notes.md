# CIT 90
## Week 1
### Course Outlines
https://goo.gl/K5GccY
https://goo.gl/

#### Things to consider when developing in go.
Concurrency vs parallelism

### Go Workspace
Go is dependent on having a 'go workspace'. You need three folders: 'bin', 'pkg', and 'src'.

### Terminal Commands
__pwd__: Print working Directory
__mkdir__: Make a folder at the current path with given name
__cd__: change directory to specified name, or up a folder

### Hash Function
Maps any arbitrary size of data to a fixed size. Regardless of how much data is provided. The same data will always result in the same hash value. If any binary bit of data should change somewhre in the file, it should result in a different hash.

### Computing a hash on Windows
To compute a hash on windows, we can use a PowerShell Commandlet
__Get-FileHash__: 
  * __-Path__: the path to the file hash to compute
  * __-Algorithm__: the hashing algorithm to use when computing the hash. Default is SHA256.


## Environment Variables
__Using Powershell__: 
  * Get-Childiten env:
  * dir env:

### HW:
Variables, values & type
Exercises - Ninja Level 1

## __Variales Values and Type__
### __Playground__
We have access to an online go playground @ play.golang.org where we can test code and share the code by URL.

To output text to the user in the console/terminal we can use the go package 'fmt' or format.

---
###  __Hello World__
Go files need to have a package declaration at the top of the file. A go program needs to have a file with package main and a function named main.

We can import other go packages using the import keyword.

In go, there is no 'public' or 'private'. There is 'exported', 'not exported', 'visibile', 'not visible'. This is to avoid any confusion with other languages.

#### Throwing away return values in go
the underscore ( _ ) value in go alows us to capture a return from a function and to signal to go, that the we do not need to hold onto the value. This character is called the __blank identifier__.


#### Go documentation - triple dots (...)
The triple dots in the go documentation signals that the function can take any number of arguments.

---
### Short Declaration Operator
syntax in go to concisely declare and assign a value to a variable. The short declaration operator can only be used inside of a code block.

__keywords__
  * these are words that ar reserved for use by Go.
__operator__
  * a character or short character sequence that represents an action.
  * the values being operated on.
__statement__
  * the smallest standalone element of an imperative programming language that expresses some action.
__expression__
  * a combination of one or more explicit values, operators and functions that the programming language interprets to produce another value.

---
### The Var Keyword
If you want to have a variable that has package level scope, you can declare and initialize the variable outside of any function using the __var__ keyword.

---
### Exploring Type
Go is a statically typed language. When you declare a variable, the variable will only be able to hold values of the declared type.

The syntax is as follows
```Go
var myVar string = "Hello"
```

#### Primitive Data Types
either a basic type or built-in type. These are types that come with the programming language. 

#### Composite Data Types
A data type that is made up of other data types. The act of constructing a composite type is known as composition.

---
### Zero Value
every type of go has a 'zero value'. If we have declared but not assigned an initial value, go will set one for us.

---
### The fmt package
* __Println:__ Outputs a string to the terminal and appends a newline
* __Printf:__ Outputs a formatted string to the terminal, similar to printf in other languages.
* __Print:__ Outputs a sting to the terminal.
* __Sprint:__ Formats any number of inputs and returns the resulting string.
* __Sprintf:__ Formats any number of inputs using supplied format speccifier and returns the string.
* __Sprintln:__ Formats any number of inputs, appends a newline to the result and returns the string.

---
### Creating Your Own Types
Go is all about type. Meaning that once a type is defined, it cannot be changed. When we create our own types in Go, we give our types a 'type name' and specify and underlying data type. We cannot assign values that are stored as a different type in our created type, EVEN IF they share the same underlying type. To clarify:
```go
var a int
type myDataType int
var b myDataType

func main() {
  a = 34 // works
  b = 21 // no problem
  b = a // Doesn't work. variable 'a' is of type "int", not of type "myDataType"
}
```

---
### Conversion, Not Casting
We can convert between types in Go if the __underlying types__ are compatible or if Go has a method for converting between the two types. For reference see the Go Spec on [Conversions](https://golang.org/ref/spec#Conversions).

---
### Strings

#### Unicode
An effort to create a single character set that included every reasonable writing system on the planet. In unicode, a letter maps to something called a 'code point'. A unicode code point looks like so U+0639, each code point is made up of two bytes and represented in hexidecimal.

#### Unicode Encodings
__UTF-8__:
  * a variable width charater encoding
  * capable of encode over a million valid points using 1-4 bytes.
  * Only uses 1 byte for characters that don't go above 127, ANSI.

### Numeric systems
__Decimal Base 10:__ is how humans like to count: 0-9
__Binary Base 2:__ is how computers store values: 0-1
__Hexidecimal Base 16:__ is a condesed representation that is compatible with Binary.

### Constants
the __const__ keyword in go allows us to declare a variable as immutable. We cannot change its value after it is set.

### Iota
Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. Its value is the index of the respective ConstSpec in that constant declaration, starting at zero.

__Iota will:__
  * only work inside of a constant declaraction
  * self increment, starting at 0
  * reset to zero if used in another const declaration.