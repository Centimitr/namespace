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

type MapHandler struct {
	key string
	m   *Map
}

func (this *MapHandler) Has() bool {
	return this.m.has(this.key)
}

func (this *MapHandler) Get() (interface{}, bool) {
	return this.m.get(this.key)
}

func (this *MapHandler) MustGet() interface{} {
	return this.m.mustGet(this.key)
}

func (this *MapHandler) Set(value interface{}) {
	this.m.set(this.key, value)
}

func (this *MapHandler) Delete() {
	this.m.delete(this.key)
}
