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

	// Add interaction handler
	dg.AddHandler(handleInteraction)

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

	// Register the slash command to display the modal
	registerModalSlashCommand(dg, guildID)

	// get channel ID
	channelID, err := getTestChannelID(dg, guildID)
	if err != nil {
		log.Println("Failed to get the test channel ID:", err)
		return
	}

	// send test channel
	_, err = dg.ChannelMessageSend(channelID, "Hello, Test Channel!")
	if err != nil {
		log.Println("Failed to send message to test channel:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {} // Wait indefinitely.
}

func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		showModal(s, i)
		return
	}

	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	// Handle modal submission
	if i.MessageComponentData().CustomID == "modal_submit" {
		// data := i.MessageComponentData()
		// userid := strings.Split(data.CustomID, "_")[2]

		// Do something with modal data here

		// Send thanks to user
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Thank you for filling out the modal!",
			},
		})
	}
}

func showModal(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Create the modal with input fields and a submit button
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "modals_survey_" + i.Interaction.Member.User.ID,
			Title:    "Modals survey",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "opinion",
							Label:       "What is your opinion on them?",
							Style:       discordgo.TextInputShort,
							Placeholder: "Don't be shy, share your opinion with us",
							Required:    true,
							MaxLength:   300,
							MinLength:   10,
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("Failed to send modal: %v", err)
	}
}

func registerModalSlashCommand(s *discordgo.Session, guildID string) {
	// Define the slash command
	cmd := &discordgo.ApplicationCommand{
		Name:        "show-modal",
		Description: "Show a test modal",
	}

	_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
	if err != nil {
		log.Println("Failed to create command:", err)
	}
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
