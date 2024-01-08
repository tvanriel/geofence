package app

import (
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
)

type Geo struct {
	db *geoip2.Reader
}

func NewGeo(filename string) (*Geo, error) {
	db, err := geoip2.Open(filename)

	if err != nil {
		return nil, err
	}

	return &Geo{
		db: db,
	}, nil
}

func (g *Geo) City(address string) string {
	ip := net.ParseIP(address)
	record, err := g.db.City(ip)
	if err != nil {
		return ""
	}
	return record.City.Names["en-EN"]
}

func (g *Geo) Country(address string) string {
	ip := net.ParseIP(address)
	record, err := g.db.City(ip)
	if err != nil {
		return ""
	}
	return record.Country.Names["en-EN"]
}
