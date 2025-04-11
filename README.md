# tshop

CLI for terminal.shop

## Inspiration

The `ssh terminal.shop` experience is too bloated.

I present you the _real_ terminal shopping experience. Only pure text, no fancy frontend.

## Install

```bash
go install github.com/peterszarvas94/tshop@v0.0.1
```

## Env

You should set up the following env variables:

> [!NOTE]
> If you don't have valid tokens, visit `ssh terminal.shop -t token` (or `ssh dev.terminal.shop -t token` for dev server)

```sh
TERMINAL_TOKEN_ID=pat_...
TERMINAL_TOKEN=trm_...
TERMINAL_ENV=dev/prod
```

## Commands

More info about commands and subcommands are available with the `--help` flag.

```bash
tshop --help
tshop address --help
tshop address create --help
```

### Address

```bash
tshop address
tshop address create
tshop address delete
tshop address list
```

### Card

```bash
tshop card
tshop card create
tshop card delete
tshop card list
```

### Shopping cart

```bash
tshop cart
tshop cart address
tshop cart card
tshop cart clear
tshop cart info
tshop cart update
tshop cart order
```

### Order

```bash
tshop order
tshop order info
tshop order list
```

### Product

```bash
tshop product
tshop product info
tshop product list
```

### Subscriptions

> [!CAUTION]
> Not tested in production. It possibly needs future improvement..

```bash
tshop subscription
tshop subscription cancel
tshop subscription create
tshop subscription info
tshop subscription list
```

### Token

> [!TIP]
> If you delete the token you are using, you need to confirm it.

```bash
tshop token
tshop token create
tshop token delete
tshop token list
```

### User

```bash
tshop user
tshop user info
tshop user update
```

### Version

```bash
tshop version
```
