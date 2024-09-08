package staticdata

import (
	_ "embed"
	"encoding/json"
	"sync"
)

//go:embed countries.json
var countriesJSONBytes []byte

// Country is a struct that represents a country
// with its name in spanish and english, its iso2 and iso3 codes
type Country struct {
	NameES    string `json:"nameES"`
	NameEN    string `json:"nameEN"`
	Iso2      string `json:"iso2"`
	Iso3      string `json:"iso3"`
	PhoneCode string `json:"phoneCode"`
}

var parsedCountries []Country
var parsedCountriesSync = sync.Once{}

// Countries returns a json string with all the countries
// with their names in spanish and english, and their iso2 and iso3 codes
func Countries() []Country {
	parsedCountriesSync.Do(func() {
		_ = json.Unmarshal(countriesJSONBytes, &parsedCountries)
	})

	return parsedCountries
}
