package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	key_value "github.com/cemezgn/keyValueApp/pkg/key-value"
)

func Read() *key_value.DataStore{

	var items []key_value.Item

	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)

	if err != nil {
		fmt.Println(err)
	}

	itemList := make(map[string]key_value.Item, len(items))

	for _, v := range items {
		itemList[v.Key] = v
	}

	rest, _ := json.Marshal(itemList)
	fmt.Println(bytes.NewBuffer(rest))

	return &key_value.DataStore{
		M:       itemList,
		RWMutex: &sync.RWMutex{},
	}

}
