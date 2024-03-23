// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/ent/identifiersentry.go
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
	"github.com/bomctl/bomctl/internal/ent/identifiersentry"
)

// IdentifiersEntry is the model entity for the IdentifiersEntry schema.
type IdentifiersEntry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SoftwareIdentifierType holds the value of the "software_identifier_type" field.
	SoftwareIdentifierType identifiersentry.SoftwareIdentifierType `json:"software_identifier_type,omitempty"`
	// SoftwareIdentifierValue holds the value of the "software_identifier_value" field.
	SoftwareIdentifierValue string `json:"software_identifier_value,omitempty"`
	node_identifiers        *string
	selectValues            sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IdentifiersEntry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case identifiersentry.FieldID:
			values[i] = new(sql.NullInt64)
		case identifiersentry.FieldSoftwareIdentifierType, identifiersentry.FieldSoftwareIdentifierValue:
			values[i] = new(sql.NullString)
		case identifiersentry.ForeignKeys[0]: // node_identifiers
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IdentifiersEntry fields.
func (ie *IdentifiersEntry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case identifiersentry.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ie.ID = int(value.Int64)
		case identifiersentry.FieldSoftwareIdentifierType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field software_identifier_type", values[i])
			} else if value.Valid {
				ie.SoftwareIdentifierType = identifiersentry.SoftwareIdentifierType(value.String)
			}
		case identifiersentry.FieldSoftwareIdentifierValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field software_identifier_value", values[i])
			} else if value.Valid {
				ie.SoftwareIdentifierValue = value.String
			}
		case identifiersentry.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_identifiers", values[i])
			} else if value.Valid {
				ie.node_identifiers = new(string)
				*ie.node_identifiers = value.String
			}
		default:
			ie.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IdentifiersEntry.
// This includes values selected through modifiers, order, etc.
func (ie *IdentifiersEntry) Value(name string) (ent.Value, error) {
	return ie.selectValues.Get(name)
}

// Update returns a builder for updating this IdentifiersEntry.
// Note that you need to call IdentifiersEntry.Unwrap() before calling this method if this IdentifiersEntry
// was returned from a transaction, and the transaction was committed or rolled back.
func (ie *IdentifiersEntry) Update() *IdentifiersEntryUpdateOne {
	return NewIdentifiersEntryClient(ie.config).UpdateOne(ie)
}

// Unwrap unwraps the IdentifiersEntry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ie *IdentifiersEntry) Unwrap() *IdentifiersEntry {
	_tx, ok := ie.config.driver.(*txDriver)
	if !ok {
		panic("ent: IdentifiersEntry is not a transactional entity")
	}
	ie.config.driver = _tx.drv
	return ie
}

// String implements the fmt.Stringer.
func (ie *IdentifiersEntry) String() string {
	var builder strings.Builder
	builder.WriteString("IdentifiersEntry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ie.ID))
	builder.WriteString("software_identifier_type=")
	builder.WriteString(fmt.Sprintf("%v", ie.SoftwareIdentifierType))
	builder.WriteString(", ")
	builder.WriteString("software_identifier_value=")
	builder.WriteString(ie.SoftwareIdentifierValue)
	builder.WriteByte(')')
	return builder.String()
}

// IdentifiersEntries is a parsable slice of IdentifiersEntry.
type IdentifiersEntries []*IdentifiersEntry
