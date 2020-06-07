package main
import (
	"fmt"
	"html/template"
  "io/ioutil"
	// "bytes"
	// "encoding/json"
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
	fmt.Println(content)

	t := template.Must(template.ParseFiles("template.tmpl"))

	// reqBodyBytes := new(bytes.Buffer)
	// json.NewEncoder(reqBodyBytes).Encode(content)
  t.Execute(w, content)

	// fmt.Println(reqBodyBytes.Bytes())
	// reqBodyBytes.Bytes()
	stringToWrite := string(cont)
  err2 := ioutil.WriteFile("first-post.html", stringToWrite, 0644)
  if err2 != nil {
      panic(err)
  }

}
