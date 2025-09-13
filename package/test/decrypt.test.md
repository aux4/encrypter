### Using parameters

```execute
aux4 encrypter decrypt --iv nhBGMMMQDJeMncEs --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile --text 9V3GJr82MrlrEpZKyUF5OA==
```

```expect
Hello World
```

### Using Environment Variables

```execute
export AUX4_IV=nhBGMMMQDJeMncEs
export AUX4_SECRET=9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
aux4 encrypter decrypt --text 9V3GJr82MrlrEpZKyUF5OA==
```

```expect
Hello World
```

### Using aux4/config

```file:config.yaml
config:
  test:
    iv: nhBGMMMQDJeMncEs
    secret: 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
```

```execute
aux4 encrypter decrypt --config test --text '9V3GJr82MrlrEpZKyUF5OA=='
```

```expect
Hello World
```

### Using File

```execute
echo '9V3GJr82MrlrEpZKyUF5OA==' > token.txt
cat token.txt | aux4 encrypter decrypt --iv nhBGMMMQDJeMncEs --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
rm -f token.txt
```

```expect
Hello World
```

### Multiple texts

```execute
aux4 encrypter decrypt --iv nhBGMMMQDJeMncEs --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile --text "iOgicUiHYmWpbq41dfiN0Q==" --text "hmEsKxb2XqCAQ2eRSY0mPA==" --text "N7m7RL3+045hv6/0TA2QjA=="
```

```expect
Hello
World
Test
```
