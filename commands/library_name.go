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
	libByNameURL = "https://api.cdnjs.com/libraries/"
)

// LibByNameCommand is a command which returns specefic library
func LibByNameCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
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

		resp, err := http.Get(libByNameURL + args[2])
		h.HandleError(err)
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			s.ChannelMessageSend(m.ChannelID, "❌ Library not found")
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

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

		libHeader := "ℹ️ *CDNJS LIB NAME RESULTS*:\n\n"
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
