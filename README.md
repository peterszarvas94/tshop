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
tshop address create --name "John Doe" --country "US" --province "Montana" --city "Bozeman" --zip "59715" --street1 "123 Main Street" --street2 "Apt 4B" --phone "406-555-1234"
tshop address delete "shp_xxx"
tshop address list
```

### Card

```bash
tshop card create
tshop card delete "crd_xxx"
tshop card list
```

### Shopping cart

```bash
tshop cart address "shp_xxx"
tshop cart card "crd_xxx"
tshop cart clear
tshop cart info
tshop cart add --variant "var_xxx" --quantity "1"
tshop cart remove --variant "var_xxx" --quantity "1"
tshop cart order
```

### Order

```bash
tshop order info "ord_xxx"
tshop order list
```

### Product

```bash
tshop product info "prd_xxx"
tshop product list
```

### Subscriptions

> [!CAUTION]
> Not tested in production.

```bash
tshop subscription cancel
tshop subscription create --address "shp_xxx" --card "crd_xxx" --variant "var_xxx" --quantity "1" --type "weekly" --interval "3"
tshop subscription info "sub_xxx"
tshop subscription list
```

### Token

> [!WARNING]
> If you delete the token you are using, you need to confirm it.

```bash
tshop token create
tshop token delete "pat_xxx"
tshop token list
```

### User

```bash
tshop user info
tshop user update --name "John Doe" --email "johndoe@terminal.shop"
```

### Version

```bash
tshop version
```
