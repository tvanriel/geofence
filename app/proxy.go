package app

import (
	"fmt"
	"io"
	"net"
)

type Proxy struct {
	Middleware []Middleware
	Listen     string
	Upstream   string
	Name       string

	Reporter *DiscordReporter
}

func NewProxy(middleware []Middleware, listen string, upstream string, name string, reporter *DiscordReporter) *Proxy {

	return &Proxy{
		Listen:     listen,
		Upstream:   upstream,
		Reporter:   reporter,
		Middleware: middleware,
		Name:       name,
	}
}

func (p *Proxy) ListenAndServe() error {
	listener, err := net.Listen("tcp", p.Listen)
	fmt.Printf("Listening for Service %s on %v\n", p.Name, p.Listen)
	if err != nil {
		return err
	}
	for {
		cl, err := listener.Accept()
		fmt.Printf("server: accept: %v\n", cl.RemoteAddr().String())

		if err != nil {
			break
		}

		go func() {
			for i := range p.Middleware {
				if !p.Middleware[i].Passes(cl.RemoteAddr().String()) {
					cl.Close()
					// go p.Reporter.ReportBlocked(
					// 	p.Name,
					// 	cl.RemoteAddr().String(),
					// 	p.Middleware[i].Rule,
					// )
					fmt.Printf("Blocked connection from %s: %s\n",
						cl.RemoteAddr().String(),
						p.Middleware[i].Rule,
					)
					return
				}
			}
			up, err := net.Dial("tcp", p.Upstream)
			if err != nil {
				cl.Close()

				// go p.Reporter.CannotDial(
				// 	p.Name,
				// 	cl.RemoteAddr().String(),
				// 	err,
				// )
				fmt.Printf("[%s] Cannot dial upstream for %s\n",
					p.Name,
					cl.RemoteAddr().String(),
				)
				return
			}
			fmt.Printf("[*] Accepted from: %s\n", cl.RemoteAddr())

			// go p.Reporter.ReportAccepted(
			// 	p.Name,
			// 	cl.RemoteAddr().String(),
			// )
			handleConnection(cl, up)
		}()
	}
	return nil
}

func handleConnection(client, up net.Conn) {
	go io.Copy(up, client)
	io.Copy(client, up)
}
