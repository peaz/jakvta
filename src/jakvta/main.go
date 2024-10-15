package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addKeyValueRecursively(data interface{}, key string, value interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		v[key] = value
		for _, val := range v {
			addKeyValueRecursively(val, key, value)
		}
	case []interface{}:
		for _, val := range v {
			addKeyValueRecursively(val, key, value)
		}
	}
}

func main() {
	inputFile := flag.String("i", "", "input file name and path")
	key := flag.String("k", "", "key to add to each object")
	valueStr := flag.String("v", "", "value to assign to the key (string, number, or boolean)")
	dataType := flag.String("t", "", "optional: specify datatype (bool, int, float, string)")
	flag.Parse()

	if *inputFile == "" || *key == "" || *valueStr == "" {
		fmt.Println("Usage: go run main.go -i <input_file> -k <key> -v <value>")
		flag.PrintDefaults()
		return
	}

	// Read the JSON file
	byteValue, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var result interface{}
	json.Unmarshal(byteValue, &result)

	var value interface{}
	var errValue error

	if *dataType != "" {
		value, errValue = parseValueWithType(*valueStr, *dataType)
	} else {
		value, errValue = parseValueAutomatic(*valueStr)
	}

	if errValue != nil {
		fmt.Println("Error parsing value:", errValue)
		return
	}

	addKeyValueRecursively(result, *key, value)

	// Convert the modified data back to JSON
	modifiedJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the modified JSON back to file
	baseName := strings.TrimSuffix(*inputFile, ".json")
	outputFile := baseName + "-modified.json"
	err = os.WriteFile(outputFile, modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("Successfully added '%s': '%s' to all objects in the JSON. Output written to %s\n", *key, *valueStr, outputFile)
}

func parseValueWithType(valueStr, dataType string) (interface{}, error) {
	switch dataType {
	case "bool":
		return strconv.ParseBool(valueStr)
	case "int":
		return strconv.ParseInt(valueStr, 10, 64)
	case "float":
		return strconv.ParseFloat(valueStr, 64)
	case "string":
		return valueStr, nil
	default:
		return nil, errors.New("invalid datatype specified")
	}
}

func parseValueAutomatic(valueStr string) (interface{}, error) {
	if b, err := strconv.ParseBool(valueStr); err == nil {
		return b, nil
	}
	if i, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return i, nil
	}
	if f, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return f, nil
	}
	return valueStr, nil
}
