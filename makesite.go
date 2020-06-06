package main
import (
	// "fmt"
  "io/ioutil"
	"html/template"
	"os"
)

type FileCont struct {
	Content string
}

func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
			panic(err)
	}
	// sending data to struct
	content := FileCont{
		Content: string(fileContents),
	}

	t := template.Must(template.ParseFiles("template.tmpl"))
	t.Execute(os.Stdout, content)
}
