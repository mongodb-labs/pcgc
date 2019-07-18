package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func prettyJSON(obj interface{}) {
	prettyJson, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		er(err)
	}

	fmt.Println(string(prettyJson))
}
