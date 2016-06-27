package maps_test

import (
	"fmt"
	maps "github.com/Centimitr/namespace/maps"
)

func Example() {
	m := maps.Map{}
	m.Init()
	_, p := m.Namespace.NewPrefix("SERVICE", "KVSTORE")
	_, p = p.Extend("REALTIME")
	_, p = p.Extend("LISTENER_LIST")
	_, s := p.Apply("LISTENERS")
	key := s.Get("User")
	fmt.Println(key)
	h := s.Handler("User")
	h.Set("123")
	v := h.MustGet()
	fmt.Println(v)
}
