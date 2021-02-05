package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	h "../helpers"
	t "../types"
	"github.com/bwmarrin/discordgo"
)

const (
	whitelistURL = "https://api.cdnjs.com/whitelist"
)

// WhiteListCommand is a command which returns cdnjs extension whitelist
func WhiteListCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Trim(m.Content, " ") == "!cdn whitelist" {
		resp, err := http.Get(whitelistURL)
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		var whitelistResp t.WhiteListResponse
		json.Unmarshal(body, &whitelistResp)
		fmt.Println()

		rawExtensions := whitelistResp.Categories

		extensions := make([]string, 0, len(rawExtensions))
		for key, val := range rawExtensions {
			extensions = append(extensions, "- "+val+" : "+key)
		}

		whitelistHeader := "📃 *CDNJS Extension Whitelist*:\n"
		whitelistMsg := strings.Join(extensions[:], "\n")

		s.ChannelMessageSend(m.ChannelID, whitelistHeader+whitelistMsg)
	}
}
