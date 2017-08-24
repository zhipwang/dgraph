package main

import (
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/worker"
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
		},
	}
}

func (s schemaStore) fixEdge(de *protos.DirectedEdge, isUIDEdge bool) {

	if isUIDEdge {
		de.ValueType = uint32(protos.Posting_UID)
	}

	schTyp := types.DefaultID
	if sch, ok := s.m[de.Attr]; ok {
		schTyp = types.TypeID(sch.ValueType)
	}

	_ = worker.ValidateAndConvert(de, schTyp)
	// A return error indicates that the conversion failed. Dgraph simply
	// continues on in that case, so we do as well to maintain compatibility.

	if _, ok := s.m[de.Attr]; !ok {
		s.m[de.Attr] = &protos.SchemaUpdate{ValueType: de.ValueType}
	}
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
