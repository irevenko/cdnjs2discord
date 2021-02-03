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

// LibByNameCommand is a command which returns specefic library by name
func LibByNameCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn lib") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		resp, err := http.Get(libByNameURL + args[2])
		h.HandleError(err)

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		var specificLibResp t.SpecificLibResponse
		json.Unmarshal(body, &specificLibResp)

		libName := specificLibResp.Name
		libLicense := specificLibResp.License
		libDesc := specificLibResp.Description
		libSource := specificLibResp.Repository.URL
		libHomePage := specificLibResp.HomePage
		libVersion := specificLibResp.Version
		libAuthor := specificLibResp.Author
		libLatestLink := specificLibResp.LatestLink

		s.ChannelMessageSend(m.ChannelID, libName+" "+libVersion+"\n"+libDesc+"\n"+libAuthor+"\n"+libLicense+"\n"+libLatestLink+"\n"+libSource+"\n"+libHomePage)
	}
}
