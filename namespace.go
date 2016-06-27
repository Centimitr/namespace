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

type Namespace struct {
	scopes             map[string]Scope
	scopePrefixes      map[string]ScopePrefix
	ruleOfGet          func(scope, name string) string
	ruleOfPrefixConcat func(prefix, extprefix string) string
}

// try to use a prefix that haven't been used yet, and get the scope
func (n *Namespace) Apply(scopeName string) (ok bool, _ Scope) {
	if _, ok := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = Scope{
			name:                scopeName,
			namespace:           n,
			UseDefaultRuleOfGet: true,
		}
		return true, n.scopes[scopeName]
	}
	return false, Scope{}
}

// get the scope with specific prefix, apply if it is not exist
func (n *Namespace) Use(scopeName string) Scope {
	if _, ok := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = Scope{
			name:                scopeName,
			namespace:           n,
			UseDefaultRuleOfGet: true,
		}
	}
	return n.scopes[scopeName]
}

// how to concat prefix and string
func (n *Namespace) SetGetRule(fn func(string, string) string) {
	n.ruleOfGet = fn
}

func (n *Namespace) DirectGet(scope, name string) string {
	if s, ok := n.scopes[scope]; ok {
		return s.ruleOfGet(scope, name)
	} else {
		return n.ruleOfGet(scope, name)
	}
}

// init maps and ruleOfGet functions
func (n *Namespace) Init() {
	n.scopes = make(map[string]Scope)
	n.scopePrefixes = make(map[string]ScopePrefix)
	n.ruleOfGet = func(scope, name string) string {
		return scope + ":" + name
	}
	n.ruleOfPrefixConcat = func(scope, name string) string {
		return scope + "." + name
	}
}

func (n *Namespace) NewPrefix(name string) (ok bool, _ ScopePrefix) {
	if _, ok := n.scopePrefixes[name]; !ok {
		n.scopePrefixes[name] = ScopePrefix{
			name:      name,
			namespace: n,
		}
		return true, n.scopePrefixes[name]
	}
	return false, ScopePrefix{}
}

func New() Namespace {
	ns := Namespace{}
	ns.Init()
	return ns
}
