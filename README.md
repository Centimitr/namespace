# namespace

[![Build Status](https://travis-ci.org/Centimitr/namespace.svg?branch=master)](https://travis-ci.org/Centimitr/namespace)

A tiny library for prefix generation and using in map string keys.

### A simple example

Now we want to store an item named like "Service.Chat.Message.20160701" into a map. Actually, the prefix "Service.Chat.Message" is always fixed during the related transactions. So in this condition, what we need are:

1. Have a namespace
2. Manage potential key prefixes
3. Generate a non-conflict name

```Go
package main

import (
	"fmt"
	"github.com/Centimitr/namespace"
)

func main() {
  
    // generate a namespace
	n := namespace.New()
	
	// get a prefix
	// prefixes are consist of groups of description
	_, p := n.Prefix("Service")
	
	// extend the prefix
	_, p = p.Extend("Chat", "Message")
	
	// generate a scope from the prefix with Apply()
	// one scope must not be a part of another scope
	_, s := p.Apply()
	
	// use Key() method to generate a key with the given string
	k := s.Key("20160701")
	fmt.Println(k) 
}

```

```
  Output:
  Service.Chat.Message.20160701
```

### Notice

The package is firstly build for my other projects, and now, the API are not stable and may be changed at any timepoint. Prefix string are not checked when added into a Prefix struct. Features like conflict detectition and containers binding is under consideration.
