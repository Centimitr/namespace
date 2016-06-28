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

type Scope struct {
	// scope's name
	name string
	// the namespace it belongs to
	namespace           *Namespace
	ruleOfGet           func(scope, name string) string
	UseDefaultRuleOfGet bool
}

func (s *Scope) SetGetRule(fn func(string, string) string) {
	s.UseDefaultRuleOfGet = false
	s.ruleOfGet = fn
}

// generate key string with namespace\scope and string
func (s *Scope) Get(name string) string {
	if s.UseDefaultRuleOfGet {
		return s.namespace.ruleOfGet(s.name, name)
	} else {
		return s.ruleOfGet(s.name, name)
	}
}

func (s *Scope) Handler(name string) Handler {
	if s.namespace.binding == nil {
		panic("Namespace hasn't have a binding yet.")
	}
	return Handler{
		key:       name,
		Interface: s.namespace.binding,
	}
}
