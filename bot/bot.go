package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

var tab struct {
	tableau [3][3]string
	joueur  bool
}
var id struct {
	baben string
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	goBot.AddHandler(messageHandler)
	BotID = u.ID

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	fmt.Println("!ascii / !ascii2 / !ascii3")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {

			return
		}

		split := strings.Split(m.Content, " ")

		for _, str := range split {
			fmt.Println("arg: ", str)
		}

		fmt.Println("commande: ", split[1:])
		if len(split) >= 1 {
			if split[0] == "!spam" {
				if len(split) > 1 {
					for i := 0; i <= 5; i++ {
						s.ChannelMessageSend(m.ChannelID, split[1])

					}
				} else {
					s.ChannelMessageSend(m.ChannelID, "!spam + mention")
				}
			}
			if split[0] == "!helpme" {
				s.ChannelMessageSend(m.ChannelID, "```"+"!ascii1 + un text \n!ascii2 + un text \n!ascii3 + un text \n!ticstart \n!spam + mention"+"```")
			}
			if split[0] == "!ascii" {
				myString := strings.Join(split[1:], " ")

				file, _ := os.Open("standard.txt")
				fileVal := ScanFile(file)

				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner

			}
			if split[0] == "!ascii1" {
				myString := strings.Join(split[1:], " ")

				file, _ := os.Open("standard.txt")
				fileVal := ScanFile(file)

				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner

			}
			if split[0] == "!ascii2" {
				myString := strings.Join(split[1:], " ")

				file, _ := os.Open("thinkertoy.txt")
				fileVal := ScanFile(file)

				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner

			}

			if split[0] == "!ascii3" {

				myString := strings.Join(split[1:], " ")

				file, _ := os.Open("shadow.txt")
				fileVal := ScanFile(file)

				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner

			}
			if split[0] == "!tic1" || split[0] == "!tic2" || split[0] == "!tic3" || split[0] == "!tic4" || split[0] == "!tic5" || split[0] == "!tic6" || split[0] == "!tic7" || split[0] == "!tic8" || split[0] == "!tic9" || split[0] == "!ticstart" {
				if split[0] == "!ticstart" {
					tab.joueur = true
					s.ChannelMessageSend(m.ChannelID, "```"+"La partie commence"+"```")
					tab.tableau[0][0] = "!tic7"
					tab.tableau[0][1] = "!tic8"
					tab.tableau[0][2] = "!tic9"
					tab.tableau[1][0] = "!tic4"
					tab.tableau[1][1] = "!tic5"
					tab.tableau[1][2] = "!tic6"
					tab.tableau[2][0] = "!tic1"
					tab.tableau[2][1] = "!tic2"
					tab.tableau[2][2] = "!tic3"
					s.ChannelMessageSend(m.ChannelID, "```"+tab.tableau[0][0]+"|"+tab.tableau[0][1]+"|"+tab.tableau[0][2]+"\n"+tab.tableau[1][0]+"|"+tab.tableau[1][1]+"|"+tab.tableau[1][2]+"\n"+tab.tableau[2][0]+"|"+tab.tableau[2][1]+"|"+tab.tableau[2][2]+"```")
					tab.tableau[0][0] = " "
					tab.tableau[0][1] = " "
					tab.tableau[0][2] = " "
					tab.tableau[1][0] = " "
					tab.tableau[1][1] = " "
					tab.tableau[1][2] = " "
					tab.tableau[2][0] = " "
					tab.tableau[2][1] = " "
					tab.tableau[2][2] = " "
				}
				if split[0] == "!tic7" {
					if tab.tableau[0][0] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[0][0] = "O"

						} else {
							tab.joueur = true
							tab.tableau[0][0] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}

				}
				if split[0] == "!tic8" {
					if tab.tableau[0][1] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[0][1] = "O"
						} else {
							tab.joueur = true
							tab.tableau[0][1] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic9" {
					if tab.tableau[0][2] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[0][2] = "O"
						} else {
							tab.joueur = true
							tab.tableau[0][2] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic4" {
					if tab.tableau[1][0] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[1][0] = "O"
						} else {
							tab.joueur = true
							tab.tableau[1][0] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic5" {
					if tab.tableau[1][1] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[1][1] = "O"
						} else {
							tab.joueur = true
							tab.tableau[1][1] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic6" {
					if tab.tableau[1][2] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[1][2] = "O"
						} else {
							tab.joueur = true
							tab.tableau[1][2] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic1" {
					if tab.tableau[2][0] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[2][0] = "O"
						} else {
							tab.joueur = true
							tab.tableau[2][0] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic2" {
					if tab.tableau[2][1] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[2][1] = "O"
						} else {
							tab.joueur = true
							tab.tableau[2][1] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				if split[0] == "!tic3" {
					if tab.tableau[2][2] == " " {
						if tab.joueur == true {
							tab.joueur = false
							tab.tableau[2][2] = "O"
						} else {
							tab.joueur = true
							tab.tableau[2][2] = "X"
						}
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"La case est deja prise"+"```")
					}
				}
				s.ChannelMessageSend(m.ChannelID, "```"+tab.tableau[0][0]+"|"+tab.tableau[0][1]+"|"+tab.tableau[0][2]+"\n"+tab.tableau[1][0]+"|"+tab.tableau[1][1]+"|"+tab.tableau[1][2]+"\n"+tab.tableau[2][0]+"|"+tab.tableau[2][1]+"|"+tab.tableau[2][2]+"```")
				if verif() == 1 {
					if tab.joueur == true {
						s.ChannelMessageSend(m.ChannelID, "```"+"Bravo joueur 2 !!!"+"```")
					} else {
						s.ChannelMessageSend(m.ChannelID, "```"+"Bravo joueur 1 !!!"+"```")
					}
					tab.tableau[0][0] = " "
					tab.tableau[0][1] = " "
					tab.tableau[0][2] = " "
					tab.tableau[1][0] = " "
					tab.tableau[1][1] = " "
					tab.tableau[1][2] = " "
					tab.tableau[2][0] = " "
					tab.tableau[2][1] = " "
					tab.tableau[2][2] = " "
				}
				if verif() == 2 {
					s.ChannelMessageSend(m.ChannelID, "```"+"DRAW !!!"+"```")
					tab.tableau[0][0] = " "
					tab.tableau[0][1] = " "
					tab.tableau[0][2] = " "
					tab.tableau[1][0] = " "
					tab.tableau[1][1] = " "
					tab.tableau[1][2] = " "
					tab.tableau[2][0] = " "
					tab.tableau[2][1] = " "
					tab.tableau[2][2] = " "
				}
			}
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "il n'y a pas de message")
	}

}

