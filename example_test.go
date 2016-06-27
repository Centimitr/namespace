package namespace_test

import (
	"fmt"
	ns "github.com/Centimitr/namespace"
)

func Example() {
	n := ns.New()
	_, p := n.NewPrefix("SERVICE", "NS")
	_, p = prefixOfService.Extend("REALTIME", "GET")
	_, s := prefixOfRealtime.Apply("get")
	key := s.Get("123")
	fmt.Println(key)
}
