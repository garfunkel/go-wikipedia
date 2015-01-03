package wikipedia

import (
	"io"
	"log"
	"testing"
)

func TestParser(t *testing.T) {
	parser, err := NewParserFromFile("tests/test1.xml")

	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err = parser.NextPage()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
