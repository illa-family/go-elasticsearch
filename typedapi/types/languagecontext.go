// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// LanguageContext type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_global/get_script_languages/types.ts#L22-L25
type LanguageContext struct {
	Contexts []string                      `json:"contexts"`
	Language scriptlanguage.ScriptLanguage `json:"language"`
}

// LanguageContextBuilder holds LanguageContext struct and provides a builder API.
type LanguageContextBuilder struct {
	v *LanguageContext
}

// NewLanguageContext provides a builder for the LanguageContext struct.
func NewLanguageContextBuilder() *LanguageContextBuilder {
	r := LanguageContextBuilder{
		&LanguageContext{},
	}

	return &r
}

// Build finalize the chain and returns the LanguageContext struct
func (rb *LanguageContextBuilder) Build() LanguageContext {
	return *rb.v
}

func (rb *LanguageContextBuilder) Contexts(contexts ...string) *LanguageContextBuilder {
	rb.v.Contexts = contexts
	return rb
}

func (rb *LanguageContextBuilder) Language(language scriptlanguage.ScriptLanguage) *LanguageContextBuilder {
	rb.v.Language = language
	return rb
}