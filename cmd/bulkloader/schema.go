package main

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/worker"
	"github.com/dgraph-io/dgraph/x"
)

type schemaState struct {
	strict bool
	*protos.SchemaUpdate
}

type schemaStore struct {
	m map[string]schemaState
}

func newSchemaStore(initial []*protos.SchemaUpdate) schemaStore {
	s := schemaStore{
		map[string]schemaState{
			"_predicate_": {true, nil},
			"_lease_":     {true, &protos.SchemaUpdate{ValueType: uint32(protos.Posting_INT)}},
		},
	}
	for _, sch := range initial {
		p := sch.Predicate
		sch.Predicate = ""
		s.m[p] = schemaState{true, sch}
	}
	return s
}

func (s schemaStore) fixEdge(de *protos.DirectedEdge, isUIDEdge bool) {

	if isUIDEdge {
		de.ValueType = uint32(protos.Posting_UID)
	}

	sch, ok := s.m[de.Attr]
	if !ok {
		sch = schemaState{false, &protos.SchemaUpdate{ValueType: de.ValueType}}
		s.m[de.Attr] = sch
	}

	schTyp := types.TypeID(sch.ValueType)
	err := worker.ValidateAndConvert(de, schTyp)
	if sch.strict && err != nil {
		// TODO: It's unclear to me as to why it's only an error to have a bad
		// conversion if the schema was established explicitly rather than
		// automatically. I suspect I have a misunderstanding about how things
		// should work.
		fmt.Printf("BAD RDF: %v\n", err) // TODO: bubble back properly
		os.Exit(1)
	}
}

func (s schemaStore) write(kv *badger.KV) {
	fmt.Println("Schema:")
	for pred, sch := range s.m {
		fmt.Printf("%s: %+v\n", pred, sch)
		k := x.SchemaKey(pred)
		var v []byte
		var err error
		if sch.SchemaUpdate != nil {
			v, err = sch.SchemaUpdate.Marshal()
			x.Check(err)
		}
		x.Check(kv.Set(k, v, 0))
	}
	fmt.Println()
}
