#### Description

The `generate-secret` command generates a random secret using the selected encryption provider and prints it to stdout. It is typically used to produce AES key material for the `aux4` provider.

The provider is resolved as: `--provider` flag > `AUX4_ENCRYPTER_PROVIDER` env > `--config` value > default `aux4`. Some providers (e.g. KMS-backed ones) may not implement this command. If the selected provider is not installed, the command prints a clean hint — `Encryption provider '<provider>' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-<provider>` — instead of a generic command-not-found error.

#### Usage

```bash
aux4 encrypter generate-secret [length] [--provider <name>]
```

--provider  The encryption provider (default: aux4; env: AUX4_ENCRYPTER_PROVIDER)
length      The length of the secret (positional arg; default: 32)

#### Example

```bash
aux4 encrypter generate-secret 32
```

```text
9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
```
