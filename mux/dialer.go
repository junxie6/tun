package mux

import (
	"net"

	"github.com/4396/tun/conn"
	"github.com/xtaci/smux"
)

func Dial(addr string) (d *Dialer, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	sess, err := smux.Client(conn, nil)
	if err != nil {
		return
	}

	d = &Dialer{sess: sess}
	return
}

type Dialer struct {
	sess *smux.Session
}

func (d *Dialer) Dial() (c net.Conn, err error) {
	st, err := d.sess.OpenStream()
	if err != nil {
		return
	}

	c = conn.WithSnappy(st)
	return
}

func (d *Dialer) Close() error {
	return d.sess.Close()
}
