// Build and use an anonymous func

package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hey look at me, I'm doin the thing.")
		fmt.Println("Don't put your labels on me....")
	}()
}
