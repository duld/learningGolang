# CIT 90

## Things to consider when developing in go

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

### HW

Variables, values & type
Exercises - Ninja Level 1

## __Variales Values and Type__

### __Playground__

We have access to an online go playground @ play.golang.org where we can test code and share the code by URL.

To output text to the user in the console/terminal we can use the go package 'fmt' or format.

---

### __Hello World__

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

keywords

* these are words that ar reserved for use by Go.

operator

* a character or short character sequence that represents an action.
* the values being operated on.

statement

* the smallest standalone element of an imperative programming language that expresses some action.

expression

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

```Golang
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

#### UTF-8

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
  * Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in composite literals of the struct.
    * whats this mean?

### Anonymous Structs

We can create anonymous structs in go, by that I mean that we can define adata structure (struct) that is not named, it has no identifer. This would besimilar to how in JS or Python we can create Object/Dictionary literalswithout creating a Class to represent our data structure.

```JavaScript
// JavaScript
let myData = { // this object isn't given a type name, but is still valid
  keys : {},
  dataList : [],
  _rev : 23
}
```

```Golang
// Golang
myData := struct{ // this describes the shape of our data structure
  keys map[string]string
  dataList []string,
  _rev int
}{ // now lets initialize it
  keys : make(map[string]string),
  dataList: make([]string, 10),
  _rev: 23,
}

```

In the above example we are achieving very similar resulst. Go is a little more verbose in this example, but only because we must explicitly declare the types of the data to be housed. The dynamic nature of JS allows for more flexiblity but at the cost of storage and possibly performance. ./golang-apologist

## Functions

### Function Syntax

func (r receiver) identifier(parameters) (returns) {...}

* __func:__ the go keyword used to declare a function
* __receiver:__ who does this function belong to
* __identifier:__ what is the name we will use to reference this function
* __parameters:__ what arguments do we need to pass in, when calling the function
* __returns:__ does the function return value(s), if so how many and of what type?

### Variadic Parameter

the variadic parameter allows a function to accept any number of valid arguments. We can also use the variadic parameter to apply the values of a slice as individual arguments in a function call. This behavior is similar to the __spread__ operator in JavaScript.

```Golang
var nums []interface{}
nums = append(nums, 0, 1, 2)

fmt.Println(nums) // is the same as
fmt.Println(nums[0], nums[1], nums[2])
```

### Defer

defer can be used to execute a function after the enclosing function has finished execution. If we defer a multiple functions then the order of defered functions will be last in first out (LIFO). This can be usefull if we need to execute some kind of cleanup code at the end of our function, closing any IO or open connections etc.

### Methods

we can attatch methods to types in go. A method is just a function that is associated with a type, and as a result we can call the method from a variable of that type. There are two types of Methods: Methods that can manipulate the stored values/fields of the type and methods that only have a copy of that value. Any changes that the latter attempts will not be stored on the caller.

If we want to mutate the data on the caller, then we must use Pointers

```Golang
type Person struct {
  first, last string
  age int
  weight float64
}

func (p *Person) WorkOut() { // this function can modify the fields of Person
  p.weight = p.weight - p.weight * 0.01
  fmt.Println("Keep it up!")
}

func (p Person) GrowOld() { // this function will not modify the fields of Person
  p.age = p.age * 5
  fmt.Println("You should be dead!")
}
```

### Pointers

Pointers reference a location in memeory where a value is stored rather than the value itself. What the pointer is allowed to 'point' to is determined by the type of the pointer. The type of the pointer is important because it informs the computer about the size and type of the value to be read. By using pointers we can update the values stored at the memory location. When using pointers we have two operators to keep in mind

#### * - Pointer

The asterisk character followed by the type of the stored value represents a pointer. If we want to create a variable that is of __type pointer to type__ we need to add an asterisk to the beginning of the variable type declaration

```Golang
var n int // this can hold values of type int
var x *int // this can hold a pointer to an int
```

#### * - Dereference Operator

The asterisk also serves double duty as the dereference operator when working with pointers. Dereferencing a pointer gives us access to the value the pointer points to.

```Golang
n := 33
np := &n // we use the & to get the memory address of n
y := *np // (*) will dereference the pointer 'np' and access the value it points to 'n'.
```

#### & - Address of Operator

The & is used to find the address of a variable. If we use a the & on a pointer we get a pointer to a pointer of some type.

```Golang
n := 45
np := &n // results in a pointer of int *int
npp := &np // results in a pointer to a pointer of int **int
```

### When To Use Pointers

Pointers come in handy when you need to pass a large object to a function in Go. Everything in Go is passed by value, even collection and aggregate types. This behavior is even present when working with methods in Go. We must still explicitly use pointers on our reciever to allow us to mutate field values.

### Interfaces

Like a struct an interface is created using the __type__ keyword, followed by a name and the keyword __interface__. But instead of defining fields, we define a "method set". A method set is a list of methods that a type must have in order to "implement" the interface.

[interfaces in go](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)

### Method sets

A method set is a list of methods that a type must have in order to implement the interface. An Interface allows for a type of polymorphism in Go.

A method set for an interface and the methods on a type are not the same thing. Methods associated with a type (where a reciever is specified) can be called on a pointer to the type or on the type itself, without conflict. If however, a method set has a specific reciever format, then the type must follow the reciever format for that method in order for the type to be an implementation of the interface.

What that means essentially:

```Golang
type People interface {
  Greet()
}

