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

## Variales Values and Type
#### Playground
We have access to an online go playground @ play.golang.org where we can test code and share the code by URL.

To output text to the user in the console/terminal we can use the go package 'fmt' or format.

####  Hello World
Go files need to have a package declaration at the top of the file. A go program needs to have a file with package main and a function named main.

We can import other go packages using the import keyword.

In go, there is no 'public' or 'private'. There is 'exported', 'not exported', 'visibile', 'not visible'. This is to avoid any confusion with other languages.

#### Throwing away return values in go
the underscore ( _ ) value in go alows us to capture a return from a function and to signal to go, that the we do not need to hold onto the value. This character is called the __blank identifier__.

#### Go documentation - triple dots (...)
The triple dots in the go documentation signals that the function can take any number of arguments.

#### Short Declaration Operator
syntax in go to concisely declare and assign a value to a variable. The short declaration operator can only be used inside of a code block.

__keywords__
  * these are words that ar reserved for use by Go.
__operator__
  * a character or short character sequence that represents an action.
__operand__
  * the values being operated on.
__statement__
  * the smallest standalone element of an imperative programming language that expresses some action.
__expression__
  * a combination of one or more explicit values, operators and functions that the programming language interprets to produce another value.

### The Var Keyword
If you want to have a variable that has package level scope, you can declare and initialize the variable outside of any function using the __var__ keyword.

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

### Zero Value
every type of go has a 'zero value'. If we have declared but not assigned an initial value, go will set one for us.

#### The fmt package
