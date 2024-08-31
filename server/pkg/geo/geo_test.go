package geo

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestCity(t *testing.T) {
	city, err := City(net.ParseIP("43.142.136.80"))
	if !assert.NoError(t, err) {
		return
	}
	t.Logf("%+v\n", city.City)
	t.Logf("%+v\n", city.Country)
	t.Logf("%+v\n", city.Postal)
	t.Logf("%+v\n", city.Continent)
	t.Logf("%+v\n", city.Location)
	t.Logf("%+v\n", city.Subdivisions)
	t.Logf("%+v\n", city.RepresentedCountry)
	t.Logf("%+v\n", city.RegisteredCountry)
}
