# GoToPost - gtp to Postman

## Motivation

At first, I just wanted to create a small utility to automatically convert the content of router.go into the format compatible with Postman. However, as I thought about it, why not open-source it and make it available for everyone? Even allowing others to fork and use it with their unique routers. And that's how this tool was born.

## Usage

cd to the router.go directory.

```bash
gtp <arguments> <ip> <CollectionName>

-t                  convert to thunder-client
-p                  convert to postman
```

## Installation

```bash
curl -sSfL https://raw.githubusercontent.com/Rudy1021/GoToPost/master/install.sh | sh
```

## Q&A

### "command not found: gtp" or "No such file or directory"

```zsh
alias gtp="$GOPATH/bin/gtp"
```

### make sure this command is under the $GOPATH
