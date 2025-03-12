package main

import (
	"fmt"
)

func ListLanguagesMap() {
	languagesMap := make(map[string]string, 5)
	languagesMap["Angular"] = "Frontend"
	languagesMap["Node"] = "Backend"
	languagesMap["React"] = "Frontend"
	languagesMap["Go"] = "Backend"
	languagesMap["Python"] = "Backend"
	fmt.Println(languagesMap) // map[Angular:Frontend Go:Backend Node:Backend Python:Backend React:Frontend]
	GoValue, hasGo := languagesMap["Go"]
	fmt.Println(GoValue, hasGo) // Backend
	_, hasJava := languagesMap["Java"]
	fmt.Println(hasJava)              // false
	fmt.Println(languagesMap["Go"])   // Backend
	fmt.Println(languagesMap["Java"]) // ""
	delete(languagesMap, "Java")
	fmt.Println(languagesMap) // map[Angular:Frontend Go:Backend Node:Backend Python:Backend React:Frontend]
	for key, value := range languagesMap {
		fmt.Println(key, value)
	}
	// Angular Frontend
	// Node Backend
	// React Frontend
	// Go Backend
	// Python Backend
	delete(languagesMap, "React")
	fmt.Println(languagesMap) // map[Angular:Frontend Go:Backend Node:Backend Python:Backend]
	languagesMap["Java"] = "Backend"
	fmt.Println(languagesMap) // map[Angular:Frontend Go:Backend Java:Backend Node:Backend Python:Backend]
	languagesMap["Next"] = "Frontend"
	fmt.Println(languagesMap) // map[Angular:Frontend Go:Backend Java:Frontend Node:Backend Python:Backend]
	// Check if the Vue language exists in the map
	if _, ok := languagesMap["Vue"]; ok {
		fmt.Println("Vue exists")
	} else {
		fmt.Println("Vue does not exist")
	}
	// Vue does not exist
}
