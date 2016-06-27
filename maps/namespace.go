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

package maps

import (
	. "github.com/Centimitr/namespace"
)

type MapNamespace struct {
	m *Map
	Namespace
}

func (this *MapNamespace) Apply(scopeName string) (ok bool, _ MapScope) {
	ok, scope := this.Namespace.Apply(scopeName)
	return ok, MapScope{m: this.m, Scope: scope}
}

// get the scope with specific prefix, apply if it is not exist
func (this *MapNamespace) Use(scopeName string) (ok bool, _ MapScope) {
	ok, scope := this.Namespace.Use(scopeName)
	return ok, MapScope{m: this.m, Scope: scope}
}

func (this *MapNamespace) NewPrefix(names ...string) (ok bool, _ MapScopePrefix) {
	ok, prefix := this.Namespace.NewPrefix(names...)
	return ok, MapScopePrefix{m: this.m, namespace: this, ScopePrefix: prefix}
}
