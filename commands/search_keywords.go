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
	searchKeyWordsParams = "&fields=description&limit=25&search_fields=keywords"
)

// SearchKeyWordsCommand is a command which returns search results based on keywords
func SearchKeyWordsCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn search keywords") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		if len(args) < 4 {
			s.ChannelMessageSend(m.ChannelID, "⚠️ Command **search keywords** requires 1 argument (keywords 1 word or separated by comma)")
			return
		}

		resp, err := http.Get(baseSearchURL + args[3] + searchKeyWordsParams)
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		var keywordsSearchResp t.SearchResponse
		json.Unmarshal(body, &keywordsSearchResp)
		searchResults := keywordsSearchResp.Results

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

		searchKeyWordsHeader := "🔎 *CDNJS NAME SEARCH RESULTS*:\n"
		searchKeyWordsMsg := strings.Join(firstPage[:], "\n")

		keywordsMsg, err := s.ChannelMessageSend(m.ChannelID, searchKeyWordsHeader+searchKeyWordsMsg)
		h.HandleError(err)

		s.MessageReactionAdd(keywordsMsg.ChannelID, keywordsMsg.ID, "1️⃣")
		s.MessageReactionAdd(keywordsMsg.ChannelID, keywordsMsg.ID, "2️⃣")
		s.MessageReactionAdd(keywordsMsg.ChannelID, keywordsMsg.ID, "3️⃣")
		s.MessageReactionAdd(keywordsMsg.ChannelID, keywordsMsg.ID, "4️⃣")
		s.MessageReactionAdd(keywordsMsg.ChannelID, keywordsMsg.ID, "5️⃣")
		s.MessageReactionAdd(keywordsMsg.ChannelID, keywordsMsg.ID, "😑")

		s.AddHandler(func (s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			go func() {
				switch r.Emoji.Name {
				case "1️⃣":
					if len(firstPage) != 0 {
						if keywordsMsg.ID == r.MessageID {
							s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, searchKeyWordsHeader+strings.Join(firstPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "2️⃣":
					if len(secondPage) != 0 {
						if keywordsMsg.ID == r.MessageID {
							s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, searchKeyWordsHeader+strings.Join(secondPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "3️⃣":
					if len(thirdPage) != 0 {
						if keywordsMsg.ID == r.MessageID {
							s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, searchKeyWordsHeader+strings.Join(thirdPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "4️⃣":
					if len(fourthPage) != 0 {
						if keywordsMsg.ID == r.MessageID {
							s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, searchKeyWordsHeader+strings.Join(fourthPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "5️⃣":
					if len(fifthPage) != 0 {
						if keywordsMsg.ID == r.MessageID {
							s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, searchKeyWordsHeader+strings.Join(fifthPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(keywordsMsg.ChannelID, keywordsMsg.ID, "⛔ Page Error (limit reached)")
					}
				default:
					break
				}
			}()
		})
	}
}