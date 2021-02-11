package commands

import (
	h "../helpers"
	t "../types"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	baseSearchURL = "https://api.cdnjs.com/libraries/?search="
	searchNameParams = "&fields=description&limit=25"
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
			libDescriptions = append(libDescriptions, " "+v.Description)
			cdnjsLinks = append(cdnjsLinks, "\n"+ baseLibNameURL + v.Name +"\n")
		}

		for i, _ := range searchResults{
			libraries = append(libraries, libNames[i])
			libraries = append(libraries, libDescriptions[i])
			libraries = append(libraries, cdnjsLinks[i])
		}

		pages := h.SplitIntoPages(libraries, 15)

		var firstPage []string
		var secondPage []string
		var thirdPage []string
		var fourthPage []string
		var fifthPage []string

		for i, v := range pages {
			switch i {
			case 0:
				firstPage = append(firstPage, strings.Join(v, ""))
			case 1:
				secondPage = append(secondPage, strings.Join(v, ""))
			case 2:
				thirdPage = append(thirdPage, strings.Join(v, ""))
			case 3:
				fourthPage = append(fourthPage, strings.Join(v, ""))
			case 4:
				fifthPage = append(fifthPage, strings.Join(v, ""))
			default:
				break
			}
		}

		searchNameHeader := "🔎 *CDNJS NAME SEARCH RESULTS*:\n"
		searchNameMsg := strings.Join(firstPage[:], "\n")

		nameMsg, err := s.ChannelMessageSend(m.ChannelID, searchNameHeader+searchNameMsg)
		h.HandleError(err)

		s.MessageReactionAdd(nameMsg.ChannelID, nameMsg.ID, "1️⃣")
		s.MessageReactionAdd(nameMsg.ChannelID, nameMsg.ID, "2️⃣")
		s.MessageReactionAdd(nameMsg.ChannelID, nameMsg.ID, "3️⃣")
		s.MessageReactionAdd(nameMsg.ChannelID, nameMsg.ID, "4️⃣")
		s.MessageReactionAdd(nameMsg.ChannelID, nameMsg.ID, "5️⃣")
		s.MessageReactionAdd(nameMsg.ChannelID, nameMsg.ID, "😑")

		s.AddHandler(func (s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			go func() {
				switch r.Emoji.Name {
				case "1️⃣":
					if len(firstPage) != 0 {
						if nameMsg.ID == r.MessageID {
							s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, searchNameHeader + strings.Join(firstPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "2️⃣":
					if len(secondPage) != 0 {
						if nameMsg.ID == r.MessageID {
							s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, searchNameHeader+strings.Join(secondPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "3️⃣":
					if len(thirdPage) != 0 {
						if nameMsg.ID == r.MessageID {
							s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, searchNameHeader+strings.Join(thirdPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "4️⃣":
					if len(fourthPage) != 0 {
						if nameMsg.ID == r.MessageID {
							s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, searchNameHeader+strings.Join(fourthPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "5️⃣":
					if len(fifthPage) != 0 {
						if nameMsg.ID == r.MessageID {
							s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, searchNameHeader+strings.Join(fifthPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(nameMsg.ChannelID, nameMsg.ID, "⛔ Page Error (limit reached)")
					}
				default:
					break
				}
			}()
		})
	}
}