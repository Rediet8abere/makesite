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
	// 5. flag to your command named file. This flag represents the name of any .txt file in the same directory as your program
	file_flag := flag.String("latest-post.html", "defaultValue", ".txt")
	flag.Parse()
	fmt.Println("file:", *file_flag)
}

func main() {
	save("first-post.txt")
}

// Questions: how to write an html file
// 						what is flag and what does it do?
