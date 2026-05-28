# Chip-in Calculator

## Flags

| Flag             | Required | Default       | Description                                 |
| ---------------- | -------- | ------------- | ------------------------------------------- |
| `-r`, `--rate`   | Yes      | —             | EUR to UAH rate, e.g. 44.12                 |
| `-c`, `--config` | No       | `config.json` | Path to config file                         |
| `--mentions`     | No       | —             | Mentions to put into message after greeting |

## Config

See [`config.schema.json`](config.schema.json) for the full schema. Example:

```json
{
  "subscriptions": [
    { "name": "Spotify", "price": 21.99, "membersCount": 5 },
    { "name": "Netflix", "price": 21.99, "membersCount": 3 }
  ]
}
```

## Docker

The image does not include a config file. Mount your own `config.json` at `/config.json`.

```bash
docker run --rm \
  -v $(pwd)/config.json:/config.json \
  -e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
  -e TELEGRAM_CHAT_ID=123456 \
  ghcr.io/alexsoft/chip-in-calculator:1.0.1 --rate 45.33

# With mentions
docker run --rm \
  -v $(pwd)/config.json:/config.json \
  -e TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
  -e TELEGRAM_CHAT_ID=123456 \
  ghcr.io/alexsoft/chip-in-calculator:1.0.1 --rate 45.33 --mentions "@user1 @user2"
```

## Build locally, Go 1.26+

```bash
go build ./ && ./chip-in-calculator --rate 45.33

# With Telegram
TELEGRAM_BOT_API_KEY=1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \
TELEGRAM_CHAT_ID=123456 \
./chip-in-calculator --rate 45.33
```
