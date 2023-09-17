## Text To Speech Message

- Text To Speech (TTS) 機能を利用して、テキストメッセージを音声として再生する。

<video width="320" height="240" controls>
  <source src="./movies/tts.mov" type="video/mp4">
</video>

```go
dg, err := discordgo.New("Bot " + TOKEN)
...
_, err = dg.ChannelMessageSendTTS(channelID, "Hello in Text-to-Speech!")
```

- ユーザーが TTS を無効にしている場合、メッセージは通常のテキストとして表示される。
- TTS の音声のピッチや速度などの細かい設定をカスタマイズすることはできない。
- TTS メッセージはテキストチャンネルでのみ再生される。
