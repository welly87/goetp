package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
	"github.com/stevenle/topsort"
)

type Message struct {
	Namespace string
	Name      string
	Fields    []Field
}

type Field struct {
	Name string
	Type interface{}
}

func run() ([]string, error) {
	searchDir := "Energistics"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}

		return err
	})

	if e != nil {
		println(e.Error())
	}

	return fileList, nil
}

func recurse(json string) {
	jsonparser.ArrayEach([]byte(json), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tp, _ := jsonparser.GetUnsafeString(value, "type")

		fmt.Print(dataType.String() + " ")

		fmt.Println(tp)

		recurse(tp)

	}, "fields")
}

func recurseField(field interface{}) {
	switch v := field.(type) {
	case string:
		switch v {
		case "string":
		case "int":
		case "null":
		case "map":
		case "bytes":
		case "long":
		case "array":
		case "boolean":
		case "float":
		case "double":
		default:
			//fmt.Println(v)

			if !graph.ContainsNode(v) {
				graph.AddNode(v)
			}

			dependencies[v] = true
			// set[v] = true
		}

	case map[string]interface{}:
		for k, v := range v {
			switch vv := v.(type) {
			case string:
				recurseField(v)
			case float64:
				fmt.Println(k, "is float64", vv)
			}
		}
	case []interface{}:
		for _, u := range v {
			recurseField(u)
		}
	case Field:
		recurseField(v.Type)
	}
}

// var set = make(map[string]bool)

var dependencies = make(map[string]bool)

func createDependencies(filePath string) {
	data, _ := ioutil.ReadFile(filePath)

	var m Message

	e := json.Unmarshal(data, &m)

	typeName := m.Namespace + "." + m.Name

	graph.AddNode(typeName)

	if e != nil {
		fmt.Println("error => ", e.Error(), filePath)
	}

	for _, field := range m.Fields {

		dependencies = make(map[string]bool)

		recurseField(field)

		for depend := range dependencies {
			graph.AddEdge(typeName, depend)
		}
	}
}

var graph = topsort.NewGraph()

func main() {

	fileList, _ := run()

	for _, path := range fileList {
		createDependencies(path)
	}

	result, _ := graph.TopSort("Energistics.Protocol.Core.OpenSession")

	for _, value := range result {
		fmt.Println(value)
	}
}
