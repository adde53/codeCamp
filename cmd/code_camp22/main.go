package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Identification struct {
	ID    string
	Phone int64
	Email string
}

func main() {
	/*var jsonText = []byte(`[
        {"ID": "ID1", "Phone": 0, "Email": "email@email.com"}
    ]`)*/

	// define slice of Identification
	var idents []Identification

	// Unmarshall it
	/*if err := json.Unmarshal([]byte(jsonText), &idents); err != nil {
		log.Println(err)
	}*/

	idents = append(idents, Identification{ID: "ID2", Phone: 15555555555, Email: "Email"})

	// iterating it
	for _, v := range idents {
		fmt.Println(v)
	}
	fmt.Println()

	// now Marshal it
	result, err := json.Marshal(idents)
	if err != nil {
		log.Println(err)
	}

	// now result has your targeted JSON structure
	fmt.Println(string(result))
}