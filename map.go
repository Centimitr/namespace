// Copyright 2016 Centimitr

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use m file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required` by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package namespace

type Map struct {
	m map[string]interface{}
}

func (m *Map) Has(key string) bool {
	_, ok := m.m[key]
	return ok
}

func (m *Map) Get(key string) interface{} {
	value, _ := m.m[key]
	return value
}

func (m *Map) Set(key string, value interface{}) {
	m.m[key] = value
}

func (m *Map) Delete(key string) {
	delete(m.m, key)
}

func (m *Map) Init() {
	m.m = make(map[string]interface{})
}