type Person struct {
  First, Last string
}

func (p *Person) Greet(){
  fmt.Printf("Hello, my name is, %v %v. It is a pleasure to make your acquaintence.", p.First, p.Last)
}

jm := Person{"J", "M"}

xpeople := []People{jm}
xpeople[0].Greet() // will not run
/*
This example wont run because the person type's Greet method requires that we use a  pointer to person, not a person literal. YES it does implement the People interface, but no it wont be able to run because we are attempting to pass a literal to a method that requires a pointer to type Person.
*/
```

### Concurrency pt1

Making progress on more than one task simultaneously is known as concurrency. Go has rich support for concurrency using goroutines and channels.

## JSON

### Marshal

the Marshal function is used to convert data in Go into JSON format. The resulting data is a slice of bytes.

### Unmarshal

To convert JSON back to Go, we need to pass valid JSON represented as a slice of bytes. The resulting Data structure needs to be either a struct that matches the JSON data or an interface{}. Important, we must provide a memory address to where the unmarshaled JSON is going to be stored as our interface{}/struct argument.

### Writer Interface

The Writer interface is simple but widely used interface in Go. The interface defines one method 'Write' which takes a slice of bytes and returns the number of bytes written, and returns an error if if the write was stopped early due to a problem.

### Sort

The sort package contains functions to sort a collection of primitive types. If the user defined type implements three methods it can be sorted.

* Len() int
* Less(i, j int) bool
* Swap(i, j int)

[Person struct - example](https://play.golang.org/p/JiudEdQMoQl)

### Bcrypt

bcrypt is an industry standard hashing algorithm for protecting sensitive data. Go has a library that implements a version of bcrypt.

### Concurrency pt2

Concurrency is a way to build things. It is the composition of independantly executing processes. Concurrency is about __dealing__ with a lot of things at once and Parallelism is about __doing__ a lot of things at once. These two things are similar but not the same.

Concurrency is about structuring things so that maybe they can be done in parallel, but parallelism is not the goal. Having a good structure is the goal.

### Wait Groups

sync.WaitGroup
WaitGroup.Add()
WaitGroup.Done()

### Race Condition

A race condition is an undesirable situation that occurs when a device or system attempts to perform two or more operations a the same time, but because of the nature of the device or system, the operations must be done in the proper sequence to be done correctly. It becomes a bug when the events do not happen in the order the programmer intended.

### Mutex

Mutual Exclusion or mutex, a Golang provides a data structure that can lock chunks of code from being manipulated by other go routines.

### Atomic

Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

These functions require great care to be used correctly. Except for special, low-level applications, synchronization is better done with channels or the facilities of the sync package. Share memory by communicating; don't communicate by sharing memory.

Atomic allows us to precisely read and write to a location in memory.

## Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

By default sends and receives block until both the sender and reciver are ready. Receivers always block until there is data to recieve. If the channel is unbuffered, the sender blocks until the receiver has received the value. If the channel has a buffer, the sender bocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.

### Comma Ok idiom

Sometimes you need to distinguish a smissing entry from a zero value. Is there an entry for "UTC" or is that 0, because it's not in th emap at all? You can descriminate with a form of multiple assignment.

```Golang
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

For obvious reasons this is called the "comma ok" idiom. In this example, if tz is present, seconds will be set appropriately and ok will be true; if not, seconds will be set to zero and ok will be false. Here's a function that puts it together with a nice error report.

