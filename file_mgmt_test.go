package main

import (
	"io/ioutil"
	"testing"
)

func TestOpenDataFile(t *testing.T) {

	file, closer, err := openAddressDataFile("./my_address.json")
	ok(t, err)

	if closer != nil {
		defer closer()
	}

	byteResult, _ := ioutil.ReadAll(file)
	result := string(byteResult)

	equals(t, 1574, len(result))
}

func TestCreateCsvFile(t *testing.T) {

	_, closer, err := createCsvFile("./output.csv")
	ok(t, err)

	if closer != nil {
		defer closer()
	}
}
