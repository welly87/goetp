package main

// ***********************
// NOTICE this file was changed beginning in November 2016 by the team maintaining
// https://github.com/go-avro/avro. This notice is required to be here due to the
// terms of the Apache license, see LICENSE for details.
// ***********************

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/avro.v0"
)

type schemas []string

func (i *schemas) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *schemas) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var schema schemas
var output = flag.String("out", "", "Output file name.")

func main() {
	parseAndValidateArgs()

	schemas := []string{
		"Energistics.Datatypes.Version",
		"Energistics.Datatypes.ArrayOfDouble",
		"Energistics.Datatypes.DataValue",
		"Energistics.Datatypes.SupportedProtocol",
		"Energistics.Protocol.Core.RequestSession",
		"Energistics.Protocol.Core.OpenSession"}

	for i, val := range schemas {
		schemas[i] = strings.Replace(val, ".", "/", -1) + ".avsc"
	}

	for i, _ := range schemas {
		result, _ := ioutil.ReadFile(schemas[i])
		fmt.Println(schemas[i])
		schemas[i] = string(result)
	}

	gen := avro.NewCodeGenerator(schemas)

	code, err := gen.Generate()
	checkErr(err)

	createDirs()
	err = ioutil.WriteFile(*output, []byte(code), 0664)
	checkErr(err)
}

func parseAndValidateArgs() {
	flag.Var(&schema, "schema", "Path to avsc schema file.")
	flag.Parse()

	if len(schema) == 0 {
		fmt.Println("At least one --schema flag is required.")
		os.Exit(1)
	}

	if *output == "" {
		fmt.Println("--out flag is required.")
		os.Exit(1)
	}
}

func createDirs() {
	index := strings.LastIndex(*output, "/")
	if index != -1 {
		path := (*output)[:index]
		err := os.MkdirAll(path, 0777)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