func printLetter(s string, fileVal []string) string {
	k := ""

	for i := 1; i <= 8; i++ {
		for _, arg := range s {
			indexBase := int(rune(arg)-32) * 9
			if indexBase > 856 {
				k = "index out of range"
				return k
			} else {
				k += fileVal[indexBase+i]
			}

		}
		k += "\n"

	}
	return k
}
func ScanFile(arg *os.File) []string {
	var fileValue []string
	scanner := bufio.NewScanner(arg)
	for scanner.Scan() {
		lines := scanner.Text()
		fileValue = append(fileValue, lines)
	}
	return fileValue
}
func verif() int {
	if tab.tableau[0][0] == tab.tableau[1][0] && tab.tableau[1][0] == tab.tableau[2][0] && tab.tableau[2][0] != " " {
		return 1
	}
	if tab.tableau[0][0] == tab.tableau[0][1] && tab.tableau[0][1] == tab.tableau[0][2] && tab.tableau[0][2] != " " {
		return 1
	}
	if tab.tableau[1][0] == tab.tableau[1][1] && tab.tableau[1][1] == tab.tableau[1][2] && tab.tableau[1][2] != " " {
		return 1
	}
	if tab.tableau[2][0] == tab.tableau[2][1] && tab.tableau[2][1] == tab.tableau[2][2] && tab.tableau[2][2] != " " {
		return 1
	}
	if tab.tableau[0][1] == tab.tableau[1][1] && tab.tableau[1][1] == tab.tableau[1][2] && tab.tableau[1][2] != " " {
		return 1
	}
	if tab.tableau[0][2] == tab.tableau[1][2] && tab.tableau[1][2] == tab.tableau[2][2] && tab.tableau[2][2] != " " {
		return 1
	}
	if tab.tableau[0][0] == tab.tableau[1][1] && tab.tableau[1][1] == tab.tableau[2][2] && tab.tableau[2][2] != " " {
		return 1
	}
	if tab.tableau[0][2] == tab.tableau[1][1] && tab.tableau[1][1] == tab.tableau[2][0] && tab.tableau[2][0] != " " {
		return 1
	}
	if tab.tableau[0][0] != " " && tab.tableau[0][1] != " " && tab.tableau[0][2] != " " && tab.tableau[1][0] != " " && tab.tableau[1][1] != " " && tab.tableau[1][2] != " " && tab.tableau[2][0] != " " && tab.tableau[2][1] != " " && tab.tableau[2][2] != " " {
		return 2
	}
	return 0
}
func activity() {
}
