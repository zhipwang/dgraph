/*
 * Copyright (C) 2017 Dgraph Labs, Inc. and Contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package schema

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dgraph-io/dgraph/group"
	"github.com/dgraph-io/dgraph/protos/typesp"
	"github.com/dgraph-io/dgraph/store"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/x"
)

type nameType struct {
	name string
	typ  *typesp.Schema
}

func checkSchema(t *testing.T, h map[string]*typesp.Schema, expected []nameType) {
	require.Len(t, h, len(expected))
	for _, nt := range expected {
		typ, found := h[nt.name]
		require.True(t, found, nt)
		require.EqualValues(t, *nt.typ, *typ)
	}
}

func TestSchema(t *testing.T) {
	require.NoError(t, ReloadData("testfiles/test_schema", 1))
	checkSchema(t, State().get(1).predicate, []nameType{
		{"name", &typesp.Schema{ValueType: uint32(types.StringID)}},
		{"address", &typesp.Schema{ValueType: uint32(types.StringID)}},
		{"http://scalar.com/helloworld/", &typesp.Schema{ValueType: uint32(types.StringID)}},
		{"age", &typesp.Schema{ValueType: uint32(types.Int32ID)}},
	})

	typ, err := State().TypeOf("age")
	require.NoError(t, err)
	require.Equal(t, types.Int32ID, typ)

	typ, err = State().TypeOf("agea")
	require.Error(t, err)
}

func TestSchema1_Error(t *testing.T) {
	require.Error(t, ReloadData("testfiles/test_schema1", 1))
}

func TestSchema2_Error(t *testing.T) {
	require.Error(t, ReloadData("testfiles/test_schema2", 1))
}

func TestSchema3_Error(t *testing.T) {
	require.Error(t, ReloadData("testfiles/test_schema3", 1))
}

/*
func TestSchema5_Error(t *testing.T) {
	str = make(map[string]types.TypeID)
	err := Parse("testfiles/test_schema5")
	require.Error(t, err)
}

func TestSchema6_Error(t *testing.T) {
	str = make(map[string]types.TypeID)
	err := Parse("testfiles/test_schema6")
	require.Error(t, err)
}
*/
// Correct specification of indexing
func TestSchemaIndex(t *testing.T) {
	require.NoError(t, ReloadData("testfiles/test_schema_index1", 1))
	require.Equal(t, 2, len(State().IndexedFields(1)))
}

// Indexing can't be specified inside object types.
func TestSchemaIndex_Error1(t *testing.T) {
	require.Error(t, ReloadData("testfiles/test_schema_index2", 1))
}

// Object types cant be indexed.
func TestSchemaIndex_Error2(t *testing.T) {
	require.Error(t, ReloadData("testfiles/test_schema_index5", 1))
}

// Missing comma.
func TestSchemaIndex_Error3(t *testing.T) {
	require.Error(t, ReloadData("testfiles/test_schema_index3", 1))
}

func TestSchemaIndexCustom(t *testing.T) {
	require.NoError(t, ReloadData("testfiles/test_schema_index4", 1))
	checkSchema(t, State().get(1).predicate, []nameType{
		{"name", &typesp.Schema{ValueType: uint32(types.StringID), Tokenizer: []string{"exact"}}},
		{"address", &typesp.Schema{ValueType: uint32(types.StringID), Tokenizer: []string{"term"}}},
		{"age", &typesp.Schema{ValueType: uint32(types.Int32ID), Tokenizer: []string{"int"}}},
		{"id", &typesp.Schema{ValueType: uint32(types.StringID), Tokenizer: []string{"exact", "term"}}},
	})
	require.True(t, State().IsIndexed("name"))
	require.False(t, State().IsReversed("name"))
	require.Equal(t, "int", State().Tokenizer("age")[0].Name())
	require.Equal(t, 4, len(State().IndexedFields(1)))
}

var ps *store.Store

func TestMain(m *testing.M) {
	x.SetTestRun()
	x.Init()

	dir, err := ioutil.TempDir("", "storetest_")
	x.Check(err)
	ps, err = store.NewStore(dir)
	x.Check(err)
	x.Check(group.ParseGroupConfig("groups.conf"))
	Init(ps)
	defer os.RemoveAll(dir)
	defer ps.Close()

	os.Exit(m.Run())
}
