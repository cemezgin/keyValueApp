package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	key_value "github.com/cemezgn/keyValueApp/pkg/key-value"
)

func Read()  {
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var items key_value.Items

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &items)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(items.Items); i++ {

	}


}
