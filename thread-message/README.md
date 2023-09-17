## Thread Text Message

- メッセージに対してスレッドを作成する

![Alt text](images/README/image.png)

```go
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
```
