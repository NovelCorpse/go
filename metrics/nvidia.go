package metrics

import (
	"time"

	libNvidia "github.com/BradburyLab/go/metrics/nvidia"
)

type NvidiaResult struct {
	v libNvidia.Devices
	e error
}

func (it *NvidiaResult) set(v libNvidia.Devices) *NvidiaResult { it.v = v; return it }
func (it *NvidiaResult) setErr(v error) *NvidiaResult          { it.e = v; return it }

func (it *NvidiaResult) V() interface{} { return it.v }
func (it *NvidiaResult) Name() string   { return "" }
func (it *NvidiaResult) Kind() Kind     { return KIND_NVIDIA }
func (it *NvidiaResult) OK() bool       { return true }
func (it *NvidiaResult) Err() error     { return it.e }

func NewNvidiaResult() *NvidiaResult { return new(NvidiaResult) }

type nvidia struct {
	scheme string
	host   string
	port   int
	path   string

	timeout                   time.Duration
	maxIdleConnectionsPerHost int
	dialTimeout               time.Duration
	dialKeepAlive             time.Duration
	tlsHandshakeTimeout       time.Duration
	tlsInsecureSkipVerify     bool

	c *libNvidia.Nvidia
}

func (it *nvidia) SetScheme(v string) *nvidia         { it.scheme = v; return it }
func (it *nvidia) SetHost(v string) *nvidia           { it.host = v; return it }
func (it *nvidia) SetPort(v int) *nvidia              { it.port = v; return it }
func (it *nvidia) SetPath(v string) *nvidia           { it.path = v; return it }
func (it *nvidia) SetTimeout(v time.Duration) *nvidia { it.timeout = v; return it }
func (it *nvidia) SetMaxIdleConnectionsPerHost(v int) *nvidia {
	it.maxIdleConnectionsPerHost = v
	return it
}
func (it *nvidia) SetDialTimeout(v time.Duration) *nvidia   { it.dialTimeout = v; return it }
func (it *nvidia) SetDialKeepAlive(v time.Duration) *nvidia { it.dialKeepAlive = v; return it }
func (it *nvidia) SetTLSHandshakeTimeout(v time.Duration) *nvidia {
	it.tlsHandshakeTimeout = v
	return it
}
func (it *nvidia) SetTLSInsecureSkipVerify(v bool) *nvidia { it.tlsInsecureSkipVerify = v; return it }

func (it *nvidia) Len() int { return 1 }

func (it *nvidia) client() *libNvidia.Nvidia {
	if it.c == nil {
		it.c = libNvidia.
			NewNvidia(it.host, it.port).
			SetScheme(it.scheme).
			SetPath(it.path).
			SetTimeout(it.timeout).
			SetMaxIdleConnectionsPerHost(it.maxIdleConnectionsPerHost).
			SetDialTimeout(it.dialTimeout).
			SetDialKeepAlive(it.dialKeepAlive).
			SetTLSHandshakeTimeout(it.tlsHandshakeTimeout).
			SetTLSInsecureSkipVerify(it.tlsInsecureSkipVerify)
	}

	return it.c
}

func (it *nvidia) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	result := NewNvidiaResult()
	devices, e := it.client().Status()
	if e != nil {
		result.setErr(e)
	} else {
		result.set(devices)
	}

	out <- result
	return
}

func Nvidia() *nvidia {
	it := new(nvidia)

	it.scheme = "http"
	it.host = "127.0.0.1"
	it.port = 4459
	it.path = "/"

	it.timeout = 1 * time.Second
	it.maxIdleConnectionsPerHost = 16
	it.dialTimeout = 1 * time.Second
	it.dialKeepAlive = 10 * time.Second
	it.tlsHandshakeTimeout = time.Second
	it.tlsInsecureSkipVerify = true

	return it
}
