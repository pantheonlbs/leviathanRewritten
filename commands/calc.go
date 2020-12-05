package commands

import (
	"fmt"
	"leviathanRewritten/utils"
	"strings"

	"net/url"

	"github.com/bwmarrin/discordgo"
)

// CommandCalcExec representa o comando de calculadora
func CommandCalcExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	expr := strings.TrimSpace(strings.Join(args, " "))

	responseBody, err := utils.GetDoc("https://api.mathjs.org/v4/?expr=" + url.QueryEscape(expr))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := string(responseBody)

	// deixar resultado bonitinho
	result = strings.ReplaceAll(result, "Infinity", "♾️")
	result = strings.ReplaceAll(result, "Error:", "❌")
	result = strings.ToLower(result)

	// deixar expressão bonitinha
	title := expr
	title = strings.ReplaceAll(title, "*", "×")
	title = strings.ReplaceAll(title, "-", "−")

	e := utils.NewEmbed()

	e.SetTitle(title)
	e.SetColor(0x36393F)
	e.SetDescription(result)

	s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
}
