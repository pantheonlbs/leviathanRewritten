package modules

import (
	"leviathanRewritten/utils"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Dice(s *discordgo.Session, m *discordgo.Message, args ...string) {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		utils.SendWarning("Error", s, m)
		return
	}

	dice := utils.RandomRange(1, num)
	msg := utils.NewEmbed()
	msg.SetAuthor(m.Author.AvatarURL("1024"), m.Author.Username)
	msg.SetDescription(strconv.Itoa(dice))
	return
}
