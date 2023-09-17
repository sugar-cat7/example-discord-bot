## Simple Text Message

- 単純なテキスト形式のメッセージを送信する

![Alt text](images/README/image.png)

```go
dg, err := discordgo.New("Bot " + TOKEN)
...
_, err = dg.ChannelMessageSend(channelID, "Hello, Test Channel!")
```
