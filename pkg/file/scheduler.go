package file

import (
	"fmt"
	"io/ioutil"
	"time"

	key_value "github.com/cemezgn/keyValueApp/pkg/key-value"
)

func Run(repository *key_value.Repository)  {
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				list, err := repository.List()

				if err != nil {
					return
				}
				_ = ioutil.WriteFile("data.json", list, 0644)
				fmt.Println("Data saved to file.")
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}