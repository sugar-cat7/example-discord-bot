package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	TOKEN := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// ギルドのIDを取得
	guildID, err := getFirstGuildID(dg)
	if err != nil {
		log.Println("Failed to get the first guild ID:", err)
		return
	}

	// チャンネルのIDを取得
	channelID, err := getTestChannelID(dg, guildID)
	if err != nil {
		log.Println("Failed to get the test channel ID:", err)
		return
	}

	// テスト用チャンネルにメッセージを送信
	embed := &discordgo.MessageEmbed{
		Title:       "Hello, World!",
		Description: "A description about the embed content.",
		URL:         "https://example.com",
		Timestamp:   "2022-03-01T15:04:05.000Z",
		Color:       0x00ff00,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Footer Text",
			IconURL: "https://3.bp.blogspot.com/-f7RP4FBbHHs/WKbKWONd1UI/AAAAAAABB0Y/xIurZIPrtPgV7X-yDbavc4v98DcYZ8onQCLcB/s800/character_computer_screen_hakase.png",
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://1.bp.blogspot.com/-hEH4NwnHbIw/U400-nkcbbI/AAAAAAAAg7c/lnPJPI4cOz0/s800/earth_good.png",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://3.bp.blogspot.com/-DAaMXKLWyTQ/XGjxyDlHrNI/AAAAAAABRbs/FJ4hpfKdzykk8PKjOByKTIuwpZ0URAwbwCLcBGAs/s800/earth_nature_futaba.png",
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Author's Name",
			URL:     "https://example.com/author",
			IconURL: "https://2.bp.blogspot.com/-Ten5Y3wa1s8/VMItaHv6ikI/AAAAAAAAqtU/HVC0kvCwPYo/s800/character_hakase.png",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Field 1",
				Value:  "Value 1",
				Inline: true,
			},
			{
				Name:   "Field 2",
				Value:  "Value 2",
				Inline: true,
			},
		},
	}
	_, err = dg.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Println("Failed to send message to test channel:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {} // Wait indefinitely.
}

func getFirstGuildID(s *discordgo.Session) (string, error) {
	guilds, err := s.UserGuilds(1, "", "")
	if err != nil {
		return "", fmt.Errorf("failed to get guilds: %w", err)
	}

	if len(guilds) == 0 {
		return "", fmt.Errorf("bot is not a member of any guilds")
	}

	return guilds[0].ID, nil
}

func getTestChannelID(s *discordgo.Session, guildID string) (string, error) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return "", fmt.Errorf("failed to get channels for guild %s: %w", guildID, err)
	}

	for _, ch := range channels {
		if ch.Name == "テスト用" {
			return ch.ID, nil
		}
	}

	return "", fmt.Errorf("no channel named 'テスト用' found")
}
