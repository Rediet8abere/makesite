package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func (item *Result) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Println(err)
		return err
	}
	item.Data = v[0].(interface{}).([]interface{})[0].([]interface{})[0].(string)
	return nil
}

type Result struct {
	Data interface{}
}

func main() {
	// fmt.Printf("\nEnter text\t: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()

	data, err := ioutil.ReadFile("latest-post.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:")
	fmt.Println(string(data))
	query := string(data)
	// scanner.Text()
	response, err := http.Get("https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=es&dt=t&q=" + url.QueryEscape(query))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var result Result
		if err := json.Unmarshal(bytes, &result); err != nil {
			fmt.Println(err)
		}
		value := result.Data
		fmt.Printf("Translation\t: %s\n", value)
	}
}
