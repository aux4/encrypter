# aux4/encrypter 3.0.1

## Fixed

- `encrypt`, `decrypt` and `generate-secret` no longer interpolate the raw
  `--provider` value into shell commands. The provider name is now passed with
  `value(provider)`, so a provider name containing shell metacharacters is
  treated as text and can never run a command. The "provider not installed"
  message is unchanged.
