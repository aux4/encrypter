#### Description

The `decrypt` command decrypts a base64-encoded ciphertext using the selected encryption provider and prints the recovered plaintext to stdout.

The ciphertext is taken from the positional `text` argument if provided, otherwise it is read from stdin — so both `aux4 encrypter decrypt --secret KEY "<base64>"` and `echo -n "<base64>" | aux4 encrypter decrypt --secret KEY` work. This matches how aux4 core and other callers pass ciphertext as a positional argument.

On a wrong key, corrupt input, or tampering, the command exits non-zero with a clean error message on stderr (no panic or stack trace).

`--secret` is **optional and never prompts**. The facade always forwards `--secret` (even empty) to the provider, so the provider decides whether a key is required and reports a clean error — the facade itself never blocks on a prompt. This lets keyless providers (e.g. KMS) decrypt without any key configured.

**Note:** When `--secret` and `AUX4_SECRET` are both omitted entirely, the fully non-prompting behavior depends on the ENC-006 core fix. On an older core an omitted `secret` can trigger a variable-resolution stack overflow; pass `--secret` explicitly (even `''`) to avoid it until the core fix is released.

The provider is resolved as: `--provider` flag > `AUX4_ENCRYPTER_PROVIDER` env > `--config` value > default `aux4`. If the selected provider is not installed, the command prints a clean hint — `Encryption provider '<provider>' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-<provider>` — instead of a generic command-not-found error.

#### Usage

```bash
aux4 encrypter decrypt --secret <keyInput> [text] [--provider <name>]
```

--secret    The key input, interpreted by the provider (optional; never prompts; env: AUX4_SECRET)
--provider  The encryption provider (default: aux4; env: AUX4_ENCRYPTER_PROVIDER)
text        The base64 ciphertext to decrypt (positional arg; falls back to stdin)

#### Example

```bash
echo -n 'AZ1nQ2f8...base64...==' | aux4 encrypter decrypt --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
```

```text
hello world
```