```Golang
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

### Fan Out

Is the idea of spreading work out across multiple functions concurrently. For example, you can spin up 'n' number of functions using the 'go' keyword to put the function on a seperate process.

### Fan In

Fan in is the idea of doing a bunch of work at the same time, and then aggregating the results of that work at a single source.

### Context

[Notes from Golang UK Conference 2016 - A Beginner's Guide to Context](https://www.youtube.com/watch?v=r4Mlm6qEWRs)

#### Problem

* In Go servers, each new request spawns it's own goroutine
* Goroutines don't have any 'thread local' state
* Your code is responsible for things like cancellation, time outs and data

#### Solution

* The __context__ package provides a standard way to solve the problems of managing state during a request.

#### __context__ addresses

* Request scoped data
* Cancellation, Deadlines & Timeouts
* It is safe for concurrent use

__Ref__: Cancellation, Context, and Pluming by Sameer Ajmani(GothamGo 2014)

A context should be concerned with holding relevant state information regarding the request. This includes headers, session data, security credentials.

#### A __context__ is made up four parts

* Done() __<-chan struct{}__
  * returns a channel that is CLOSED when this context is canceled or times out.
* Err() __error__
  * Err indicates why this context was canceled, after the Done channel is closed.
* Deadline() (deadline __time.Time__, ok __bool__)
  * Deadline returns the time when this Context will be canceled, if any.
* Value(key __interface{}__)
  * Value returns the value associated with key or nil if none.

## Errors

### Checking Errors

Functions in go can have multiple return values. We can return an error from a function when something goes wrong. Generally the right-most return of a function is the error, if an error is indeed returned. We can check if an error was indeed returned after the function call ends.

```Golang
retVal, err := usefulFunc(someValue)
if err != nil { // if an error was returned, the value will not be 'nil'
  fmt.Println(err)
}
```

### Printing & Logging

When we print, using fmt.Println() we direct the output to standard out (os.Stdout). We can have this same behavior when using log aswell. But log by default contians more information than regular printing and is safer when using multiple goroutines.

#### Logging [taken from stackoverflow](https://stackoverflow.com/questions/41389933/when-to-use-log-over-fmt-for-debugging-and-printing-error/41390023)

* The log functions print to stderr by default and can be directed to an arbitrary writer. The fmt.Printf functions prints to stdout.
* The log functions can print timestamp, source code location and other info.
* The log functions and fmt.Printf are both thread safe, but concurrent writes by fmt.Printf above an OS dependent size can be interleaved.

## __log.Fatalln()__

* uses os.Exit() to quit the application, deferred function calls will not run.

## __log.Panicln()__

* deferred functions will run
* we can recover from a 'panic'

__Panic:__ is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

__Recover:__ is a built-in function that regains control of a panicking goroutine. __Recover is only useful inside deferred functions__. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

### Errors with Info

We can use the __errors.New()__ call to add our own information to an error. The method __errors.New()__ returns a function called __Error()__ which returns a string

## Testing

Testing in go requires:

* A file that has '_test' as a suffix
* testing functions in the file that start with the prefix 'Test'
  * these functions must take in a single argument of type __*testing.T__
* run go test to execute your tests.

```Golang
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

## Golint

Go lint is a command line tool that offers linting suggestions on your files. You can pass a directory for the linter to check

## Benchmarking

Benchmarking in go requires:

* A file that has a '_test' as a suffix.
* A function in the file that start with the prefix 'Benchmark'.
  * these functions take in a single argument of type __*testing.B__.
  * Go must call the function __b.N__ number of times.
* Inside the function a forloop must be written as so:
  * __for i := 0; i < b.N; i++ { FunctionCall(Args) }__

```Golang
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```

## Coverage

Coverage has to do testing your code and how much of your code is actually 'covered' by tests. Are you testing all of the functions in your package? Are all branches of execution tested for as well?

### Coverage Commands

* __go test -cover__
  * checks the coverage of tests against the functions in a package.
* __go tool cover -func=coverage.out__
  * automatically sets the cover flag to true.
  * however, the results are saved to a named file.
* __go tool cover -html=coverage.out__
  * when the command is run, a browser window pops up, showing the covered (green), uncovered (red), and uninstrumented (grey) source. Here's a screen dump.

## BET - remember to bet

BET stands for Benchmark, Example and Test your code. Remember when writing documentation and tests for your code to include all three of these actions.

## Benchmark

Check that your code is performant. Is it performant? How sure are you?

## Example

Give an example of how your code is to be executed, this example will be used in your documentation

## Test

* Test that your code behaves in the way you expect it to.
* Tests keep your code honest.
* Test it multiple times with different values.