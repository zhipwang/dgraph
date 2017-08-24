package main

import (
	"fmt"
	"strconv"

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

	if sch, ok := s.m[nq.GetPredicate()]; ok {
		if nq.ObjectType == int32(protos.Posting_DEFAULT) {
			convertFromDefaultType(nq, sch)
		}
	} else {
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
}

func convertFromDefaultType(nq *protos.NQuad, sch *protos.SchemaUpdate) {
	x.AssertTruef(nq.ObjectType == int32(protos.Posting_DEFAULT), "nquad object must have default type")
	switch protos.Posting_ValType(sch.ValueType) {
	case protos.Posting_INT:
		i, err := strconv.ParseInt(nq.GetObjectValue().GetDefaultVal(), 10, 64)
		if err != nil {
			// Conversion failed. Leave NQuad as is (with default type).
			//
			// TODO: This doesn't seem correct to me given the documentation,
			// it should be rejected. But it seems to be what dgraph_lodaer
			// does.
			return
		}
		nq.GetObjectValue().Val = &protos.Value_IntVal{IntVal: i}
	case protos.Posting_UID, protos.Posting_DEFAULT:
		// Don't have to do any special conversions.
	default:
		// TODO: Other cases. Or better yet, code that already does this.
		x.AssertTrue(false)
	}
	nq.ObjectType = int32(sch.GetValueType())
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
