### Using parameters

```execute
aux4 encrypter encrypt --iv nhBGMMMQDJeMncEs --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile --text "Hello World"
```

```expect
9V3GJr82MrlrEpZKyUF5OA==
```

### Using Environment Variables

```execute
export AUX4_IV=nhBGMMMQDJeMncEs
export AUX4_SECRET=9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
aux4 encrypter encrypt --text "Hello World"
```

```expect
9V3GJr82MrlrEpZKyUF5OA==
```

### Using aux4/config

```file:config.yaml
config:
  test:
    iv: nhBGMMMQDJeMncEs
    secret: 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
```

```execute
aux4 encrypter encrypt --config test --text "Hello World"
```

```expect
9V3GJr82MrlrEpZKyUF5OA==
```

### Using File

```execute
echo 'Hello World' > hello.txt
cat hello.txt | aux4 encrypter encrypt --iv nhBGMMMQDJeMncEs --secret 9j48xBeBCQeV7R6Gnhe1V3acoBSS7ile
rm -f hello.txt
```

```expect
9V3GJr82MrlrEpZKyUF5OA==
```
