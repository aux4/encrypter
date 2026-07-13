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

Note: real encrypt/decrypt round-trips are covered in the provider packages' own test suites
(e.g. aux4/encrypter-aux4), where installing the provider pulls in this facade. The facade suite
only exercises provider-independent behavior, since CI installs the facade alone.

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
