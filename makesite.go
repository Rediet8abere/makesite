package main

// final submission
import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"

	// "bytes"
	// "encoding/json"
	"os"
)

type FileCont struct {
	Title   string
	Content string
}

func save(filename string) {
	// 1. Read in the contents of the provided first-post.txt file.
	fileContents, err := ioutil.ReadFile(filename)

	// 3. Render the contents of first-post.txt using Go Templates and print it to stdout.
	fmt.Println(string(fileContents))
	if err != nil {
		panic(err)
	}

	content := FileCont{
		Title:   "first-post-w",
		Content: string(fileContents),
	}
	t := template.Must(template.ParseFiles("template.tmpl"))
	// filename

	file, err := os.Create(filename[:len(filename)-4] + ".html")
	// 2. Edit the provided HTML template (template.tmpl)to display the contents of first-post.txt
	t.Execute(os.Stdout, content)
	// 4. Write the HTML template to the filesystem to a file. Name it first-post.html
	t.Execute(file, content)

}

func testFlag() {
	// https://stackoverflow.com/questions/1726891/explanation-of-flags-in-go#:~:text=flags%20are%20a%20common%20way,options%20for%20command%2Dline%20programs.&text=Try%20out%20the%20run%20program,automatically%20take%20their%20default%20values.&text=If%20you%20provide%20a%20flag,will%20overwrite%20by%20passed%20one.
	// 5. flag to your command named file. This flag represents the name of any .txt file in the same directory as your program
	fileFlag := flag.String("file", ".txt", "a string")
	dirFlag := flag.String("dir", "search", "a string")
	flag.Parse()
	fmt.Println("file:", *fileFlag)
	fmt.Println("looking for txt file:", *dirFlag)
}

func main() {
	save("first-post.txt")
	save("test1.txt")
	save("test2.txt")
	save("test3.txt")
	testFlag()
}
