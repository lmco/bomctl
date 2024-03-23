// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/ent/externalreference.go
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------
package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bomctl/bomctl/internal/ent/externalreference"
)

// ExternalReference is the model entity for the ExternalReference schema.
type ExternalReference struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Authority holds the value of the "authority" field.
	Authority string `json:"authority,omitempty"`
	// Type holds the value of the "type" field.
	Type externalreference.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExternalReferenceQuery when eager-loading is set.
	Edges        ExternalReferenceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ExternalReferenceEdges holds the relations/edges for other nodes in the graph.
type ExternalReferenceEdges struct {
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"nodes,omitempty"`
	// Hashes holds the value of the hashes edge.
	Hashes []*HashesEntry `json:"hashes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e ExternalReferenceEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[0] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// HashesOrErr returns the Hashes value or an error if the edge
// was not loaded in eager-loading.
func (e ExternalReferenceEdges) HashesOrErr() ([]*HashesEntry, error) {
	if e.loadedTypes[1] {
		return e.Hashes, nil
	}
	return nil, &NotLoadedError{edge: "hashes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExternalReference) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case externalreference.FieldID:
			values[i] = new(sql.NullInt64)
		case externalreference.FieldURL, externalreference.FieldComment, externalreference.FieldAuthority, externalreference.FieldType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExternalReference fields.
func (er *ExternalReference) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case externalreference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			er.ID = int(value.Int64)
		case externalreference.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				er.URL = value.String
			}
		case externalreference.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				er.Comment = value.String
			}
		case externalreference.FieldAuthority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authority", values[i])
			} else if value.Valid {
				er.Authority = value.String
			}
		case externalreference.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				er.Type = externalreference.Type(value.String)
			}
		default:
			er.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExternalReference.
// This includes values selected through modifiers, order, etc.
func (er *ExternalReference) Value(name string) (ent.Value, error) {
	return er.selectValues.Get(name)
}

// QueryNodes queries the "nodes" edge of the ExternalReference entity.
func (er *ExternalReference) QueryNodes() *NodeQuery {
	return NewExternalReferenceClient(er.config).QueryNodes(er)
}

// QueryHashes queries the "hashes" edge of the ExternalReference entity.
func (er *ExternalReference) QueryHashes() *HashesEntryQuery {
	return NewExternalReferenceClient(er.config).QueryHashes(er)
}

// Update returns a builder for updating this ExternalReference.
// Note that you need to call ExternalReference.Unwrap() before calling this method if this ExternalReference
// was returned from a transaction, and the transaction was committed or rolled back.
func (er *ExternalReference) Update() *ExternalReferenceUpdateOne {
	return NewExternalReferenceClient(er.config).UpdateOne(er)
}

// Unwrap unwraps the ExternalReference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (er *ExternalReference) Unwrap() *ExternalReference {
	_tx, ok := er.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExternalReference is not a transactional entity")
	}
	er.config.driver = _tx.drv
	return er
}

// String implements the fmt.Stringer.
func (er *ExternalReference) String() string {
	var builder strings.Builder
	builder.WriteString("ExternalReference(")
	builder.WriteString(fmt.Sprintf("id=%v, ", er.ID))
	builder.WriteString("url=")
	builder.WriteString(er.URL)
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(er.Comment)
	builder.WriteString(", ")
	builder.WriteString("authority=")
	builder.WriteString(er.Authority)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", er.Type))
	builder.WriteByte(')')
	return builder.String()
}

// ExternalReferences is a parsable slice of ExternalReference.
type ExternalReferences []*ExternalReference
