## File Message

- ファイルを添付しメッセージを送信する

![Alt text](images/README/image.png)

```go
dg, err := discordgo.New("Bot " + TOKEN)
...
resp, err := http.Get("")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

// ファイルをメッセージとして送信
message := &discordgo.MessageSend{
    Content: "画像:",
    Files: []*discordgo.File{
        {
            Name:   "sugar.png",
            Reader: resp.Body,
        },
    },
}

_, err = dg.ChannelMessageSendComplex(channelID, message)

```

- 添付できるファイルの種類ですが io.Reader の interface を実装しているものであれば、どのような形式のものでも添付することができます。

```go
type File struct {
	Name        string
	ContentType string
	Reader      io.Reader
}
```

基本的な形式はサポートされており、以下のようなものが添付できます。

```text
画像: .jpg, .jpeg, .png, .gif, .webp
動画: .mp4, .mov
オーディオ: .mp3, .wav
テキスト: .txt, .md, .log
ドキュメント: .pdf, .doc, .docx, .ppt, .pptx, .xls, .xlsx
圧縮ファイル: .zip, .rar, .tar.gz
```

※1 メッセージに対しては最大 8MB までのファイルを添付することができます。
