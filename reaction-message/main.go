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

	message, _ := dg.ChannelMessageSend(channelID, "React to this message!")
	err = dg.MessageReactionAdd(channelID, message.ID, "😀")
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
