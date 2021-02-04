package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	h "../helpers"
	t "../types"
	"github.com/bwmarrin/discordgo"
)

const (
	statsURL = "https://api.cdnjs.com/stats"
)

// StatsCommand is a command which returns cdnjs stats
func StatsCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn stats") {
		resp, err := http.Get(statsURL)
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		var statsResp t.StatsResponse
		json.Unmarshal(body, &statsResp)

		librariesNumber := strconv.Itoa(statsResp.LibrariesNumber)

		s.ChannelMessageSend(m.ChannelID, "📊 *CDNJS STATS*\nThe total number of libraries available on CDNJS : "+"```yaml\n"+librariesNumber+"```")
	}
}
