package geo

import (
	"github.com/dstgo/lobby/pkg/geo/internal"
	"github.com/oschwald/geoip2-golang"
	"net"
)

var geo *geoip2.Reader

func init() {
	db, err := OpenGeoDB()
	if err != nil {
		panic(err)
	}
	geo = db
}

// OpenGeoDB returns a geoip2.Reader that loads in memory
func OpenGeoDB() (*geoip2.Reader, error) {
	bytes, err := internal.FS.ReadFile(internal.GeoIp2CityDB)
	if err != nil {
		return nil, err
	}
	geo, err := geoip2.FromBytes(bytes)
	if err != nil {
		return nil, err
	}
	return geo, nil
}

func Enterprise(ipAddress net.IP) (*geoip2.Enterprise, error) {
	return geo.Enterprise(ipAddress)
}

func City(ipAddress net.IP) (*geoip2.City, error) {
	return geo.City(ipAddress)
}

func Country(ipAddress net.IP) (*geoip2.Country, error) {
	return geo.Country(ipAddress)
}

func AnonymousIP(ipAddress net.IP) (*geoip2.AnonymousIP, error) {
	return geo.AnonymousIP(ipAddress)
}

func ISP(ipAddress net.IP) (*geoip2.ISP, error) {
	return geo.ISP(ipAddress)
}

func Domain(ipAddress net.IP) (*geoip2.Domain, error) {
	return geo.Domain(ipAddress)
}

func ConnectionType(ipAddress net.IP) (*geoip2.ConnectionType, error) {
	return geo.ConnectionType(ipAddress)
}

func ASN(ipAddress net.IP) (*geoip2.ASN, error) {
	return geo.ASN(ipAddress)
}
