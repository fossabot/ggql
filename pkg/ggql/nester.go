// Copyright 2019-2020 University Health Network
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ggql

// Nester is used to create new context for fields. It allows a tree of user
// data to be built up for fields in an executable.
type Nester interface {

	// Nest should create a context. Usually it will be a new Nester but that
	// is not a requirement. The field provided is the current field the
	// context is being created for.
	Nest(field *Field) interface{}
}
