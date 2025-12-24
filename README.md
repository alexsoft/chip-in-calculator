# Chip-in Calculator

## How to run

### Required flags

| Flag           | Description                 |
| -------------- | --------------------------- |
| `-r`, `--rate` | EUR to UAH rate, e.g. 44.12 |

### Use docker image

```bash
docker run --rm \
-e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
-e TELEGRAM_CHAT_ID=123456 \
IMAGE \
-r 45.33
```

#### Providing your config
There is a default config from repo copied inside the image. You can override it by mounting your own config file.

```bash
docker run --rm \
-v /path/to/config:/config.json \
-e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
-e TELEGRAM_CHAT_ID=123456 \
IMAGE \
-r 45.33
```

### Build docker image

```bash
docker build -t chip-in-calculator . \
&& docker run --rm chip-in-calculator -r 45.33

docker build -t chip-in-calculator . \
&& docker run --rm chip-in-calculator --rate 47.11
```

#### Send message to Telegram

```bash
docker build -t chip-in-calculator . \
&& docker run --rm \
-e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
-e TELEGRAM_CHAT_ID=123456 \
chip-in-calculator --rate 45.33
```

### Build locally, Go 1.25+

```bash
go build ./ \
&& ./chip-in-calculator -r 45.33

go build ./ \
&& ./chip-in-calculator --rate 47.11
```

#### Send message to Telegram
```bash
go build ./ \
&& TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
TELEGRAM_CHAT_ID=123456 \
./chip-in-calculator --rate 47.11
```
