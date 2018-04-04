package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)

	var dat interface{}
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}

	ar, ok := dat.([]interface{})
	if ok && len(ar) == 0 {
		os.Exit(0)
	} else {
		output, _ := json.Marshal(dat)
		fmt.Println(string(output))
		os.Exit(1)
	}
}
