package commands

import (
	"leviathanRewritten/utils"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func CommandDiceExec(s *discordgo.Session, m *discordgo.Message, args ...string) {
	if len(args) == 0 {
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	if num <= 1 {
		return
	}

	dice := utils.RandomRange(1, num)
	msg := utils.NewEmbed()
	msg.SetAuthor(m.Author.AvatarURL("1024"), m.Author.Username)
	msg.SetDescription(strconv.Itoa(dice))
	sent, _ := s.ChannelMessageSendEmbed(m.ChannelID, msg.MessageEmbed)
	lastCommandOutputMsgChannelID = sent.ChannelID
	lastCommandOutputMsgID = sent.ID
	return
}
