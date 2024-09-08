package staticdata

import (
	_ "embed"
	"encoding/json"
	"errors"
	"sync"

	"github.com/adrianolmedo/components/strutil"
)

//go:embed provinces.json
var provincesAndCitiesJSONBytes []byte

// ProvinceWithCities is a struct that represents a province
// with its name for each provider and its cities
type ProvinceWithCities struct {
	ProvinceName     string   `json:"provinceName"`
	ProvinceNameAire string   `json:"provinceNameAire"`
	ProvinceNameFi   string   `json:"provinceNameFi"`
	Cities           []string `json:"cities"`
}

var parsedProvincesWithCities []ProvinceWithCities
var parsedProvincesWithCitiesSync = sync.Once{}

// ProvincesWithCities returns a slice of items that represents
// all the provinces with their names for each provider
// and their cities.
func ProvincesWithCities() []ProvinceWithCities {
	parsedProvincesWithCitiesSync.Do(func() {
		_ = json.Unmarshal(provincesAndCitiesJSONBytes, &parsedProvincesWithCities)
	})

	return parsedProvincesWithCities
}

// GetCitiesByProvinceName returns a slice of strings that represents
// all the cities of a province or an empty slice if the province
// doesn't exist.
func GetCitiesByProvinceName(provinceName string) []string {
	for _, province := range ProvincesWithCities() {
		if strutil.EqualNoStrict(province.ProvinceName, provinceName) {
			return province.Cities
		}
	}

	return []string{}
}

// GetAireProvinceName returns the name of a province for Aire
// or an error if the province doesn't exist.
func GetAireProvinceName(provinceName string) (string, error) {
	for _, province := range ProvincesWithCities() {
		if strutil.EqualNoStrict(province.ProvinceName, provinceName) {
			return province.ProvinceNameAire, nil
		}
	}

	return "", errors.New("province not found")
}

// GetFiProvinceName returns the name of a province for Fi
// or an error if the province doesn't exist.
func GetFiProvinceName(provinceName string) (string, error) {
	for _, province := range ProvincesWithCities() {
		if strutil.EqualNoStrict(province.ProvinceName, provinceName) {
			return province.ProvinceNameFi, nil
		}
	}

	return "", errors.New("province not found")
}
