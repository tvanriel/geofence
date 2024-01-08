package cmd

import "github.com/tvanriel/geofence/app"

func NewDiscordReporter(config *Configuration, geo *app.Geo) (*app.DiscordReporter, error) {
	return app.NewDiscordReporter(config.Report.BotToken, config.Report.ChannelID, geo)
}
