# tshop

The essence of the [terminal](https://www.terminal.shop) shopping experience - a command-line interface

## Install

```bash
go install github.com/peterszarvas94/tshop@v1.0.0
```

## Env

You should set up the following env variables:

```sh
TERMINAL_TOKEN_ID=pat_...
TERMINAL_TOKEN=trm_...
TERMINAL_ENV=dev/prod
```

> [!NOTE]
> If you don't have valid tokens, visit `ssh terminal.shop -t token` for prod, or `ssh dev.terminal.shop -t token` for dev

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

```bash
tshop subscription cancel
tshop subscription create --address "shp_xxx" --card "crd_xxx" --variant "var_xxx" --quantity "1" --type "weekly" --interval "3"
tshop subscription info "sub_xxx"
tshop subscription list
```

> [!CAUTION]
> Not tested in production.

### Token

```bash
tshop token create
tshop token delete "pat_xxx"
tshop token list
```

> [!WARNING]
> If you delete the token you are using, you need to confirm it.

### User

```bash
tshop user info
tshop user update --name "John Doe" --email "johndoe@terminal.shop"
```

### Version

```bash
tshop version
```
