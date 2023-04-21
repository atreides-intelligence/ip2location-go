package ip2location_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/atreides-intelligence/ip2location-go"
)

func TestIP2L(m *testing.T) {
	ip := "8.8.8.8"

	location, _ := ioutil.ReadFile("./IP2LOCATION-LITE-DB11.BIN")

	locationdb, err := ip2location.OpenDBWithBytes(location)
	if err != nil {
		panic(err)
	}

	locationInfo, err := locationdb.Get_all(ip)
	if err != nil {
		panic(err)
	}

	// Print results
	fmt.Println("IP:", ip)
	x1, _ := json.Marshal(locationInfo)
	fmt.Println(string(x1))
}
