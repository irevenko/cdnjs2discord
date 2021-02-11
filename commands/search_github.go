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
	searchGitHubParams = "&fields=description&limit=25&search_fields=github.user"
)

// SearchGitHubCommand is a command which returns search results based on github username
func SearchGitHubCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!cdn search github") {
		command := strings.Trim(m.Content, " ")
		args := strings.Split(command, " ")

		if len(args) < 4 {
			s.ChannelMessageSend(m.ChannelID, "⚠️ Command **search github** requires 1 argument (github username)")
			return
		}

		resp, err := http.Get(baseSearchURL + args[3] + searchGitHubParams)
		h.HandleError(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		h.HandleError(err)

		var githubSearchResp t.SearchResponse
		json.Unmarshal(body, &githubSearchResp)
		searchResults := githubSearchResp.Results

		if len(searchResults) == 0 {
			s.ChannelMessageSend(m.ChannelID, "❌ User **"+args[3]+"** not found!")
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

		for i := range searchResults{
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

		searchGithubHeader := "🔎 *CDNJS GITHUB SEARCH RESULTS*:\n"
		searchGithubMsg := strings.Join(firstPage[:], "\n")

		githubMsg, err := s.ChannelMessageSend(m.ChannelID, searchGithubHeader+searchGithubMsg)
		h.HandleError(err)

		s.MessageReactionAdd(githubMsg.ChannelID, githubMsg.ID, "1️⃣")
		s.MessageReactionAdd(githubMsg.ChannelID, githubMsg.ID, "2️⃣")
		s.MessageReactionAdd(githubMsg.ChannelID, githubMsg.ID, "3️⃣")
		s.MessageReactionAdd(githubMsg.ChannelID, githubMsg.ID, "4️⃣")
		s.MessageReactionAdd(githubMsg.ChannelID, githubMsg.ID, "5️⃣")
		s.MessageReactionAdd(githubMsg.ChannelID, githubMsg.ID, "😑")

		s.AddHandler(func (s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			go func() {
				switch r.Emoji.Name {
				case "1️⃣":
					if len(firstPage) != 0 {
						if githubMsg.ID == r.MessageID {
							s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, searchGithubHeader+strings.Join(firstPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "2️⃣":
					if len(secondPage) != 0 {
						if githubMsg.ID == r.MessageID {
							s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, searchGithubHeader+strings.Join(secondPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "3️⃣":
					if len(thirdPage) != 0 {
						if githubMsg.ID == r.MessageID {
							s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, searchGithubHeader+strings.Join(thirdPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "4️⃣":
					if len(fourthPage) != 0 {
						if githubMsg.ID == r.MessageID {
							s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, searchGithubHeader+strings.Join(fourthPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, "⛔ Page Error (limit reached)")
					}
				case "5️⃣":
					if len(fifthPage) != 0 {
						if githubMsg.ID == r.MessageID {
							s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, searchGithubHeader+strings.Join(fifthPage[:], "\n"))
						}
					} else {
						s.ChannelMessageEdit(githubMsg.ChannelID, githubMsg.ID, "⛔ Page Error (limit reached)")
					}
				default:
					break
				}
			}()
		})
	}
}