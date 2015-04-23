package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestJSONDecoding(t *testing.T) {

	jsonStream := `{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612517766347352,"y":38.352004138620323},"score":93.200000000000003,"attributes":{}}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var a Address
		if err := dec.Decode(&a); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		equals(t, "main.Address", reflect.TypeOf(a).String())
	}
}

func TestJSONArrayDecoding(t *testing.T) {

	jsonStream := `{"candidates":[{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612517766347352,"y":38.352004138620323},"score":93.200000000000003,"attributes":{}},{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612229940922532,"y":38.353295843541105},"score":93.200000000000003,"attributes":{}},{"address":"507 S PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.61248740139186,"y":38.352168562340495},"score":100,"attributes":{}},{"address":"507 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612448720360234,"y":38.351815864580452},"score":100,"attributes":{}},{"address":"507 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612391256749248,"y":38.35311332552839},"score":90.870000000000005,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612757556175097,"y":38.351796024941883},"score":79,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612397010945358,"y":38.351714488143401},"score":79,"attributes":{}},{"address":"508 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.61254859196886,"y":38.353055632076831},"score":69.870000000000005,"attributes":{}},{"address":"505 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612366186770942,"y":38.351796905938166},"score":68.5,"attributes":{}},{"address":"521 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.614312182687073,"y":38.352233582879819},"score":68.489999999999995,"attributes":{}}]}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var c Candidates
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		equals(t, "main.Candidates", reflect.TypeOf(c).String())
		equals(t, 10, len(c.Candidates))
	}
}

func TestJSONWkid(t *testing.T) {

	jsonStream := `{"wkid":4326,"latestWkid":4326}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var c Wkid
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		equals(t, "main.Wkid", reflect.TypeOf(c).String())
		equals(t, 4326, c.Wkid)
		equals(t, 4326, c.LatestWkid)
	}
}

func TestJSONSpatialReference(t *testing.T) {

	jsonStream := `{"spatialReference":{"wkid":4326,"latestWkid":4326}}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var c SpatialReference
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		equals(t, "main.SpatialReference", reflect.TypeOf(c).String())
		equals(t, 4326, c.SpatialReference.Wkid)
		equals(t, 4326, c.SpatialReference.LatestWkid)
	}
}

func TestJSONFullSinglePayload(t *testing.T) {

	jsonStream := `{"spatialReference":{"wkid":4326,"latestWkid":4326},"candidates":[{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612517766347352,"y":38.352004138620323},"score":93.200000000000003,"attributes":{}},{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612229940922532,"y":38.353295843541105},"score":93.200000000000003,"attributes":{}},{"address":"507 S PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.61248740139186,"y":38.352168562340495},"score":100,"attributes":{}},{"address":"507 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612448720360234,"y":38.351815864580452},"score":100,"attributes":{}},{"address":"507 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612391256749248,"y":38.35311332552839},"score":90.870000000000005,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612757556175097,"y":38.351796024941883},"score":79,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612397010945358,"y":38.351714488143401},"score":79,"attributes":{}},{"address":"508 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.61254859196886,"y":38.353055632076831},"score":69.870000000000005,"attributes":{}},{"address":"505 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612366186770942,"y":38.351796905938166},"score":68.5,"attributes":{}},{"address":"521 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.614312182687073,"y":38.352233582879819},"score":68.489999999999995,"attributes":{}}]}`

	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	var c Candidates
	for {
		if err := decoder.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		equals(t, "main.Candidates", reflect.TypeOf(c).String())
		equals(t, 10, len(c.Candidates))
	}
}

func TestJSONMultipleResponses(t *testing.T) {

	jsonStream := `[{"spatialReference":{"wkid":4326,"latestWkid":4326},"candidates":[{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612517766347352,"y":38.352004138620323},"score":93.200000000000003,"attributes":{}},{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612229940922532,"y":38.353295843541105},"score":93.200000000000003,"attributes":{}},{"address":"507 S PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.61248740139186,"y":38.352168562340495},"score":100,"attributes":{}},{"address":"507 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612448720360234,"y":38.351815864580452},"score":100,"attributes":{}},{"address":"507 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612391256749248,"y":38.35311332552839},"score":90.870000000000005,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612757556175097,"y":38.351796024941883},"score":79,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612397010945358,"y":38.351714488143401},"score":79,"attributes":{}},{"address":"508 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.61254859196886,"y":38.353055632076831},"score":69.870000000000005,"attributes":{}},{"address":"505 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612366186770942,"y":38.351796905938166},"score":68.5,"attributes":{}},{"address":"521 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.614312182687073,"y":38.352233582879819},"score":68.489999999999995,"attributes":{}}]},{"spatialReference":{"wkid":4326,"latestWkid":4326},"candidates":[{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612517766347352,"y":38.352004138620323},"score":93.200000000000003,"attributes":{}},{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612229940922532,"y":38.353295843541105},"score":93.200000000000003,"attributes":{}},{"address":"507 S PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.61248740139186,"y":38.352168562340495},"score":100,"attributes":{}},{"address":"507 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612448720360234,"y":38.351815864580452},"score":100,"attributes":{}},{"address":"507 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612391256749248,"y":38.35311332552839},"score":90.870000000000005,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612757556175097,"y":38.351796024941883},"score":79,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612397010945358,"y":38.351714488143401},"score":79,"attributes":{}},{"address":"508 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.61254859196886,"y":38.353055632076831},"score":69.870000000000005,"attributes":{}},{"address":"505 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612366186770942,"y":38.351796905938166},"score":68.5,"attributes":{}},{"address":"521 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.614312182687073,"y":38.352233582879819},"score":68.489999999999995,"attributes":{}}]}]`

	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	var c []*Candidates
	for {
		if err := decoder.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		equals(t, "[]*main.Candidates", reflect.TypeOf(c).String())
		equals(t, 2, len(c))

		for i := 0; i < len(c); i++ {
			equals(t, 10, len(c[i].Candidates))
		}
	}
}

func TestCandidatesBestMatch(t *testing.T) {

	jsonStream := `{"spatialReference":{"wkid":4326,"latestWkid":4326},"candidates":[{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612517766347352,"y":38.352004138620323},"score":93.200000000000003,"attributes":{}},{"address":"507 PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.612229940922532,"y":38.353295843541105},"score":93.200000000000003,"attributes":{}},{"address":"507 S PINEHURST AVE, SALISBURY, MD, 21801","location":{"x":-75.61248740139186,"y":38.352168562340495},"score":100,"attributes":{}},{"address":"507 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612448720360234,"y":38.351815864580452},"score":100,"attributes":{}},{"address":"507 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612391256749248,"y":38.35311332552839},"score":90.870000000000005,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612757556175097,"y":38.351796024941883},"score":79,"attributes":{}},{"address":"508 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612397010945358,"y":38.351714488143401},"score":79,"attributes":{}},{"address":"508 N Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.61254859196886,"y":38.353055632076831},"score":69.870000000000005,"attributes":{}},{"address":"505 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.612366186770942,"y":38.351796905938166},"score":68.5,"attributes":{}},{"address":"521 S Pinehurst Ave, SALISBURY, MD, 21801","location":{"x":-75.614312182687073,"y":38.352233582879819},"score":68.489999999999995,"attributes":{}}]}`

	unmarshaler := JSONMarshaler{}
	candidates, err := unmarshaler.UnmarshalAddresses(strings.NewReader(jsonStream))
	ok(t, err)

	fmt.Println(candidates.GetScores())
	candidates.SortCandidates()

	sortedScores := candidates.GetScores()

	equals(t, float32(100), sortedScores[0])
}
