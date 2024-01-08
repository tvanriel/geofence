package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

type DiscordReporter struct {
	client  *disgord.Client
	Channel string
	Geo     *Geo
}

func NewDiscordReporter(botToken string, channel string, geo *Geo) (*DiscordReporter, error) {
	if botToken == "" {
		return nil, errors.New("no bot token was provided where one was expected")
	}
	discordClient := disgord.New(disgord.Config{
		Intents:  disgord.AllIntents(),
		BotToken: botToken,
	})

	bot := &DiscordReporter{
		client:  discordClient,
		Channel: channel,
	}

	bot.client.Gateway().BotReady(func() {
		fmt.Println("Bot ready")
		bot.client.Logger().Info(bot.client.BotAuthorizeURL(disgord.PermissionAddReactions, []string{}))
		user, err := bot.client.CurrentUser().Get()
		if err != nil {
			bot.client.Logger().Error("Unable to fetch own user?")
			return
		}
		bot.client.Logger().Info("Logged in as %s#%s", user.Username, user.Discriminator)
	})
	go func() {
		fmt.Println("Discord bot has booted.")
		bot.client.Gateway().StayConnectedUntilInterrupted()
	}()

	return bot, nil
}

func (r *DiscordReporter) ReportBlocked(service string, ipaddress string, rule string) {
	city := r.Geo.City(ipaddress)
	country := r.Geo.Country(ipaddress)
	flag := strings.Join([]string{
		":flag_",
		country,
		":",
	}, "")

	r.client.Channel(disgord.ParseSnowflakeString(r.Channel)).CreateMessage(&disgord.CreateMessage{
		Content: fmt.Sprintf(
			"[%s] Blocked connection from `%s`: `%s`\n%s, %s %s",
			service,
			ipaddress,
			rule,
			city,
			country,
			flag,
		),
	})
}

func (r *DiscordReporter) ReportAccepted(service string, ipaddress string) {
	city := r.Geo.City(ipaddress)
	country := r.Geo.Country(ipaddress)
	flag := strings.Join([]string{
		":flag_",
		country,
		":",
	}, "")

	r.client.Channel(disgord.ParseSnowflakeString(r.Channel)).CreateMessage(&disgord.CreateMessage{
		Content: fmt.Sprintf(
			"[%s] Opened connection from `%s`.\n%s, %s %s",
			service,
			ipaddress,
			city,
			country,
			flag,
		),
	})
}

func (r *DiscordReporter) CannotDial(service string, ipaddress string, err error) {
	city := r.Geo.City(ipaddress)
	country := r.Geo.Country(ipaddress)
	flag := strings.Join([]string{
		":flag_",
		country,
		":",
	}, "")

	r.client.Channel(disgord.ParseSnowflakeString(r.Channel)).CreateMessage(&disgord.CreateMessage{
		Content: fmt.Sprintf(
			"[%s] Cannot dial upstream for `%s`: `%v`\n%s, %s %s",
			service,
			ipaddress,
			err,
			city,
			country,
			flag,
		),
	})
}
