package router

import "net/http"

type PrefixMux struct {
	prefix string
	Mux    *http.ServeMux
}

func NewPrefixMux(prefix string, mux *http.ServeMux) *PrefixMux {
	return &PrefixMux{prefix, mux}
}

func (p *PrefixMux) Handle(pattern string, handler http.Handler) {
	p.Mux.Handle(p.prefix+pattern, handler)
}

func (p *PrefixMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	p.Mux.HandleFunc(p.prefix+pattern, handler)
}
