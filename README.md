![Logo - cert-manager-webhook-strato](/docs/images/cert-manager-webhook-strato-logo.svg)

# STRATO DNS Webhook for cert-manager

---

This [cert-manager](https://cert-manager.io/) webhook implementation enables you to solve the [ACME DNS-01 challenges](https://cert-manager.io/docs/configuration/acme/dns01/#configuring-dns01-challenge-provider) for **STRATO**.

> [!NOTE]
> Under the hood, the webhook utilizes [strato-certbot](https://github.com/Buxdehuda/strato-certbot) from [Buxdehuda](https://github.com/Buxdehuda). Thanks for the great work 🙏
> 
> Git submodule is used to include [strato-certbot](https://github.com/Buxdehuda/strato-certbot).
> 
> To execute the python code [go-embed-python](https://github.com/kluctl/go-embed-python) is used.


# Usage

---

> TODO: content

# Installation

---





> TODO: content


# Testing

---

## Run Tests

1. Add your credential within `testdata/strato/config.json` for more details see the `README.md` in the directory.



2. Download modules
```bash
go mod download
```

3. Generate python binaries for **go-embed-python**
```bash
go generate ./...
```

4. Run test
```bash
$ TEST_ZONE_NAME=example.com. make test
```
