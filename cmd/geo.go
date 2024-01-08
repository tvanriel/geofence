package cmd

import "github.com/tvanriel/geofence/app"

func NewGeo(config *Configuration) (*app.Geo, error) {
	return app.NewGeo(config.Db)
}
