package main

import (
	"encoding/csv"
	// "fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Geocoder struct {
	URL *url.URL
}

func NewGeocoder() *Geocoder {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"
	u.Path = "imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{"507 S Pinehurst Ave"},
		"City":         []string{""},
		"State":        []string{"Maryland"},
		"ZIP":          []string{"21801"},
		"SingleLine":   []string{""},
		"outFields":    []string{""},
		"maxLocations": []string{"United States"},
		"outSR":        []string{"4326"},
		"searchExtent": []string{""},
		"f":            []string{"json"},
	}
	u.RawQuery = v.Encode()

	return &Geocoder{u}
}

func (g *Geocoder) SetUrlValues(address *InRecord) {

	oldQuery := g.URL.Query()

	oldQuery.Set("Street", address.Address)
	oldQuery.Set("ZIP", address.Zip)

	g.URL.RawQuery = oldQuery.Encode()
}

func (g Geocoder) Geocode() ([]byte, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g Geocoder) GeocodeToCandidates() ([]string, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	unmarshaler := JSONMarshaler{}
	candidates, err := unmarshaler.UnmarshalAddresses(res.Body)

	bestMatch := candidates.GetBestMatchLocation()

	return bestMatch, nil
}

func UnmarshalAndGeocodeInRecords(reader *csv.Reader) (outRecords [][]string, err error) {

	eof := false
	for lino := 1; !eof; lino++ {
		line, err := reader.Read()
		if err == io.EOF {
			err = nil
			eof = true
			return outRecords, nil
		} else if err != nil {
			return nil, err
		}

		parsedLine, err := ParseAndGeocodeInRecord(line)
		outRecords = append(outRecords, parsedLine)
	}

	return outRecords, nil
}

func ParseAndGeocodeInRecord(line []string) ([]string, error) {

	inRecord := &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	gc := NewGeocoder()
	gc.SetUrlValues(inRecord)

	xyVals, err := gc.GeocodeToCandidates()

	if err != nil {
		return nil, err
	}

	line = append(line, xyVals...)
	return line, nil
}
