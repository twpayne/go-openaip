// Package openaip decodes http://www.openaip.net/ airspace files.
package openaip

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// An Alt represents an altitude.
type Alt struct {
	Unit  string
	Value float64
}

// An AltLimit represents an altitude limit.
type AltLimit struct {
	Reference string `xml:"REFERENCE,attr"`
	Value     Alt    `xml:"ALT"`
}

// A Polygon represents a two-dimensional polygon.
type Polygon struct {
	Coords [][]float64 // Longitude, latitude
}

// An Airspace represents a single airspace.
type Airspace struct {
	Category       string    `xml:"CATEGORY,attr"`
	Version        string    `xml:"VERSION"`
	ID             int       `xml:"ID"`
	Country        string    `xml:"COUNTRY"`
	Name           string    `xml:"NAME"`
	AltLimitTop    AltLimit  `xml:"ALTLIMIT_TOP"`
	AltLimitBottom AltLimit  `xml:"ALTLIMIT_BOTTOM"`
	Polygons       []Polygon `xml:"GEOMETRY>POLYGON"`
}

// An OpenAIP represents version information and zero or more Airspaces.
type OpenAIP struct {
	Version    string     `xml:"VERSION,attr"`
	DataFormat string     `xml:"DATAFORMAT,attr"`
	Airspaces  []Airspace `xml:"AIRSPACES>ASP"`
}

func (a *Alt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "UNIT" {
			a.Unit = attr.Value
		}
	}
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return err
	}
	var err error
	a.Value, err = strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	return nil
}

func (p *Polygon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return err
	}
	css := strings.Split(value, ",")
	if len(css) < 3 {
		return fmt.Errorf("expected at least 3 coordinates, got %d: %#v", len(css), css)
	}
	coords := make([][]float64, len(css))
	for i, cs := range css {
		fields := strings.Fields(cs)
		if len(fields) != 2 {
			return fmt.Errorf("expected two values, got %d: %v", len(fields), cs)
		}
		coords[i] = make([]float64, 2)
		var err error
		coords[i][0], err = strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return err
		}
		coords[i][1], err = strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return err
		}
	}
	p.Coords = append(p.Coords, coords...)
	return nil
}

// Read reads an OpenAIP from r.
func Read(r io.Reader) (*OpenAIP, error) {
	var oa OpenAIP
	d := xml.NewDecoder(r)
	if err := d.Decode(&oa); err != nil {
		return nil, err
	}
	if oa.DataFormat != "1.1" {
		return nil, fmt.Errorf("unsupported data format: %v", oa.DataFormat)
	}
	return &oa, nil
}
