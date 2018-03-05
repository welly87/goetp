package codegen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

func run(searchDir string) ([]string, error) {
	// searchDir := "Energistics"

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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var etpsortedschemas = make([]string, 0)

func EtpSortedSchemaList() []string {

	fileList, _ := run("Energistics")

	for _, path := range fileList {
		createDependencies(path)
	}

	fileList, _ = run("Energistics/Protocol")

	for _, path := range fileList {

		typeName := strings.Replace(filepath.Dir(path), "\\", ".", -1) + "." + strings.Replace(filepath.Base(path), ".avsc", "", -1)
		result, _ := graph.TopSort(typeName)

		for _, value := range result {
			if !contains(etpsortedschemas, value) {
				etpsortedschemas = append(etpsortedschemas, value)
			}
		}
	}

	return etpsortedschemas
}
