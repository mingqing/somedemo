# server

```
./tcp-server -l 10.5.17.76:4444 -w -i 1000
```

# client

```
./tcp-client -s 10.5.17.76:4444 -w -i 1000
```

# watch server and client tcp socket

```
watch -n1 ss -t -4
```
