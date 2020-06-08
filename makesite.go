package main
import (
	"fmt"
	// "html/template"
  "io/ioutil"
	"bytes"
	"encoding/json"
	// "os"
)

type FileCont struct {
	Content string
}

func save(txt string) {
	fileContents, err := ioutil.ReadFile(txt)
	if err != nil {
			panic(err)
	}
	// fmt.Println(fileContents)
	content := FileCont{
		Content: string(fileContents),
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(content)
	fmt.Println(reqBodyBytes.Bytes())		// fmt.Println(reqBodyBytes.Bytes())
	bytesToWrite := []byte(reqBodyBytes.Bytes())		// reqBodyBytes.Bytes()
  err2 := ioutil.WriteFile(txt + ".html", bytesToWrite, 0644)
	if err2 != nil {
      panic(err)
  }
}
func main() {
	save("first-post.txt")


}
