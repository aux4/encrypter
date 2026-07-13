#### Description

The `provider` command is the namespace under which encryption providers register. Each installed provider package (e.g. `aux4/encrypter-aux4`) adds a `<provider>` subcommand here that routes to its own implementation.

You normally do not call this directly — the `encrypt`, `decrypt`, and `generate-secret` facade commands dispatch to `aux4 encrypter provider <provider> ...` for you based on the resolved provider. It is available for explicit, provider-pinned invocations and for debugging.

#### Usage

```bash
aux4 encrypter provider <provider> <command> [options]
```

#### Example

```bash
echo -n 'hello world' | aux4 encrypter provider aux4 encrypt --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
```

```text
AZ1nQ2f8...base64...==
```
