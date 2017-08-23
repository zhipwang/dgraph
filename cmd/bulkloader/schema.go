package main

import (
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

type schemaStore struct {
	m map[string]*protos.SchemaUpdate
}

func newSchemaStore() schemaStore {
	return schemaStore{
		map[string]*protos.SchemaUpdate{
			"_predicate_": nil,
			"_lease_":     &protos.SchemaUpdate{ValueType: uint32(protos.Posting_INT)},
		}}
}

func (s schemaStore) add(nq *protos.NQuad) {

	fmt.Printf("NQuad: %#v\n\n", nq)

	sch := &protos.SchemaUpdate{
		ValueType: uint32(nq.GetObjectType()),
	}
	if nq.GetObjectValue() == nil {
		// RDF parser doesn't seem to pick up that objects that are nodes
		// should have UID value type.
		sch.ValueType = uint32(protos.Posting_UID)
	}
	s.m[nq.GetPredicate()] = sch
}

func (s schemaStore) write(kv *badger.KV) {
	fmt.Println("Schema:")
	for pred, sch := range s.m {
		fmt.Printf("%s: %#v\n", pred, sch)
		k := x.SchemaKey(pred)
		var v []byte
		var err error
		if sch != nil {
			v, err = sch.Marshal()
			x.Check(err)
		}
		x.Check(kv.Set(k, v, 0))
	}
	fmt.Println()
}
