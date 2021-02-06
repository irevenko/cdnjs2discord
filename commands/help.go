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

	if strings.Trim(m.Content, " ") == "!cdn" || strings.Trim(m.Content, " ") == "!cdn help" {
		helpHeader := "🆘 *HELP FOR THE CDNJS BOT*:\n"
		helpHelp := "**!cdn help** or **!cdn** - displays help info\n"
		statsHelp := "**!cdn stats** - returns cdnjs libraries number\n"
		libHelp := "**!cdn lib** <*LIBRARY NAME*> - returns specific lib\n"
		assetsHelp := "**!cdn assets** <*NAME*> <*VERSION*> - returns assets for the specific lib version\n"
		whitelistHelp := "**!cdn whitelist** - returns cdnjs extension whitelist\n"
		sourceHelp := "__Bot Source Code__: https://github.com/irevenko/cdnjs2discord"

		helpMsg := helpHeader + helpHelp + statsHelp + whitelistHelp + libHelp + assetsHelp + sourceHelp

		s.ChannelMessageSend(m.ChannelID, helpMsg)
	}
}
