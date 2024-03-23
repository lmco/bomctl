// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/ent/person.go
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

package person

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsOrg holds the string denoting the is_org field in the database.
	FieldIsOrg = "is_org"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// EdgePerson holds the string denoting the person edge name in mutations.
	EdgePerson = "person"
	// EdgeContacts holds the string denoting the contacts edge name in mutations.
	EdgeContacts = "contacts"
	// Table holds the table name of the person in the database.
	Table = "persons"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "persons"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "metadata_authors"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "persons"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_originators"
	// PersonTable is the table that holds the person relation/edge.
	PersonTable = "persons"
	// PersonColumn is the table column denoting the person relation/edge.
	PersonColumn = "person_contacts"
	// ContactsTable is the table that holds the contacts relation/edge.
	ContactsTable = "persons"
	// ContactsColumn is the table column denoting the contacts relation/edge.
	ContactsColumn = "person_contacts"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIsOrg,
	FieldEmail,
	FieldURL,
	FieldPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "persons"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"metadata_authors",
	"node_suppliers",
	"node_originators",
	"person_contacts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Person queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIsOrg orders the results by the is_org field.
func ByIsOrg(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOrg, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByMetadataField orders the results by metadata field.
func ByMetadataField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetadataStep(), sql.OrderByField(field, opts...))
	}
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}

// ByPersonField orders the results by person field.
func ByPersonField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPersonStep(), sql.OrderByField(field, opts...))
	}
}

// ByContactsField orders the results by contacts field.
func ByContactsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContactsStep(), sql.OrderByField(field, opts...))
	}
}
func newMetadataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetadataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MetadataTable, MetadataColumn),
	)
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
	)
}
func newPersonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, PersonTable, PersonColumn),
	)
}
func newContactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ContactsTable, ContactsColumn),
	)
}
