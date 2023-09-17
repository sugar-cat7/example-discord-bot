package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
)

// DiscordのパケットをPion RTPパケットに変換
func createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version:        2,
			PayloadType:    0x78, // Discord voiceのドキュメントから取得
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}

// 音声データの取り扱い（保存）
func handleVoice(c chan *discordgo.Packet) {
	files := make(map[uint32]media.Writer) // SSRCをキーとしたファイルのマップ

	for p := range c {
		file, ok := files[p.SSRC]
		if !ok {
			// 新しいOGGファイルの作成
			var err error
			file, err = oggwriter.New(fmt.Sprintf("%d.ogg", p.SSRC), 48000, 2)
			if err != nil {
				fmt.Printf("failed to create file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
				return
			}
			files[p.SSRC] = file
		}
		// DiscordGoの型からpion RTPパケットを構築
		rtp := createPionRTPPacket(p)
		err := file.WriteRTP(rtp)
		if err != nil {
			fmt.Printf("failed to write to file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
		}
	}

	// パケットの受信が終了したら全てのファイルを閉じる
	for _, f := range files {
		f.Close()
	}
}

func main() {
	TOKEN := os.Getenv("DISCORD_TOKEN")

	// Discordセッションの作成
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// セッションを開始
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// 最初のギルドIDを取得
	guildID, err := getFirstGuildID(dg)
	if err != nil {
		log.Println("Failed to get the first guild ID:", err)
		return
	}

	// テスト用のボイスチャンネルIDを取得
	channelID, err := getTestChannelID(dg, guildID)
	if err != nil {
		log.Println("Failed to get the test channel ID:", err)
		return
	}

	// ボイスチャンネルに参加
	v, err := dg.ChannelVoiceJoin(guildID, channelID, true, false)
	if err != nil {
		fmt.Println("failed to join voice channel:", err)
		return
	}

	// 10秒後にボイスチャンネルを退出
	go func() {
		time.Sleep(10 * time.Second)
		close(v.OpusRecv)
		v.Close()
	}()

	// 音声データを取り扱う
	handleVoice(v.OpusRecv)

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {} // 無限に待機
}

// ボットが参加している最初のギルドIDを取得
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

// テスト用のボイスチャンネルIDを取得
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
