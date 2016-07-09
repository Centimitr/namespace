// Copyright 2016 Centimitr

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required` by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package namespace

import (
	"strings"
)

type Namespace struct {
	scopes           map[string]Scope
	prefixes         map[string]Prefix
	keyConcatRule    func(scope, name string) string
	prefixConcatRule func(prefixes ...string) string
	// a kv like structs like map that this namespace bind to
	// binding Interface
}

func (n *Namespace) hasNoConflict(scopeName string) bool {
	for _, s := range n.scopes {
		if strings.HasPrefix(s.name, scopeName) || strings.HasPrefix(scopeName, s.name) {
			return false
		}
	}
	return true
}

func (n *Namespace) assignNewScope(scopeName string) {
	n.scopes[scopeName] = Scope{
		name:      scopeName,
		namespace: n,
		// UseDefaultRuleOfGet: true,
	}
}

// try to use a prefix that haven't been used yet, and get the scope
func (n *Namespace) Apply(scopeName string) (ok bool, _ Scope) {
	if _, ok := n.scopes[scopeName]; !ok && n.hasNoConflict(scopeName) {
		n.assignNewScope(scopeName)
		return true, n.scopes[scopeName]
	}
	return false, Scope{}
}

// get the scope with specific prefix, apply if it is not exist
func (n *Namespace) Use(scopeName string) (ok bool, _ Scope) {
	if _, ok := n.scopes[scopeName]; !ok {
		if !n.hasNoConflict(scopeName) {
			return false, Scope{}
		}
		n.assignNewScope(scopeName)
	}
	return true, n.scopes[scopeName]
}

// how to concat prefix and string
func (n *Namespace) SetKeyConcatRule(fn func(string, string) string) {
	n.keyConcatRule = fn
}

func (n *Namespace) SetPrefixConcatRule(fn func(...string) string) {
	n.prefixConcatRule = fn
}

// init maps and keyConcatRule functions
func (n *Namespace) Init() {
	n.scopes = make(map[string]Scope)
	n.prefixes = make(map[string]Prefix)
	n.keyConcatRule = func(scope, name string) string {
		return scope + ":" + name
	}
	n.prefixConcatRule = func(names ...string) string {
		return strings.Join(names, ".")
	}
}

func (n *Namespace) Prefix(names ...string) (ok bool, _ Prefix) {
	key := getPrefixNamesKey(names)
	if _, ok := n.prefixes[key]; !ok {
		n.prefixes[key] = Prefix{
			names:     names,
			namespace: n,
		}
		return true, n.prefixes[key]
	}
	return false, Prefix{}
}

func New() Namespace {
	ns := Namespace{}
	ns.Init()
	return ns
}
