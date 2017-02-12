# Tailer
Notifies if something happened in log file

###How to use:

```go build```

```./tailer -subject 'error' -recepient 'recepient@example.com' -pattern 'HTTP\/1.0" (4\d{2}|5\d{2}|6\d{2})' -logfile 'test.log'```
