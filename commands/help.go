package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// HelpCommand - we all know what it does
func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn help") {
		helpHeader := "🧾 *HELP FOR THE CDNJS BOT*\n"
		statsHelp := "**!cdn stats** - returns cdnjs libraries number\n"
		libHelp := "**!cdn lib** <*LIBRARY NAME*> - search for a specific lib\n"
		sourceHelp := "__Source Code__: https://github.com/irevenko/cdnjs2discord"

		helpMsg := helpHeader + statsHelp + libHelp + sourceHelp

		s.ChannelMessageSend(m.ChannelID, helpMsg)
	}
}
