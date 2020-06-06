// https://blog.gopheracademy.com/advent-2017/using-go-templates/
package main

import (
	"html/template"
	"os"
	// "fmt"
  // "io/ioutil"
)

type entry struct {
  Name string
  Done bool
}

type ToDo struct {
  User string
  List []entry
}


func main() {

	// Data 
	data := ToDo{
    User: "Rediet",
    List: []entry{
        {Name: "Task 1", Done: false},
        {Name: "Task 2", Done: true},
        {Name: "Task 3", Done: true},
    },
}
	// Parse tempalte and Execute
  t := template.Must(template.ParseFiles("todo.tmpl"))
	t.Execute(os.Stdout, data)
}
