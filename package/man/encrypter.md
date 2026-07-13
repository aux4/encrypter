#### Description

The `encrypter` command is the entry point for the aux4 encryption facade. It groups the `encrypt`, `decrypt`, and `generate-secret` commands and a `provider` namespace under which provider packages register.

The facade contains no cryptography of its own — it routes each command to the selected provider (`aux4 encrypter provider <provider> ...`). Install a provider package to enable encryption:

```bash
aux4 aux4 pkger install aux4/encrypter-aux4
```

The active provider is resolved as: `--provider` flag > `AUX4_ENCRYPTER_PROVIDER` env > `--config` value > default `aux4`.

#### Usage

```bash
aux4 encrypter <command> [options]
```

encrypt          Encrypt plaintext using the selected provider
decrypt          Decrypt base64 ciphertext using the selected provider
generate-secret  Generate a random secret
provider         Provider namespace (aux4 encrypter provider <name> ...)

#### Example

```bash
SECRET=$(aux4 encrypter generate-secret)
echo -n 'hello world' | aux4 encrypter encrypt --secret "$SECRET"
```

```text
AZ1nQ2...base64...==
```
