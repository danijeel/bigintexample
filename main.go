package main

import (
	"encoding/json"
	"fmt"

	"github.com/danijeel/bigintexample/openapi"
)

// openapi generated with
// openapi-generator-cli generate -g  go -i openapi.yml --package-name openapi -o openapi

func main() {

	jsonSrc := []byte(`{
		"id": 162259276829213363391578010288126,
		"name": "test"
	}`)

	var myJson openapi.Category
	json.Unmarshal(jsonSrc, &myJson)

	fmt.Printf("%d\n", myJson.Id)
	// output: 824634459072%
	// expects: 162259276829213363391578010288126
}
