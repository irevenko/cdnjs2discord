package commands

import (
	h "../helpers"
	t "../types"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseSearchURL = "https://api.cdnjs.com/libraries/?search="
	baseSearchParams = "&fields=description&limit=5"
)

// SearchNameCommand is a command which returns search results based on lib name
func SearchNameCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn search name") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		if len(args) < 4 {
			s.ChannelMessageSend(m.ChannelID, "⚠️ Command **search name** requires 1 argument (library name)")
			return
		}

		resp, err := http.Get(baseSearchURL + args[3] + baseSearchParams)
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		var nameSearchResp t.SearchResponse
		json.Unmarshal(body, &nameSearchResp)
		searchResults := nameSearchResp.Results

		if len(searchResults) == 0 {
			s.ChannelMessageSend(m.ChannelID, "❌ Library **"+args[3]+"** not found!")
			return
		}


		var libNames []string
		var libDescriptions []string
		var libAssets []string
		var cdnjsLinks []string

		for _, v := range searchResults {
			libNames = append(libNames, v.Name)
			libDescriptions = append(libDescriptions, v.Description)
			libAssets = append(libAssets, v.LatestLink)
			cdnjsLinks = append(cdnjsLinks, baseLibNameURL + v.Name)
		}

		//libName := strings.Join(libNames, "\n")
		//libDesc := strings.Join(libDescriptions, "\n")
		//libAssetLink := "<" + nameSearchResp.LatestLink + ">"
		//cdnjsLink := "<" + "https://cdnjs.com/libraries/" + libName + ">"

		//searchNameHeader := "🔎 *CDNJS NAME SEARCH RESULTS*:\n"
		//searchNameMsg :=
		//	"➡️ **Name**: " + libName + "\n" +
		//		"📜 **Description**: " + libDesc + "\n" +
		//		"🔗 **Asset Link**: " + libAssetLink + "\n" +
		//		"⚓ **CDNJS Reference**: " + cdnjsLink + "\n"

		s.ChannelMessageSend(m.ChannelID, strings.Join(libNames, "\n"))//searchNameHeader+searchNameMsg)
	}
}