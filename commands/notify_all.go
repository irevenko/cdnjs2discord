package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func NotifyEveryoneCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!notif_all") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		duration, _ := strconv.Atoi(args[1])
		go func() {
			time.Sleep(time.Second * time.Duration(duration))
			notificationMsg := "@everyone 📡\n" + ">>> " + strings.Join(args[2:], " ")
			s.ChannelMessageSend(m.ChannelID, notificationMsg)
		}()
	}

}
