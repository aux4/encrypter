# aux4 encrypter

Encryption facade for aux4. This package provides the `encrypter` command group (`encrypt`, `decrypt`, `generate-secret`) and dispatches to provider packages like `aux4/encrypter-aux4` that implement the actual cryptography.

The facade contains **no crypto** — it only routes. Install the provider you want; it pulls this facade in via `dependencies` (just like installing `aux4/secret-1password` pulls `aux4/secret`).

## Installation

Install a provider — it brings the facade with it:

```bash
aux4 aux4 pkger install aux4/encrypter-aux4
```

## Quick Start

```bash
# Generate a key
SECRET=$(aux4 encrypter generate-secret)

# Encrypt (plaintext via stdin) -> base64 on stdout
echo -n 'hello world' | aux4 encrypter encrypt --secret "$SECRET"

# Decrypt (base64 via stdin OR positional arg) -> plaintext on stdout
echo -n '<base64>' | aux4 encrypter decrypt --secret "$SECRET"
aux4 encrypter decrypt --secret "$SECRET" '<base64>'
```

## Provider Selection

The active provider is resolved in this order (first match wins):

1. `--provider <name>` flag
2. `AUX4_ENCRYPTER_PROVIDER` environment variable
3. `--config` value (via `aux4/config`)
4. Default: `aux4`

```bash
aux4 encrypter encrypt --provider aux4 --secret "$SECRET" 'text'
AUX4_ENCRYPTER_PROVIDER=aws-kms aux4 encrypter encrypt --secret alias/my-key 'text'
```

If the selected provider is not installed, the facade prints a clean install hint instead of a generic command-not-found error:

```text
Encryption provider 'aws-kms' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-aws-kms
```

```bash
aux4 aux4 pkger install aux4/encrypter-<provider>
```

## The `--secret` Flag

`--secret` is **optional and never prompts** on `encrypt` and `decrypt`. The facade always forwards `--secret`
(even empty) to the selected provider, so the provider — not the facade — decides whether a key is required
and reports a clean error. This is what lets keyless providers (e.g. `aws-kms`, whose ciphertext blob is
self-describing) decrypt without any key configured.

```bash
# aux4 provider: a key is required — an empty/wrong key yields a clean "wrong key or corrupt data" error
aux4 encrypter decrypt --secret "$SECRET" '<base64>'

# KMS provider: no key needed on decrypt
AUX4_ENCRYPTER_PROVIDER=aws-kms aux4 encrypter decrypt '<base64>'
```

**Note:** Fully non-prompting behavior when `--secret` and `AUX4_SECRET` are both omitted requires the ENC-006
core fix (unreleased at the time of writing). On an older core an omitted `secret` can trigger a
variable-resolution stack overflow; passing `--secret` explicitly (even `''`) avoids it until the core fix
ships.

## Commands

| Command | Description |
|---------|-------------|
| `encrypt` | Encrypt plaintext (positional arg or stdin) → base64 on stdout |
| `decrypt` | Decrypt base64 (positional arg or stdin) → plaintext on stdout |
| `generate-secret` | Generate a random secret (provider-specific) |
| `provider` | Namespace under which providers register (`aux4 encrypter provider <name> ...`) |

## Provider Command Contract

Every encrypter provider registers under the shared `encrypter:provider` profile with a `<provider>` command that routes to its own `encrypter:provider:<provider>` sub-profile. The commands and their interfaces are the same across all providers, so the facade can dispatch to any of them uniformly.

### encrypt (required)

Encrypts plaintext to a base64 blob.

```bash
aux4 encrypter provider <provider> encrypt --secret <keyInput> [text]
```

| Variable | Description | Source |
|----------|-------------|--------|
| `secret` | The key input, interpreted by the provider (AES key material, KMS alias/ARN, or a `secret://` reference resolved by core) | `--secret` flag |
| `text` | The plaintext | positional arg **or** stdin |

- Plaintext is read from the **positional arg** if present, otherwise from **stdin**.
- Output is **plain base64** on stdout — **no `--iv`**, **no value prefix** (no `aux4enc:` string). The nonce/IV, if any, rides inside the blob.

### decrypt (required)

Decrypts a base64 blob back to plaintext.

```bash
aux4 encrypter provider <provider> decrypt --secret <keyInput> [base64]
```

| Variable | Description | Source |
|----------|-------------|--------|
| `secret` | The key input, interpreted by the provider | `--secret` flag |
| `text` | The base64 ciphertext | positional arg **or** stdin |

- Ciphertext is read from the **positional arg** if present, otherwise from **stdin**.
- Output is the recovered plaintext on stdout.
- On a wrong key, corrupt input, or tamper, the provider must exit **non-zero** with a clean stderr message — never a panic or stack trace.

### generate-secret (optional)

Generates a random secret suitable for the provider.

```bash
aux4 encrypter provider <provider> generate-secret [length]
```

| Variable | Description | Default |
|----------|-------------|---------|
| `length` | Length of the generated secret | `32` |

Some providers (e.g. KMS-backed ones) may omit or no-op this command.

### The `--secret` value

`--secret` is **provider-interpreted key input**, not a fixed key format:

- The `aux4` provider treats it as AES key material.
- A KMS provider treats it as a key alias/ARN (needed on encrypt; decrypt is self-describing).
- The value may be a literal, an environment value, or a `secret://...` reference — core resolves `secret://` via the vault mechanism **before** the provider sees it.

There is **no `--iv`** and **no `--key`** parameter anywhere in the contract.

### Value format

The ciphertext value is **plain base64** with no human-visible prefix. Any nonce/IV and authentication data live **inside** the blob. The provider is chosen by configuration, not by inspecting the value — a wrong provider or wrong key is caught by the provider's own integrity check and surfaces as a clean error.

## Building an Encrypter Provider

A provider is a standard aux4 package that:

1. Declares `aux4/encrypter` as a dependency.
2. Re-opens the `encrypter:provider` profile and adds a `<provider>` command routing to `encrypter:provider:<provider>`.
3. Implements `encrypt` and `decrypt` (and optionally `generate-secret`) in the `encrypter:provider:<provider>` profile, following the contract above.

Because each provider puts its real commands in its **own** `encrypter:provider:<provider>` sub-profile and only appends a single `<provider>` command to the shared `encrypter:provider` profile, providers never collide.

See `aux4/encrypter-aux4` for a complete provider implementation.
