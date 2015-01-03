// Package wikipedia is a simple Wikipedia library for Go.
package wikipedia

import (
	"encoding/xml"
	"io"
	"os"
	"time"
)

// Parser for MediaWiki XML dump files.
type Parser struct {
	decoder *xml.Decoder
}

// Contributor data structure.
type Contributor struct {
	Username string `xml:"username"`
	ID       int    `xml:"id"`
	IP       string `xml:"ip"`
}

// Revision data structure.
type Revision struct {
	ID          int         `xml:"id"`
	ParentID    int         `xml:"parentid"`
	Timestamp   time.Time   `xml:"timestamp"`
	Contributor Contributor `xml:"contributor"`
	Comment     string      `xml:"comment"`
	Model       string      `xml:"model"`
	Format      string      `xml:"format"`
	Text        string      `xml:"text"`
	SHA1        string      `xml:"sha1"`
}

// Page data structure.
type Page struct {
	Title     string     `xml:"title"`
	Namespace int        `xml:"ns"`
	ID        int        `xml:"id"`
	Revisions []Revision `xml:"revision"`
}

// NewParser creates a new parser from an io.Reader.
func NewParser(reader io.Reader) (parser *Parser, err error) {
	parser = &Parser{
		decoder: xml.NewDecoder(reader),
	}

	return
}

// NewParserFromFile creates a new parser from a file path.
func NewParserFromFile(path string) (parser *Parser, err error) {
	reader, err := os.Open(path)

	if err != nil {
		return
	}

	return NewParser(reader)
}

// NextPage gets the next page in the XML dump.
func (parser *Parser) NextPage() (page *Page, err error) {
	var token xml.Token

	for {
		token, err = parser.decoder.Token()

		if err != nil || token == nil {
			return
		}

		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "page" {
				err = parser.decoder.DecodeElement(&page, &tokenType)

				return
			}
		}
	}

	return
}
