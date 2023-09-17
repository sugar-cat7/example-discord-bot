## Voice Channel Audio

- ボイスチャンネルで効果音を送信する
  - - `dca`フォーマットのオーディオファイルを読み込む

<video width="320" height="240" controls>
  <source src="./movies/video.mov" type="video/mp4">
</video>

```go
func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}
	defer vc.Disconnect()

	time.Sleep(250 * time.Millisecond)
	vc.Speaking(true)
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}
	vc.Speaking(false)
	time.Sleep(250 * time.Millisecond)

	return nil
}
```
