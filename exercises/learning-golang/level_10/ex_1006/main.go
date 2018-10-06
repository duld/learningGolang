/*
write a program that
  puts 100 numbers to a channel
  pull the numbers off the channel and print them
*/
package main

import "fmt"

func main() {
	// create our channel
	c := make(chan int)

	// add 100 numbers onto the channel - blocks
	go sendNums(c)
	// read the numbers off of the channel - unblocks
	//  - print the values
	readNums(c)
	fmt.Println("End of Program!")
}

// func sendNums(c chan int) {
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()
// }

func sendNums(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func readNums(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

/*
What would be considered better style I wonder? explicitly using
the 'go' keyword when calling the function, or nesting that call
inside the function body? Both 'work' just as well.

one however causes blocking, so someone using the function later
on would have to know that it blocks and then add the 'go'
keyword to unblock the function. The other hides it, and lets
spares user the cognitive load of remembering to use 'go'. But
shouldn't someone else 'know' what the function is essentially
doing and then make the necessary steps when using the function?

hmmm... we should probably nest it. Keep things simple.
*/
