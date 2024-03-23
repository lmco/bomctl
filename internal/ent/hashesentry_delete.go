// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/ent/hashesentry.go
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
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bomctl/bomctl/internal/ent/hashesentry"
	"github.com/bomctl/bomctl/internal/ent/predicate"
)

// HashesEntryDelete is the builder for deleting a HashesEntry entity.
type HashesEntryDelete struct {
	config
	hooks    []Hook
	mutation *HashesEntryMutation
}

// Where appends a list predicates to the HashesEntryDelete builder.
func (hed *HashesEntryDelete) Where(ps ...predicate.HashesEntry) *HashesEntryDelete {
	hed.mutation.Where(ps...)
	return hed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hed *HashesEntryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hed.sqlExec, hed.mutation, hed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hed *HashesEntryDelete) ExecX(ctx context.Context) int {
	n, err := hed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hed *HashesEntryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hashesentry.Table, sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeInt))
	if ps := hed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hed.mutation.done = true
	return affected, err
}

// HashesEntryDeleteOne is the builder for deleting a single HashesEntry entity.
type HashesEntryDeleteOne struct {
	hed *HashesEntryDelete
}

// Where appends a list predicates to the HashesEntryDelete builder.
func (hedo *HashesEntryDeleteOne) Where(ps ...predicate.HashesEntry) *HashesEntryDeleteOne {
	hedo.hed.mutation.Where(ps...)
	return hedo
}

// Exec executes the deletion query.
func (hedo *HashesEntryDeleteOne) Exec(ctx context.Context) error {
	n, err := hedo.hed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hashesentry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hedo *HashesEntryDeleteOne) ExecX(ctx context.Context) {
	if err := hedo.Exec(ctx); err != nil {
		panic(err)
	}
}
