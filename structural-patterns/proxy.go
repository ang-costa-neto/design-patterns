package main

import "fmt"

// Interface
type Subject interface {
	Request() string
}

// Real subject
type RealSubject struct{}

func (r RealSubject) Request() string {
	return "Real Request"
}

// Proxy
type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	return p.realSubject.Request()
}

func mainProxy() {
	proxy := &Proxy{}
	fmt.Println(proxy.Request()) // Real Request
}
