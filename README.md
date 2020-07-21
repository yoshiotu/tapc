`tapc.Conn` wraps `net.Conn` and can be used to
record or log every byte read or written to the
underlying connection.

```
type myLogger struct{}

func (m myLogger) Log(s string) {
	fmt.Println(s)
}

mlog := myLogger{}

tappedConn := tapc.NewConn(conn, mlog)
```

If you use `tappedConn` to send and receive data,
any data read or written will be logged.
