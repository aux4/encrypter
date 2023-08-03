# @aux4/encrypt
Aux4 encrypt and decrypt tool

## Generate key

```bash
$ aux4-encrypt generate-secret

e85fd3a1f57b6d445e4bf077a9239ba74f0c6a6f07dde6a14ca0043240a2b4eeaf7159ae3d3872825940c2d3e02f081f
```

## Encrypt

### Using parameter

```bash
$ aux4-encrypt encrypt --text "Hello World"

ea776229343c22471b5a1f8876ab605a96e7c81a9429014b6ebb777323426117bbdcf82a5f4ad5fc3077dece9a24b80ec698e640b0fbb90aede4d4e35eed8c9cf90f14ccb71441abb9c3bf7a13cca98bb846aee199efec26f33e1ef4f1c5fbbbc64d4b5746066174bf3196
```

### From file

```bash
$ cat hello.txt | aux4-encrypt encrypt

ea776229343c22471b5a1f8876ab605a96e7c81a9429014b6ebb777323426117bbdcf82a5f4ad5fc3077dece9a24b80ec698e640b0fbb90aede4d4e35eed8c9cf90f14ccb71441abb9c3bf7a13cca98bb846aee199efec26f33e1ef4f1c5fbbbc64d4b5746066174bf3196
```

### Using custom secret

#### From parameter

```bash
$ aux4-encrypt encrypt --secret "e85fd3a1f57b6d445e4bf077a9239ba74f0c6a6f07dde6a14ca0043240a2b4eeaf7159ae3d3872825940c2d3e02f081f" --text "Hello World"
```

#### From environment variable

```bash
export AUX4_SECRET_KEY=e85fd3a1f57b6d445e4bf077a9239ba74f0c6a6f07dde6a14ca0043240a2b4eeaf7159ae3d3872825940c2d3e02f081f

$ aux4-encrypt encrypt --text "Hello World"
```

## Decrypt

### Using parameter

```bash
$ aux4-encrypt decrypt --token "ea776229343c22471b5a1f8876ab605a96e7c81a9429014b6ebb777323426117bbdcf82a5f4ad5fc3077dece9a24b80ec698e640b0fbb90aede4d4e35eed8c9cf90f14ccb71441abb9c3bf7a13cca98bb846aee199efec26f33e1ef4f1c5fbbbc64d4b5746066174bf3196"

Hello World
```

### From file

```bash
$ cat token.txt | aux4-encrypt decrypt

Hello World
```

### Using custom secret

#### From parameter

```bash
$ aux4-encrypt decrypt --secret "e85fd3a1f57b6d445e4bf077a9239ba74f0c6a6f07dde6a14ca0043240a2b4eeaf7159ae3d3872825940c2d3e02f081f" --token "5bf94d4b2b0ed4bb9c3ed57b68aa820e501be48a101482eb30076c9ea93d617b6f0de423b65401c6acbf5b0ae3b5680b343a4e82e2ca8657f83999c05e7d9a4b316cf66e5636e3dbf18ddf42ee0d23bca8b9df28413c6e824b66ca5c6eeddc5081ceac72edf92f6741e82f0c99857cb98a8f5384623d123da9a55096c86962"
```

#### From environment variable

```bash
export AUX4_SECRET_KEY=e85fd3a1f57b6d445e4bf077a9239ba74f0c6a6f07dde6a14ca0043240a2b4eeaf7159ae3d3872825940c2d3e02f081f

$ aux4-encrypt decrypt --token "5bf94d4b2b0ed4bb9c3ed57b68aa820e501be48a101482eb30076c9ea93d617b6f0de423b65401c6acbf5b0ae3b5680b343a4e82e2ca8657f83999c05e7d9a4b316cf66e5636e3dbf18ddf42ee0d23bca8b9df28413c6e824b66ca5c6eeddc5081ceac72edf92f6741e82f0c99857cb98a8f5384623d123da9a55096c86962"
```