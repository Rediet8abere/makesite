package main
import (
	"fmt"
  // "io/ioutil"
	"flag"
)



// Add a new flag to your command named file.
// This flag represents the name of any .txt
// file in the same directory as your program.
// Run ./makesite --file=latest-post.txt to test.
func main() {
	var file *string
	file = flag.String("file", "defaultValue", ".txt")
	flag.Parse()
	fmt.Println("file:", *file)
	// fileContents, err := ioutil.ReadFile("first-post.txt")
	// if err != nil {
	// 		// A common use of `panic` is to abort if a function returns an error
	// 		// value that we don’t know how to (or want to) handle. This example
	// 		// panics if we get an unexpected error when creating a new file.
	// 		panic(err)
	// }
	// fmt.Print(string(fileContents))
	//
	// bytesToWrite := []byte("hello\ngo\n")
  // err2 := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
  // if err2 != nil {
  //     panic(err)
  // }

	// examplePtr := flag.String("example", "defaultValue", " Help text.")
	// fmt.Println(examplePtr)
	// flag.Parse()


}
