# encrypter

## help

### should display the encrypter facade help

```execute
aux4 encrypter --help
```

```expect:partial
Encryption facade
```

### should list the encrypt command

```execute
aux4 encrypter --help
```

```expect:partial
Encrypt a text using the selected provider
```

### should list the decrypt command

```execute
aux4 encrypter --help
```

```expect:partial
Decrypt a text using the selected provider
```

### should list the generate-secret command

```execute
aux4 encrypter --help
```

```expect:partial
Generate a random secret using the selected provider
```

### should list the provider namespace

```execute
aux4 encrypter --help
```

```expect:partial
Encryption providers
```

## provider namespace

### should display the provider namespace help

```execute
aux4 encrypter provider --help
```

```expect:partial
Encryption providers
```

## secret handling

The `secret` variable is optional and non-prompting on `encrypt`/`decrypt`. The facade always forwards
`--secret` (even empty) to the provider, so the provider decides whether a key is required and reports a
clean error — the facade never prompts. This lets keyless providers (e.g. KMS) decrypt without a key.

**Note:** When `--secret` and `AUX4_SECRET` are both omitted entirely, the correct non-prompting behavior
requires the ENC-006 core fix (unreleased). On the current core an omitted `secret` triggers a
stack-overflow during variable resolution. Passing `--secret` explicitly (even `''`) — as these tests do —
exercises the intended post-fix behavior and avoids that core bug.

### should round-trip plaintext through the aux4 provider

```execute
printf 'hello world' | aux4 encrypter encrypt --secret testkey | aux4 encrypter decrypt --secret testkey
```

```expect
hello world
```

### should not prompt when secret is empty (clean provider error)

```execute
CT=$(printf 'hi' | aux4 encrypter encrypt --secret realkey); aux4 encrypter decrypt --secret '' "$CT" </dev/null
```

```error:partial
wrong key or corrupt data
```

## missing provider hint

### should hint how to install a missing provider on encrypt

```execute
aux4 encrypter encrypt --provider bogus --secret k 'x' </dev/null
```

```error:partial
Encryption provider 'bogus' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-bogus
```

### should hint how to install a missing provider on decrypt

```execute
aux4 encrypter decrypt --provider bogus --secret k 'x' </dev/null
```

```error:partial
Encryption provider 'bogus' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-bogus
```

### should hint how to install a missing provider on generate-secret

```execute
aux4 encrypter generate-secret --provider bogus </dev/null
```

```error:partial
Encryption provider 'bogus' not installed. Install it with: aux4 aux4 pkger install aux4/encrypter-bogus
```
