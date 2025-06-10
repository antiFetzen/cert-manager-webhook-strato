<p align="center">
    <img src="https://raw.githubusercontent.com/cert-manager/cert-manager/d53c0b9270f8cd90d908460d69502694e1838f5f/logo/logo-small.png" height="196" width="196" alt="cert-manager project logo" />
</p>
<p align="center">
    <img src="https://raw.githubusercontent.com/antiFetzen/cert-manager-webhook-strato/refs/heads/master/docs/images/strato-logo.svg" width="600" alt="cert-manager project logo" />
</p>


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

Generate package data for **go-embed-python**
```bash
go generate ./...
```

Download dependencies
```bash
go mod download
```



> TODO: content

```bash
$ TEST_ZONE_NAME=example.com. make test
```
