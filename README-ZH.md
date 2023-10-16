# GoToPost - gtp to Postman

## 動機

在最開始時，我只是想要寫一個小工具來自動幫我完成Postman的route，畢竟api一多起來其實有點麻煩。
但我突然想到，何不把這個小工具開源並讓大家使用呢？
所以這個專案就誕生了。

## 使用方法

cd 到router.go的資料夾

```bash
gtp <arguments> <ip> <CollectionName>

-t                  convert to thunder-client
-p                  convert to postman
```

## 安裝方法

```bash
curl -sSfL https://raw.githubusercontent.com/Rudy1021/GoToPost/master/install.sh | sh
```

## Q&A

### "command not found: gtp" or "No such file or directory"

```zsh
alias gtp="$GOPATH/bin/gtp"
```

### 請注意 此指令要在$GOPATH的下方
