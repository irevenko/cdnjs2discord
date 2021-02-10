package commands

import (
	h "../helpers"
	t "../types"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	baseSearchURL = "https://api.cdnjs.com/libraries/?search="
	searchNameParams = "&fields=description&limit=20"
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

		resp, err := http.Get(baseSearchURL + args[3] + searchNameParams)
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

		var libraries []string

		var libNames []string
		var libDescriptions []string
		var cdnjsLinks []string

		for i, v := range searchResults {
			libNames = append(libNames, strconv.Itoa(i + 1) + ") " + v.Name)
			libDescriptions = append(libDescriptions, v.Description)
			cdnjsLinks = append(cdnjsLinks, baseLibNameURL + v.Name +"\n")
		}

		for i, _ := range searchResults{
			libraries = append(libraries, libNames[i])
			libraries = append(libraries, libDescriptions[i])
			libraries = append(libraries, cdnjsLinks[i])
		}

		// 3 elems split
		// clamp function

		//firstPage := libraries[0:15]
		//secondPage := libraries[15:30]
		//thirdPage := libraries[30:45]
		//fourthPage := libraries[45:60]

		//fmt.Println(libraries[0:3])
		//fmt.Println(firstPage)
		//fmt.Println(secondPage)
		//fmt.Println(thirdPage)
		//fmt.Println(fourthPage)

		fmt.Println(libraries)

		searchNameHeader := "🔎 *CDNJS NAME SEARCH RESULTS*:\n"
		searchNameMsg := strings.Join(libraries[0:15], "\n")

		msg, err := s.ChannelMessageSend(m.ChannelID, searchNameHeader+searchNameMsg)
		h.HandleError(err)

		s.MessageReactionAdd(msg.ChannelID, msg.ID, "1️⃣")
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "2️⃣")
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "3️⃣")
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "4️⃣")
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "5️⃣")

		s.AddHandler(func (s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			go func() {
				switch r.Emoji.Name {
				case "1️⃣":
					s.ChannelMessageEdit(msg.ChannelID, msg.ID, "1")
				case "2️⃣":
					s.ChannelMessageEdit(msg.ChannelID, msg.ID, "2")
				case "3️⃣":
					s.ChannelMessageEdit(msg.ChannelID, msg.ID, "3")
				case "4️⃣":
					s.ChannelMessageEdit(msg.ChannelID, msg.ID, "4")
				case "5️⃣":
					s.ChannelMessageEdit(msg.ChannelID, msg.ID, "5")
				default:
					s.ChannelMessageSend(m.ChannelID, "⛔ Reaction Error (use only pages reactions 1-5)")
				}
			}()
		})
	}
}