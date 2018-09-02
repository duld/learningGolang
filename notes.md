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

---
## Section 7 - Control flow
### Loop - init, condition, post
Most loops take the same shape
* we describe a counter which will increment on each successful loop.
* we write out an expression to test our counter against
* finally we describe how our counter increments on each successful loop

This style of loop structure is a classic for loop. Other loops include the forEach, while, do-while and for-In loops.


### Nested loops
We can nest looping structures inside on another. 

### For Statement
Go's for statement does a lot of heavy lifting when it comes to looping. Go's 'for' statement can simulate multiple styles of looping: while, for, for-in/for-each.

### Standard For loop: for INIT; CONDITION; POST {}
The most basic of the looping syntax is the initialize, condition, post structure; found in many other languages.

### While loop: for {}
if we do not provide any condition to the for loop, it will evaluate to 'true' or, continue looping. If we wish to exit the loop, we would have to use the 'break' keyword.

### Break and Continue
__break:__
  * used to break out of a block of code

__continue:__
  * used to continue into the next iteration of a loop.

### If Statement
If statements in Go behave like in other c style languages. One thing of interest however is that you can have an Initialize statement to go along with the condition statement.
```Go
if y := true; y {
  fmt.Println("Y is indeed true")
}
```
### Switch Statement
Go has a switch statement that is similar to other languages, with one exception. There is no default fallthrough behavior. You must specify that explicitly
```Go
switch {
case false:
  // code wont run
case (2 == 4):
  fallthrough
  // code wont run
case (3 == 3):
  // code should run!
case true:
  // code wont run! Becuase we didn't specify the above case to fallthrough
default:
  // will catch any case that isn't matched.
}
```

### Arrays
Arrays hold a single type, and are a fixed length. Arrays like in other languages are zero based when indexing. However arrays in go have some differences
  * Arrays are values. Assigning one array to another copies all the elements.
  * In particular, if you pass an array to a function, it will receive a copy of the array, __not a pointer to it__.
  * The size of an array is part of its type. The types __[10]int__ and __[20]int__ are distinct.

### Slice
A slice is made up of three parts.
* A pointer to an __underlying array__
* the __length__ of the slice
* the __capacity__ of the slice 
 
The length of a slice represents the total indexable range of the slice. We can access and set values within that range. The capacity of a slice is the total size of the underlying array allocated for the slice. To add items to the slice past the length, we must call the built-in __append()__ function, passing in the slice we wish to append to. We can append a series of comma seperated values or an entire slice of a valid type using the __...__ operator.

### Map
A map is a key value store, where a unique key is used when looking up the value. Think Dictionary. The order of values in a map is not stored, and they can be returned in any order.

### Struct
A struct is a data structure that aggregates values of different types together.

### Embedded Structs
We can embed different types inside of our own struct. By including a type as a nameless parameter within another type, the exported parameters and methods defined on the embedded type are accessible through the ebedding type. The compiler deides on this by using a technique called __promotion__.

* __promotion:__
  * The exported properties and methods of the embedded type are promoted to the embedding type.