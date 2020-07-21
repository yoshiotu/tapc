package tapc

import (
	"encoding/hex"
	"fmt"
	"net"
)

type Conn struct {
	net.Conn
	logger Logger
}

func NewConn(conn net.Conn, logger Logger) *Conn {
	return &Conn{
		Conn:   conn,
		logger: logger,
	}
}

func (c *Conn) Read(b []byte) (n int, err error) {
	n, err = c.Conn.Read(b)

	if n > 0 {
		s := fmt.Sprintf("bytes read: %s\n", hex.Dump(b[:n]))
		c.logger.Log(s)
	}

	return
}

func (c *Conn) Write(b []byte) (n int, err error) {
	n, err = c.Conn.Write(b)

	if n > 0 {
		s := fmt.Sprintf("bytes written: %s\n", hex.Dump(b[:n]))
		c.logger.Log(s)
	}

	return
}
