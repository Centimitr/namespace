// Copyright 2016 Cyako Author

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
	scopes  map[string]Scope
	getRule func(scope, name string) string
}

type Scope struct {
	// scope's name
	name string
	// the namespace it belongs to
	namespace         *Namespace
	getRule           func(scope, name string) string
	UseDefaultGetRule bool
}

func (n *Namespace) Apply(scopeName string) (ok bool, _ Scope) {
	if _, ok := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = Scope{
			name:              scopeName,
			namespace:         n,
			UseDefaultGetRule: true,
		}
		return true, n.scopes[scopeName]
	}
	return false, Scope{}
}

func (n *Namespace) Use(scopeName string) Scope {
	if _, ok := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = Scope{
			name:              scopeName,
			namespace:         n,
			UseDefaultGetRule: true,
		}
	}
	return n.scopes[scopeName]
}

func (n *Namespace) SetGetRule(fn func(string, string) string) {
	n.getRule = fn
}

func (s *Scope) SetGetRule(fn func(string, string) string) {
	s.UseDefaultGetRule = false
	s.getRule = fn
}

func (s *Scope) Get(name string) string {
	if s.UseDefaultGetRule {
		return s.namespace.getRule(s.name, name)
	} else {
		return s.getRule(s.name, name)
	}
}

func (n *Namespace) DirectGet(scope, name string) string {
	if s, ok := n.scopes[scope]; ok {
		return s.getRule(scope, name)
	} else {
		return n.getRule(scope, name)
	}
}

func (n *Namespace) Init() {
	n.scopes = make(map[string]Scope)
	n.getRule = func(scope, name string) string {
		return scope + "." + name
	}
}

func New() Namespace {
	return Namespace{
		scopes: make(map[string]Scope),
		getRule: func(scope, name string) string {
			return scope + "." + name
		},
	}
}
