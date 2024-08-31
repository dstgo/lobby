package internal

import "embed"

//go:embed *
var FS embed.FS

const (
	GeoIp2CityDB = "geoip2/GeoLite2-City.mmdb"
)
