// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv1/conversion"
	"entgo.io/ent/schema/field"
)

// ConversionCreate is the builder for creating a Conversion entity.
type ConversionCreate struct {
	config
	mutation *ConversionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ConversionCreate) SetName(s string) *ConversionCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableName(s *string) *ConversionCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetInt8ToString sets the "int8_to_string" field.
func (cc *ConversionCreate) SetInt8ToString(i int8) *ConversionCreate {
	cc.mutation.SetInt8ToString(i)
	return cc
}

// SetNillableInt8ToString sets the "int8_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableInt8ToString(i *int8) *ConversionCreate {
	if i != nil {
		cc.SetInt8ToString(*i)
	}
	return cc
}

// SetUint8ToString sets the "uint8_to_string" field.
func (cc *ConversionCreate) SetUint8ToString(u uint8) *ConversionCreate {
	cc.mutation.SetUint8ToString(u)
	return cc
}

// SetNillableUint8ToString sets the "uint8_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableUint8ToString(u *uint8) *ConversionCreate {
	if u != nil {
		cc.SetUint8ToString(*u)
	}
	return cc
}

// SetInt16ToString sets the "int16_to_string" field.
func (cc *ConversionCreate) SetInt16ToString(i int16) *ConversionCreate {
	cc.mutation.SetInt16ToString(i)
	return cc
}

// SetNillableInt16ToString sets the "int16_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableInt16ToString(i *int16) *ConversionCreate {
	if i != nil {
		cc.SetInt16ToString(*i)
	}
	return cc
}

// SetUint16ToString sets the "uint16_to_string" field.
func (cc *ConversionCreate) SetUint16ToString(u uint16) *ConversionCreate {
	cc.mutation.SetUint16ToString(u)
	return cc
}

// SetNillableUint16ToString sets the "uint16_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableUint16ToString(u *uint16) *ConversionCreate {
	if u != nil {
		cc.SetUint16ToString(*u)
	}
	return cc
}

// SetInt32ToString sets the "int32_to_string" field.
func (cc *ConversionCreate) SetInt32ToString(i int32) *ConversionCreate {
	cc.mutation.SetInt32ToString(i)
	return cc
}

// SetNillableInt32ToString sets the "int32_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableInt32ToString(i *int32) *ConversionCreate {
	if i != nil {
		cc.SetInt32ToString(*i)
	}
	return cc
}

// SetUint32ToString sets the "uint32_to_string" field.
func (cc *ConversionCreate) SetUint32ToString(u uint32) *ConversionCreate {
	cc.mutation.SetUint32ToString(u)
	return cc
}

// SetNillableUint32ToString sets the "uint32_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableUint32ToString(u *uint32) *ConversionCreate {
	if u != nil {
		cc.SetUint32ToString(*u)
	}
	return cc
}

// SetInt64ToString sets the "int64_to_string" field.
func (cc *ConversionCreate) SetInt64ToString(i int64) *ConversionCreate {
	cc.mutation.SetInt64ToString(i)
	return cc
}

// SetNillableInt64ToString sets the "int64_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableInt64ToString(i *int64) *ConversionCreate {
	if i != nil {
		cc.SetInt64ToString(*i)
	}
	return cc
}

// SetUint64ToString sets the "uint64_to_string" field.
func (cc *ConversionCreate) SetUint64ToString(u uint64) *ConversionCreate {
	cc.mutation.SetUint64ToString(u)
	return cc
}

// SetNillableUint64ToString sets the "uint64_to_string" field if the given value is not nil.
func (cc *ConversionCreate) SetNillableUint64ToString(u *uint64) *ConversionCreate {
	if u != nil {
		cc.SetUint64ToString(*u)
	}
	return cc
}

// Mutation returns the ConversionMutation object of the builder.
func (cc *ConversionCreate) Mutation() *ConversionMutation {
	return cc.mutation
}

// Save creates the Conversion in the database.
func (cc *ConversionCreate) Save(ctx context.Context) (*Conversion, error) {
	return withHooks[*Conversion, ConversionMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConversionCreate) SaveX(ctx context.Context) *Conversion {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConversionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConversionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConversionCreate) check() error {
	return nil
}

func (cc *ConversionCreate) sqlSave(ctx context.Context) (*Conversion, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConversionCreate) createSpec() (*Conversion, *sqlgraph.CreateSpec) {
	var (
		_node = &Conversion{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(conversion.Table, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(conversion.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Int8ToString(); ok {
		_spec.SetField(conversion.FieldInt8ToString, field.TypeInt8, value)
		_node.Int8ToString = value
	}
	if value, ok := cc.mutation.Uint8ToString(); ok {
		_spec.SetField(conversion.FieldUint8ToString, field.TypeUint8, value)
		_node.Uint8ToString = value
	}
	if value, ok := cc.mutation.Int16ToString(); ok {
		_spec.SetField(conversion.FieldInt16ToString, field.TypeInt16, value)
		_node.Int16ToString = value
	}
	if value, ok := cc.mutation.Uint16ToString(); ok {
		_spec.SetField(conversion.FieldUint16ToString, field.TypeUint16, value)
		_node.Uint16ToString = value
	}
	if value, ok := cc.mutation.Int32ToString(); ok {
		_spec.SetField(conversion.FieldInt32ToString, field.TypeInt32, value)
		_node.Int32ToString = value
	}
	if value, ok := cc.mutation.Uint32ToString(); ok {
		_spec.SetField(conversion.FieldUint32ToString, field.TypeUint32, value)
		_node.Uint32ToString = value
	}
	if value, ok := cc.mutation.Int64ToString(); ok {
		_spec.SetField(conversion.FieldInt64ToString, field.TypeInt64, value)
		_node.Int64ToString = value
	}
	if value, ok := cc.mutation.Uint64ToString(); ok {
		_spec.SetField(conversion.FieldUint64ToString, field.TypeUint64, value)
		_node.Uint64ToString = value
	}
	return _node, _spec
}

// ConversionCreateBulk is the builder for creating many Conversion entities in bulk.
type ConversionCreateBulk struct {
	config
	builders []*ConversionCreate
}

// Save creates the Conversion entities in the database.
func (ccb *ConversionCreateBulk) Save(ctx context.Context) ([]*Conversion, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Conversion, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConversionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConversionCreateBulk) SaveX(ctx context.Context) []*Conversion {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConversionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConversionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
