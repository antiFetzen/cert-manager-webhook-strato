> [!CAUTION]
> This project is under development and could change frequently

---

![Logo - cert-manager-webhook-strato](/docs/images/cert-manager-webhook-strato-logo.svg)

# STRATO DNS Webhook for cert-manager

This [cert-manager webhook(https://cert-manager.io/docs/configuration/acme/dns01/#configuring-dns01-challenge-provider) implementation enables you to solve the **ACME DNS-01 challenges** for **STRATO** 

> [!NOTE]
> Under the hood, the webhook utilizes [strato-certbot](https://github.com/Buxdehuda/strato-certbot) from [Buxdehuda](https://github.com/Buxdehuda). Thanks for the great work 🙏
> 
> Git submodule is used to include [strato-certbot](https://github.com/Buxdehuda/strato-certbot).
> 
> To execute the python code [go-embed-python](https://github.com/kluctl/go-embed-python) is used.


# Config

## Secret
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: strato-secret
  namespace: cert-manager
type: Opaque
data:
  username: your-username-base64-encoded
  password: your-password-base64-encoded
  # In case of 2FA provide both totp values or none
  totpDeviceName: your-totp-device-name-base64-encoded
  totpSecret: your-totp-secret-base64-encoded
```

## ClusterIssuer
```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-dns-strato
spec:
  acme:
    email: anti.fetzen@gmail.com
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: fetzen-rocks-issuer-account-key
    solvers:
      - dns01:
          webhook:
            groupName: fetzen.rocks
            solverName: acme
            config:
              # !IMPORTANT: use the correct url from your country
              # e.g.:
              #   Germany:      https://www.strato.de/apps/CustomerService
              #   Netherlands:  https://www.strato.nl/apps/CustomerService#skl
              apiUrl: https://www.strato.de/apps/CustomerService
              # Reference of the secret
              secretRef: strato-secret
              # OPTIONAL - default value is `ResourceNamespace`
              secretNamespace: cert-manager
              # OPTIONAL - default value is `0`
              waitingTime: 0
```

# Usage

> TODO: add content


# Testing

## Run Tests

It is very easy to test the webhook from your local computer to check the general functionality of the module.


### 1. Modify config
Add your credential within `testdata/strato/config.json` for more details see the `README.md` in the directory.

### 2. Download modules
```bash
go mod download
```

### 3. Generate python binaries for **go-embed-python**
```bash
go generate ./...
```

### 4. Run test
```bash
$ TEST_ZONE_NAME=example.com. make test
```

# Docker Build

```bash
docker build -t antifetzen/cert-manager-webhook-strato:latest --platform=linux/amd64 .
```

```bash
docker push antifetzen/cert-manager-webhook-strato:latest
```
