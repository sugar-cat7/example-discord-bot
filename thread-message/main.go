package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	TOKEN := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.Contains(m.Content, "ping") {
			thread, err := s.MessageThreadStartComplex(m.ChannelID, m.ID, &discordgo.ThreadStart{
				Name:                "Pong game with " + m.Author.Username,
				AutoArchiveDuration: 60,
				Invitable:           false,
				RateLimitPerUser:    10,
			})
			if err != nil {
				log.Println("Error starting thread:", err)
				return
			}
			_, err = s.ChannelMessageSend(thread.ID, "pong")
			if err != nil {
				log.Println("Error sending message to thread:", err)
			}
			go func() {
				<-time.After(10 * time.Second) // Here you set the time after which the thread will be archived.
				archived := true
				locked := true
				_, err := s.ChannelEditComplex(thread.ID, &discordgo.ChannelEdit{
					Archived: &archived,
					Locked:   &locked,
				})
				if err != nil {
					log.Println("Error archiving thread:", err)
				}
			}()
		}
	})

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
	_, err = dg.ChannelMessageSend(channelID, "ping")
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
