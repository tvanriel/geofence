package app

import "fmt"

func NewProxyBooter(proxies []Proxy) *ProxyBooter {
	return &ProxyBooter{
		Proxies: proxies,
	}
}

type ProxyBooter struct {
	Proxies []Proxy
}

func (p *ProxyBooter) BootAll() {
	for i := range p.Proxies {
		fmt.Printf("Booting Proxy %s\n", p.Proxies[i].Name)
		go p.Proxies[i].ListenAndServe()
	}
}
