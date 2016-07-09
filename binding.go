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

// import (
// 	"errors"
// )

// var lackBinding = errors.New("No scope has been binded to this binding.")

// type Binding struct {
// 	scope *Scope
// }

// func (b *Binding) Bind(s *Scope) {
// 	b.scope = s
// }

// func (b *Binding) Key(key string) (string, error) {
// 	if b.scope == nil {
// 		return "", lackBinding
// 	}
// 	return b.scope.Key(key), nil
// }

// func (b *Binding) MustKey(key string) string {
// 	if b.scope == nil {
// 		return key
// 	}
// 	return b.scope.Key(key)
// }
