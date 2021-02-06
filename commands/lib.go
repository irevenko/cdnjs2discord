package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	h "../helpers"
	t "../types"
	"github.com/bwmarrin/discordgo"
)

const (
	baseLibByNameURL = "https://api.cdnjs.com/libraries/"
)

// LibCommand is a command which returns specefic library info
func LibCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn lib") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		if len(args) < 3 {
			s.ChannelMessageSend(m.ChannelID, "⚠️ Command **lib** requires 1 argument (library name)")
			return
		}

		resp, err := http.Get(baseLibByNameURL + args[2])
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		if strings.Contains(string(body), "Library not found") {
			s.ChannelMessageSend(m.ChannelID, "❌ Library **"+args[2]+"** not found!")
			return
		}

		var specificLibResp t.SpecificLibResponse
		json.Unmarshal(body, &specificLibResp)

		libName := specificLibResp.Name
		libLicense := specificLibResp.License
		libDesc := specificLibResp.Description
		libSource := "<" + specificLibResp.Repository.URL + ">"
		libHomePage := "<" + specificLibResp.HomePage + ">"
		libVersion := specificLibResp.Version
		libAuthor := specificLibResp.Author
		libAssetLink := "<" + specificLibResp.LatestLink + ">"
		keyWords := strings.Join(specificLibResp.KeyWords[:], ", ")
		autoUpdate := specificLibResp.AutoUpdate.Source + " | " + specificLibResp.AutoUpdate.Target
		cdnjsLink := "<" + "https://cdnjs.com/libraries/" + libName + ">"

		libHeader := "🔎 *CDNJS LIB RESULTS*:\n"
		libMessage :=
			"➡️ **Name**: " + libName + "\n" +
				"🔖 **Version**: " + libVersion + "\n" +
				"📜 **Description**: " + libDesc + "\n" +
				"📒 **Key words**: " + keyWords + "\n" +
				"✒️ **Author**: " + libAuthor + "\n" +
				"📑 **License**: " + libLicense + "\n" +
				"♻️ **Auto Update**: " + autoUpdate + "\n" +
				"🔗 **Asset Link**: " + libAssetLink + "\n" +
				"⚓ **CDNJS Reference**: " + cdnjsLink + "\n" +
				"🗂 **Source Code**: " + libSource + "\n" +
				"📍 **Home Page**: " + libHomePage

		s.ChannelMessageSend(m.ChannelID, libHeader+libMessage)
	}
}
