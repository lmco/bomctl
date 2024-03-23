// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/ent/document.go
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
	"github.com/bomctl/bomctl/internal/ent/document"
	"github.com/bomctl/bomctl/internal/ent/metadata"
	"github.com/bomctl/bomctl/internal/ent/nodelist"
)

// Document is the model entity for the Document schema.
type Document struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DocumentQuery when eager-loading is set.
	Edges        DocumentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DocumentEdges holds the relations/edges for other nodes in the graph.
type DocumentEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// NodeList holds the value of the node_list edge.
	NodeList *NodeList `json:"node_list,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocumentEdges) MetadataOrErr() (*Metadata, error) {
	if e.Metadata != nil {
		return e.Metadata, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: metadata.Label}
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// NodeListOrErr returns the NodeList value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocumentEdges) NodeListOrErr() (*NodeList, error) {
	if e.NodeList != nil {
		return e.NodeList, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: nodelist.Label}
	}
	return nil, &NotLoadedError{edge: "node_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Document) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Document fields.
func (d *Document) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Document.
// This includes values selected through modifiers, order, etc.
func (d *Document) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryMetadata queries the "metadata" edge of the Document entity.
func (d *Document) QueryMetadata() *MetadataQuery {
	return NewDocumentClient(d.config).QueryMetadata(d)
}

// QueryNodeList queries the "node_list" edge of the Document entity.
func (d *Document) QueryNodeList() *NodeListQuery {
	return NewDocumentClient(d.config).QueryNodeList(d)
}

// Update returns a builder for updating this Document.
// Note that you need to call Document.Unwrap() before calling this method if this Document
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Document) Update() *DocumentUpdateOne {
	return NewDocumentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Document entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Document) Unwrap() *Document {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Document is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Document) String() string {
	var builder strings.Builder
	builder.WriteString("Document(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Documents is a parsable slice of Document.
type Documents []*Document
