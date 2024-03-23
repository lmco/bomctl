// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/ent/edge.go
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
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bomctl/bomctl/internal/ent/edge"
)

// EdgeCreate is the builder for creating a Edge entity.
type EdgeCreate struct {
	config
	mutation *EdgeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetType sets the "type" field.
func (ec *EdgeCreate) SetType(e edge.Type) *EdgeCreate {
	ec.mutation.SetType(e)
	return ec
}

// SetFrom sets the "from" field.
func (ec *EdgeCreate) SetFrom(s string) *EdgeCreate {
	ec.mutation.SetFrom(s)
	return ec
}

// SetTo sets the "to" field.
func (ec *EdgeCreate) SetTo(s string) *EdgeCreate {
	ec.mutation.SetTo(s)
	return ec
}

// Mutation returns the EdgeMutation object of the builder.
func (ec *EdgeCreate) Mutation() *EdgeMutation {
	return ec.mutation
}

// Save creates the Edge in the database.
func (ec *EdgeCreate) Save(ctx context.Context) (*Edge, error) {
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EdgeCreate) SaveX(ctx context.Context) *Edge {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EdgeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EdgeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EdgeCreate) check() error {
	if _, ok := ec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Edge.type"`)}
	}
	if v, ok := ec.mutation.GetType(); ok {
		if err := edge.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Edge.type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "Edge.from"`)}
	}
	if _, ok := ec.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "Edge.to"`)}
	}
	return nil
}

func (ec *EdgeCreate) sqlSave(ctx context.Context) (*Edge, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EdgeCreate) createSpec() (*Edge, *sqlgraph.CreateSpec) {
	var (
		_node = &Edge{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(edge.Table, sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ec.conflict
	if value, ok := ec.mutation.GetType(); ok {
		_spec.SetField(edge.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ec.mutation.From(); ok {
		_spec.SetField(edge.FieldFrom, field.TypeString, value)
		_node.From = value
	}
	if value, ok := ec.mutation.To(); ok {
		_spec.SetField(edge.FieldTo, field.TypeString, value)
		_node.To = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Edge.Create().
//		SetType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EdgeUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (ec *EdgeCreate) OnConflict(opts ...sql.ConflictOption) *EdgeUpsertOne {
	ec.conflict = opts
	return &EdgeUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Edge.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EdgeCreate) OnConflictColumns(columns ...string) *EdgeUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EdgeUpsertOne{
		create: ec,
	}
}

type (
	// EdgeUpsertOne is the builder for "upsert"-ing
	//  one Edge node.
	EdgeUpsertOne struct {
		create *EdgeCreate
	}

	// EdgeUpsert is the "OnConflict" setter.
	EdgeUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *EdgeUpsert) SetType(v edge.Type) *EdgeUpsert {
	u.Set(edge.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EdgeUpsert) UpdateType() *EdgeUpsert {
	u.SetExcluded(edge.FieldType)
	return u
}

// SetFrom sets the "from" field.
func (u *EdgeUpsert) SetFrom(v string) *EdgeUpsert {
	u.Set(edge.FieldFrom, v)
	return u
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *EdgeUpsert) UpdateFrom() *EdgeUpsert {
	u.SetExcluded(edge.FieldFrom)
	return u
}

// SetTo sets the "to" field.
func (u *EdgeUpsert) SetTo(v string) *EdgeUpsert {
	u.Set(edge.FieldTo, v)
	return u
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *EdgeUpsert) UpdateTo() *EdgeUpsert {
	u.SetExcluded(edge.FieldTo)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Edge.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EdgeUpsertOne) UpdateNewValues() *EdgeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Edge.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EdgeUpsertOne) Ignore() *EdgeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EdgeUpsertOne) DoNothing() *EdgeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EdgeCreate.OnConflict
// documentation for more info.
func (u *EdgeUpsertOne) Update(set func(*EdgeUpsert)) *EdgeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EdgeUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *EdgeUpsertOne) SetType(v edge.Type) *EdgeUpsertOne {
	return u.Update(func(s *EdgeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EdgeUpsertOne) UpdateType() *EdgeUpsertOne {
	return u.Update(func(s *EdgeUpsert) {
		s.UpdateType()
	})
}

// SetFrom sets the "from" field.
func (u *EdgeUpsertOne) SetFrom(v string) *EdgeUpsertOne {
	return u.Update(func(s *EdgeUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *EdgeUpsertOne) UpdateFrom() *EdgeUpsertOne {
	return u.Update(func(s *EdgeUpsert) {
		s.UpdateFrom()
	})
}

// SetTo sets the "to" field.
func (u *EdgeUpsertOne) SetTo(v string) *EdgeUpsertOne {
	return u.Update(func(s *EdgeUpsert) {
		s.SetTo(v)
	})
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *EdgeUpsertOne) UpdateTo() *EdgeUpsertOne {
	return u.Update(func(s *EdgeUpsert) {
		s.UpdateTo()
	})
}

// Exec executes the query.
func (u *EdgeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EdgeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EdgeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EdgeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EdgeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EdgeCreateBulk is the builder for creating many Edge entities in bulk.
type EdgeCreateBulk struct {
	config
	err      error
	builders []*EdgeCreate
	conflict []sql.ConflictOption
}

// Save creates the Edge entities in the database.
func (ecb *EdgeCreateBulk) Save(ctx context.Context) ([]*Edge, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Edge, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EdgeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EdgeCreateBulk) SaveX(ctx context.Context) []*Edge {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EdgeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EdgeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Edge.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EdgeUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (ecb *EdgeCreateBulk) OnConflict(opts ...sql.ConflictOption) *EdgeUpsertBulk {
	ecb.conflict = opts
	return &EdgeUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Edge.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EdgeCreateBulk) OnConflictColumns(columns ...string) *EdgeUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EdgeUpsertBulk{
		create: ecb,
	}
}

// EdgeUpsertBulk is the builder for "upsert"-ing
// a bulk of Edge nodes.
type EdgeUpsertBulk struct {
	create *EdgeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Edge.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EdgeUpsertBulk) UpdateNewValues() *EdgeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Edge.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EdgeUpsertBulk) Ignore() *EdgeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EdgeUpsertBulk) DoNothing() *EdgeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EdgeCreateBulk.OnConflict
// documentation for more info.
func (u *EdgeUpsertBulk) Update(set func(*EdgeUpsert)) *EdgeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EdgeUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *EdgeUpsertBulk) SetType(v edge.Type) *EdgeUpsertBulk {
	return u.Update(func(s *EdgeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EdgeUpsertBulk) UpdateType() *EdgeUpsertBulk {
	return u.Update(func(s *EdgeUpsert) {
		s.UpdateType()
	})
}

// SetFrom sets the "from" field.
func (u *EdgeUpsertBulk) SetFrom(v string) *EdgeUpsertBulk {
	return u.Update(func(s *EdgeUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *EdgeUpsertBulk) UpdateFrom() *EdgeUpsertBulk {
	return u.Update(func(s *EdgeUpsert) {
		s.UpdateFrom()
	})
}

// SetTo sets the "to" field.
func (u *EdgeUpsertBulk) SetTo(v string) *EdgeUpsertBulk {
	return u.Update(func(s *EdgeUpsert) {
		s.SetTo(v)
	})
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *EdgeUpsertBulk) UpdateTo() *EdgeUpsertBulk {
	return u.Update(func(s *EdgeUpsert) {
		s.UpdateTo()
	})
}

// Exec executes the query.
func (u *EdgeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EdgeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EdgeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EdgeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
