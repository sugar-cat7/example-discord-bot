## Reaction To Text Message

- メッセージに対してリアクションを付ける

![Alt text](images/README/image.png)

![Alt text](images/README/image-2.png)

- リアクションとしてつけられる絵文字は 2 種類あります。
  - `Unicode絵文字`, `カスタム絵文字`

```go
// Unicode絵文字
err = dg.MessageReactionAdd(channelID, message.ID, "😀")

// カスタム絵文字
err = dg.MessageReactionAdd(channelID, message.ID, "emoji_name:emoji_id")
```

- emoji_id はスタンプのリンクから得られる`XXX`の部分です。
  - `https://cdn.discordapp.com/emojis/?XXX.webp`
    ![Alt text](images/README/image-1.png)
