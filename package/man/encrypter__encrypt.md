#### Description

The `encrypt` command encrypts plaintext using the selected encryption provider and prints a base64-encoded ciphertext to stdout.

The plaintext is taken from the positional `text` argument if provided, otherwise it is read from stdin — so both `aux4 encrypter encrypt --secret KEY "text"` and `echo -n "text" | aux4 encrypter encrypt --secret KEY` work.

Output is **plain base64** with no prefix and no `--iv`. Any nonce/IV rides inside the blob (the exact format is a provider detail).

`--secret` is **optional and never prompts**. The facade always forwards `--secret` (even empty) to the provider, so the provider decides whether a key is required and reports a clean error — the facade itself never blocks on a prompt.

The provider is resolved as: `--provider` flag > `AUX4_ENCRYPTER_PROVIDER` env > `--config` value > default `aux4`. If the selected provider is not installed, the command prints a clean hint — `Encryption provider '<provider>' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-<provider>` — instead of a generic command-not-found error.

#### Usage

```bash
aux4 encrypter encrypt --secret <keyInput> [text] [--provider <name>]
```

--secret    The key input, interpreted by the provider (optional; never prompts; env: AUX4_SECRET)
--provider  The encryption provider (default: aux4; env: AUX4_ENCRYPTER_PROVIDER)
text        The plaintext to encrypt (positional arg; falls back to stdin)

#### Example

```bash
echo -n 'hello world' | aux4 encrypter encrypt --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
```

```text
AZ1nQ2f8...base64...==
```
