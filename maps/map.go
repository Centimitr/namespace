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

type Map struct {
	m         map[string]interface{}
	Namespace MapNamespace
}

func (this *Map) has(key string) bool {
	_, ok := this.m[key]
	return ok
}

func (this *Map) get(key string) (interface{}, bool) {
	value, ok := this.m[key]
	return value, ok
}

func (this *Map) mustGet(key string) interface{} {
	value, _ := this.m[key]
	return value
}

func (this *Map) set(key string, value interface{}) {
	this.m[key] = value
}

func (this *Map) delete(key string) {
	delete(this.m, key)
}

func (this *Map) Init() {
	this.m = make(map[string]interface{})
	this.Namespace.Init()
	this.Namespace.m = this
}
