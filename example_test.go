package namespace_test

import (
	"fmt"
	ns "github.com/Centimitr/namespace"
)

func ExampleApply() {
	n := ns.New()
	_, p := n.Prefix("SERVICE", "KVSTORE")
	_, p = p.Extend("REALTIME")
	_, p = p.Extend("LISTENER_LIST")
	_, s := p.Apply("LISTENERS")
	key := s.Key("User")
	fmt.Println(key)
}

func ExampleHandler() {
	n := ns.New()
	_, p := n.Prefix("SERVICE")
	_, p = p.Extend("KVSTORE", "USERS")
	_, s := p.Apply("Centimitr")
	m := ns.Map{}
	m.Init()
	n.Bind(&m)
	h := s.Handler("hobby")
	h.Set("girl")
	v := h.Get()
	fmt.Println(s.Key("hobby"), v)
}
