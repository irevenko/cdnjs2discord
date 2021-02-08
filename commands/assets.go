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
	baseLibVersionURL = "https://api.cdnjs.com/libraries/"
	baseCloudFlareURL = "https://cdnjs.cloudflare.com/ajax/libs/"
)

// AssetsCommand is a command which returns assets for the specific library
func AssetsCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn assets") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		if len(args) < 4 {
			s.ChannelMessageSend(m.ChannelID, "⚠️ Command **assets** requires 2 arguments (*lib name*, *lib version*)")
			return
		}

		resp, err := http.Get(baseLibVersionURL + args[2] + "/" + args[3])
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		if strings.Contains(string(body), "Library not found") {
			s.ChannelMessageSend(m.ChannelID, "❌ Library **"+args[2]+"** not found!")
			return
		}

		if strings.Contains(string(body), "Version not found") {
			s.ChannelMessageSend(m.ChannelID, "❌ Version *"+args[3]+"* for the **"+args[2]+"** library"+" not found!")
			return
		}

		var specificLibVerResp t.SpecificLibVerResponse
		json.Unmarshal(body, &specificLibVerResp)

		rawFiles := specificLibVerResp.Files

		for idx := range rawFiles {
			orderNum := strconv.Itoa(idx + 1)
			rawFiles[idx] = orderNum + ") " + baseCloudFlareURL + specificLibVerResp.Name + "/" + specificLibVerResp.Version + "/" + rawFiles[idx]
		}

		files := strings.Join(rawFiles[:], "\n")

		assetsHeader := "🔎 *CDNJS ASSETS RESULTS*:\n"
		assetsMsg :=
			"➡️ **Library**: " + specificLibVerResp.Name +
				" ver " + specificLibVerResp.Version +
				"\n🔗 **Assets Links**:\n" + files

		s.ChannelMessageSend(m.ChannelID, assetsHeader+assetsMsg)
	}
}
