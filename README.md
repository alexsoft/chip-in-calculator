# Chip-in Calculator

## How to run

### Flags

| Flag             | Required | Default       | Description                                 |
| ---------------- | -------- | ------------- | ------------------------------------------- |
| `-r`, `--rate`   | Yes      | —             | EUR to UAH rate, e.g. 44.12                 |
| `-c`, `--config` | No       | `config.json` | Path to config file                         |
| `--mentions`     | No       | —             | Mentions to put into message after greeting |

### Build docker image

```bash
docker build -t chip-in-calculator . \
&& docker run --rm chip-in-calculator -r 45.33

docker build -t chip-in-calculator . \
&& docker run --rm chip-in-calculator --rate 47.11

# With custom config file
docker build -t chip-in-calculator . \
&& docker run --rm -v $(pwd)/config2.json:/config2.json \
chip-in-calculator -r 45.33 -c config2.json
```

#### Send message to Telegram

```bash
docker build -t chip-in-calculator . \
&& docker run --rm \
-e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
-e TELEGRAM_CHAT_ID=123456 \
chip-in-calculator --rate 45.33

# With mentions
docker build -t chip-in-calculator . \
&& docker run --rm \
-e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
-e TELEGRAM_CHAT_ID=123456 \
chip-in-calculator --rate 45.33 --mentions "@user1 @user2"
```

### Build locally, Go 1.25+

```bash
go build ./ \
&& ./chip-in-calculator -r 45.33

go build ./ \
&& ./chip-in-calculator --rate 47.11

# With custom config file
go build ./ \
&& ./chip-in-calculator -r 45.33 -c config2.json
```

#### Send message to Telegram
```bash
go build ./ \
&& TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
TELEGRAM_CHAT_ID=123456 \
./chip-in-calculator --rate 47.11

# With mentions
go build ./ \
&& TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
TELEGRAM_CHAT_ID=123456 \
./chip-in-calculator --rate 47.11 --mentions "@user1 @user2"
```
