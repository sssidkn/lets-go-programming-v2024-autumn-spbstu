package read

import (
	"bytes"
	"encoding/xml"
	"os"
	"strings"

	"github.com/artem6554/task-3/structs"
	"golang.org/x/net/html/charset"
)

func ParseXML(inputFile string) structs.ValCurs {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	data = []byte(strings.ReplaceAll(string(data), ",", "."))

	var currencies structs.ValCurs
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&currencies)
	if err != nil {
		panic(err)
	}
	return currencies
}
