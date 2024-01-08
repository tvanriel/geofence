package cmd

import "github.com/tvanriel/geofence/app"

func Serivces(config *Configuration, geo *app.Geo, reporter *app.DiscordReporter) ([]app.Proxy, error) {
	proxies := make([]app.Proxy, len(config.Services))
	for s := range config.Services {
		middlewares := make([]app.Middleware, len(config.Services[s].Rules))
		for m := range config.Services[s].Rules {
			middleware, err := app.NewMiddleware(config.Services[s].Rules[m].If, geo)
			if err != nil {
				return nil, err
			}
			middlewares[m] = *middleware
		}
		proxy := app.NewProxy(
			middlewares,
			config.Services[s].Listen,
			config.Services[s].Upstream,
			config.Services[s].Name,
			reporter,
		)
		proxies[s] = *proxy
	}
	return proxies, nil
}
