package modules

import (
	"fmt"
	"leviathanRewritten/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var last_results []utils.Result // guardar os resultados obtidos anteriormente
var last_msg_author_id string   // lembrar quem que pesquisou anteriormente

func Google(s *discordgo.Session, m *discordgo.Message, args ...string) {
	if len(args) > 0 {
		query := utils.ArgsTag(args)
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			return
		}

		query = strings.Replace(query, "_", "%20", -1)

		res, err := utils.GoogleParse(query, channel.NSFW)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//remove bad args
		if len(res) > 0 {
			last_results = res
			last_msg_author_id = m.Author.ID

			substr := strings.SplitAfter(res[0].Link, "&sa")
			final := strings.Replace(substr[0], "&sa", "", -1)
			sent_msg, err := s.ChannelMessageSend(m.ChannelID, final)

			if err != nil {
				// erro ao enviar mensagem
				fmt.Println(err.Error())
				return
			}

			// adicionar reacts na mensagem
			s.MessageReactionAdd(m.ChannelID, sent_msg.ID, "◀️")
			s.MessageReactionAdd(m.ChannelID, sent_msg.ID, "▶️")
		} else {
			msg := utils.NewEmbed()
			msg.SetColor(utils.Yellow)
			msg.SetTitle("Erro")
			msg.SetDescription("Nenhum resultado encontrado")
			s.ChannelMessageSendEmbed(m.ChannelID, msg.MessageEmbed)
		}

	}

	return

	//fmt.Println(teste[1].Link)
}
