/*
 * Copyright 2015 DGraph Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 		http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dgraph-io/dgraph/antlr4go/graphqlpm"

	"github.com/dgraph-io/dgraph/algo"
	"github.com/dgraph-io/dgraph/gql"
	"github.com/dgraph-io/dgraph/group"
	"github.com/dgraph-io/dgraph/posting"
	"github.com/dgraph-io/dgraph/query/graph"
	"github.com/dgraph-io/dgraph/schema"
	"github.com/dgraph-io/dgraph/store"
	"github.com/dgraph-io/dgraph/task"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/worker"
	"github.com/dgraph-io/dgraph/x"
)

func childAttrs(sg *SubGraph) []string {
	var out []string
	for _, c := range sg.Children {
		out = append(out, c.Attr)
	}
	return out
}

func taskValues(t *testing.T, v []*task.Value) []string {
	out := make([]string, len(v))
	for i, tv := range v {
		out[i] = string(tv.Val)
	}
	return out
}

func TestNewGraph(t *testing.T) {
	dir, err := ioutil.TempDir("", "storetest_")
	require.NoError(t, err)

	gq := &gql.GraphQuery{
		UID:  101,
		Attr: "me",
	}
	ps, err := store.NewStore(dir)
	require.NoError(t, err)

	posting.Init(ps)

	ctx := context.Background()
	sg, err := newGraph(ctx, gq)
	require.NoError(t, err)

	require.EqualValues(t,
		[][]uint64{
			[]uint64{101},
		}, algo.ToUintsListForTest(sg.uidMatrix))
}

const schemaStr = `
scalar name:string @index
scalar dob:date @index
scalar loc:geo @index
`

func addEdgeToValue(t *testing.T, ps *store.Store, attr string, src uint64,
	value string) {
	edge := &task.DirectedEdge{
		Value:  []byte(value),
		Label:  "testing",
		Attr:   attr,
		Entity: src,
	}
	l, _ := posting.GetOrCreate(x.DataKey(attr, src))
	require.NoError(t,
		l.AddMutationWithIndex(context.Background(), edge, posting.Set))
}

func addEdgeToTypedValue(t *testing.T, ps *store.Store, attr string, src uint64,
	typ types.TypeID, value []byte) {
	edge := &task.DirectedEdge{
		Value:     value,
		ValueType: uint32(typ),
		Label:     "testing",
		Attr:      attr,
		Entity:    src,
	}
	l, _ := posting.GetOrCreate(x.DataKey(attr, src))
	require.NoError(t,
		l.AddMutationWithIndex(context.Background(), edge, posting.Set))
}

func addEdgeToUID(t *testing.T, ps *store.Store, attr string, src uint64, dst uint64) {
	edge := &task.DirectedEdge{
		ValueId: dst,
		Label:   "testing",
		Attr:    attr,
		Entity:  src,
	}
	l, _ := posting.GetOrCreate(x.DataKey(attr, src))
	require.NoError(t,
		l.AddMutationWithIndex(context.Background(), edge, posting.Set))
}

func populateGraph(t *testing.T) (string, string, *store.Store) {
	// logrus.SetLevel(logrus.DebugLevel)
	dir, err := ioutil.TempDir("", "storetest_")
	require.NoError(t, err)

	ps, err := store.NewStore(dir)
	require.NoError(t, err)

	schema.ParseBytes([]byte(schemaStr))
	posting.Init(ps)
	worker.Init(ps)

	group.ParseGroupConfig("")
	dir2, err := ioutil.TempDir("", "wal_")
	require.NoError(t, err)
	worker.StartRaftNodes(dir2)

	// So, user we're interested in has uid: 1.
	// She has 5 friends: 23, 24, 25, 31, and 101
	addEdgeToUID(t, ps, "friend", 1, 23)
	addEdgeToUID(t, ps, "friend", 1, 24)
	addEdgeToUID(t, ps, "friend", 1, 25)
	addEdgeToUID(t, ps, "friend", 1, 31)
	addEdgeToUID(t, ps, "friend", 1, 101)

	// Now let's add a few properties for the main user.
	addEdgeToValue(t, ps, "name", 1, "Michonne")
	addEdgeToValue(t, ps, "gender", 1, "female")
	var coord types.Geo
	err = coord.UnmarshalText([]byte("{\"Type\":\"Point\", \"Coordinates\":[1.1,2.0]}"))
	require.NoError(t, err)
	gData, err := coord.MarshalBinary()
	require.NoError(t, err)
	addEdgeToTypedValue(t, ps, "loc", 1, types.GeoID, gData)
	data, err := types.Int32(15).MarshalBinary()
	require.NoError(t, err)

	addEdgeToTypedValue(t, ps, "age", 1, types.Int32ID, data)
	addEdgeToValue(t, ps, "address", 1, "31, 32 street, Jupiter")
	data, err = types.Bool(true).MarshalBinary()
	require.NoError(t, err)
	addEdgeToTypedValue(t, ps, "alive", 1, types.BoolID, data)
	addEdgeToValue(t, ps, "age", 1, "38")
	addEdgeToValue(t, ps, "survival_rate", 1, "98.99")
	addEdgeToValue(t, ps, "sword_present", 1, "true")
	addEdgeToValue(t, ps, "_xid_", 1, "mich")

	// Now let's add a name for each of the friends, except 101.
	addEdgeToTypedValue(t, ps, "name", 23, types.StringID, []byte("Rick Grimes"))
	addEdgeToValue(t, ps, "age", 23, "15")

	err = coord.UnmarshalText([]byte(`{"Type":"Polygon", "Coordinates":[[[0.0,0.0], [2.0,0.0], [2.0, 2.0], [0.0, 2.0]]]}`))
	require.NoError(t, err)
	gData, err = coord.MarshalBinary()
	require.NoError(t, err)
	addEdgeToTypedValue(t, ps, "loc", 23, types.GeoID, gData)

	addEdgeToValue(t, ps, "address", 23, "21, mark street, Mars")
	addEdgeToValue(t, ps, "name", 24, "Glenn Rhee")
	addEdgeToValue(t, ps, "name", 25, "Daryl Dixon")
	addEdgeToValue(t, ps, "name", 31, "Andrea")

	addEdgeToValue(t, ps, "dob", 23, "1910-01-02")
	addEdgeToValue(t, ps, "dob", 24, "1909-05-05")
	addEdgeToValue(t, ps, "dob", 25, "1909-01-10")
	addEdgeToValue(t, ps, "dob", 31, "1901-01-15")

	return dir, dir2, ps
}

func processToJSON(t *testing.T, query string) string {
	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)
	sg.DebugPrint("")

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)
	sg.DebugPrint("")

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	return string(js)
}

func TestGetUID(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:"0x01") {
				name
				_uid_
				gender
				alive
				friend {
					_uid_
					name
				}
			}
		}
	`
	js := processToJSON(t, query)
	require.JSONEq(t,
		`{"me":[{"_uid_":"0x1","alive":true,"friend":[{"_uid_":"0x17","name":"Rick Grimes"},{"_uid_":"0x18","name":"Glenn Rhee"},{"_uid_":"0x19","name":"Daryl Dixon"},{"_uid_":"0x1f","name":"Andrea"},{"_uid_":"0x65"}],"gender":"female","name":"Michonne"}]}`,
		js)
}

func TestGetUIDNotInChild(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:"0x01") {
				name
				_uid_
				gender
				alive
				friend {
					name
				}
			}
		}
	`
	js := processToJSON(t, query)
	require.JSONEq(t,
		`{"me":[{"_uid_":"0x1","alive":true,"friend":[{"name":"Rick Grimes"},{"name":"Glenn Rhee"},{"name":"Daryl Dixon"},{"name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		js)
}

func TestGetUIDCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:"0x01") {
				name
				_uid_
				gender
				alive
				friend {
					_count_
				}
			}
		}
	`
	js := processToJSON(t, query)
	require.JSONEq(t,
		`{"me":[{"_uid_":"0x1","alive":true,"friend":[{"_count_":5}],"gender":"female","name":"Michonne"}]}`,
		js)
}

func TestDebug1(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			debug(_uid_:"0x01") {
				name
				gender
				alive
				friend {
					_count_
				}
			}
		}
	`

	js := processToJSON(t, query)
	var mp map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(js), &mp))

	resp := mp["debug"]
	uid := resp.([]interface{})[0].(map[string]interface{})["_uid_"].(string)
	require.EqualValues(t, "0x1", uid)

	latency := mp["server_latency"]
	require.NotNil(t, latency)
	_, ok := latency.(map[string]interface{})
	require.True(t, ok)
}

func TestDebug2(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:"0x01") {
				name
				gender
				alive
				friend {
					_count_
				}
			}
		}
	`

	js := processToJSON(t, query)
	var mp map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(js), &mp))

	resp := mp["me"]
	uid, ok := resp.([]interface{})[0].(map[string]interface{})["_uid_"].(string)
	require.False(t, ok, "No uid expected but got one %s", uid)
}

func TestCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:"0x01") {
				name
				gender
				alive
				friend {
					_count_
				}
			}
		}
	`

	js := processToJSON(t, query)
	require.EqualValues(t,
		`{"me":[{"alive":true,"friend":[{"_count_":5}],"gender":"female","name":"Michonne"}]}`,
		js)
}

func TestCountError1(t *testing.T) {
	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_: "0x01") {
				friend {
					name
					_count_
				}
				name
				gender
				alive
			}
		}
	`
	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	_, err = ToSubGraph(ctx, gq)
	require.Error(t, err)
}

func TestCountError2(t *testing.T) {
	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_: "0x01") {
				friend {
					_count_ {
						friend
					}
				}
				name
				gender
				alive
			}
		}
	`
	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	_, err = ToSubGraph(ctx, gq)
	require.Error(t, err)
}

func TestProcessGraph(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_: "0x01") {
				friend {
					name
				}
				name
				gender
				alive	
			}
		}
	`
	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	require.EqualValues(t, childAttrs(sg), []string{"friend", "name", "gender", "alive"})
	require.EqualValues(t, childAttrs(sg.Children[0]), []string{"name"})

	child := sg.Children[0]
	require.EqualValues(t,
		[][]uint64{
			[]uint64{23, 24, 25, 31, 101},
		}, algo.ToUintsListForTest(child.uidMatrix))

	require.EqualValues(t, []string{"name"}, childAttrs(child))

	child = child.Children[0]
	require.EqualValues(t,
		[]string{"Rick Grimes", "Glenn Rhee", "Daryl Dixon", "Andrea", ""},
		taskValues(t, child.values))

	require.EqualValues(t, []string{"Michonne"},
		taskValues(t, sg.Children[1].values))
	require.EqualValues(t, []string{"female"},
		taskValues(t, sg.Children[2].values))
}

func TestToJSON(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:"0x01") {
				name
				gender
			  alive
				friend {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.JSONEq(t,
		`{"me":[{"alive":true,"friend":[{"name":"Rick Grimes"},{"name":"Glenn Rhee"},{"name":"Daryl Dixon"},{"name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestFieldAlias(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:"0x01") {
				MyName:name
				gender
				alive
				Buddies:friend {
					BudName:name
				}
			}
		}
	`

	js := processToJSON(t, query)
	require.JSONEq(t,
		`{"me":[{"alive":true,"Buddies":[{"BudName":"Rick Grimes"},{"BudName":"Glenn Rhee"},{"BudName":"Daryl Dixon"},{"BudName":"Andrea"}],"gender":"female","MyName":"Michonne"}]}`,
		string(js))
}

func TestFieldAliasProto(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:"0x01") {
				MyName:name
				gender
				alive
				Buddies:friend {
					BudName:name
				}
			}
		}
	`
	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)
	fmt.Println(proto.MarshalTextString(pb))
	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "MyName"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  properties: <
    prop: "alive"
    value: <
      bool_val: true
    >
  >
  children: <
    attribute: "Buddies"
    properties: <
      prop: "BudName"
      value: <
        str_val: "Rick Grimes"
      >
    >
  >
  children: <
    attribute: "Buddies"
    properties: <
      prop: "BudName"
      value: <
        str_val: "Glenn Rhee"
      >
    >
  >
  children: <
    attribute: "Buddies"
    properties: <
      prop: "BudName"
      value: <
        str_val: "Daryl Dixon"
      >
    >
  >
  children: <
    attribute: "Buddies"
    properties: <
      prop: "BudName"
      value: <
        str_val: "Andrea"
      >
    >
  >
>
`
	require.EqualValues(t,
		expectedPb,
		proto.MarshalTextString(pb))
}

func TestToJSONFilter(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:"0x01") {
				name
				gender
				friend @filter(anyof("name", "Andrea SomethingElse")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)
	sg.DebugPrint("  ")

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterAllOf(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(allof("name", "Andrea SomethingElse")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterUID(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea")) {
					_uid_
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"_uid_":"0x1f"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrUID(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea") || anyof("name", "Andrea Rhee")) {
					_uid_
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"_uid_":"0x18","name":"Glenn Rhee"},{"_uid_":"0x1f","name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea") || anyof("name", "Andrea Rhee")) {
					_count_
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.JSONEq(t,
		`{"me":[{"friend":[{"_count_":2}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrFirst(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(first:2) @filter(anyof("name", "Andrea") || anyof("name", "Glenn SomethingElse") || anyof("name", "Daryl")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Glenn Rhee"},{"name":"Daryl Dixon"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrOffset(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(offset:1) @filter(anyof("name", "Andrea") || anyof("name", "Glenn Rhee") || anyof("name", "Daryl Dixon")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Daryl Dixon"},{"name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

// No filter. Just to test first and offset.
func TestToJSONFirstOffset(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(offset:1, first:1) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Glenn Rhee"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrFirstOffset(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(offset:1, first:1) @filter(anyof("name", "Andrea") || anyof("name", "SomethingElse Rhee") || anyof("name", "Daryl Dixon")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Daryl Dixon"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrFirstOffsetCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(offset:1, first:1) @filter(anyof("name", "Andrea") || anyof("name", "SomethingElse Rhee") || anyof("name", "Daryl Dixon")) {
					_count_
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.JSONEq(t,
		`{"me":[{"friend":[{"_count_":1}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterOrFirstNegative(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	// When negative first/count is specified, we ignore offset and returns the last
	// few number of items.
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(first:-1, offset:0) @filter(anyof("name", "Andrea") || anyof("name", "Glenn Rhee") || anyof("name", "Daryl Dixon")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func TestToJSONFilterAnd(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea") && anyof("name", "SomethingElse Rhee")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"gender":"female","name":"Michonne"}]}`,
		string(js))
}

func getProperty(properties []*graph.Property, prop string) *graph.Value {
	for _, p := range properties {
		if p.Prop == prop {
			return p.Value
		}
	}
	return nil
}

func TestToProto(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			debug(_uid_:0x1) {
				_xid_
				name
				gender
				alive
				friend {
					name
				}
				friend {
				}
			}
		}
  `

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	gr, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	require.EqualValues(t,
		`attribute: "_root_"
children: <
  uid: 1
  xid: "mich"
  attribute: "debug"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  properties: <
    prop: "alive"
    value: <
      bool_val: true
    >
  >
  children: <
    uid: 23
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Rick Grimes"
      >
    >
  >
  children: <
    uid: 24
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Glenn Rhee"
      >
    >
  >
  children: <
    uid: 25
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Daryl Dixon"
      >
    >
  >
  children: <
    uid: 31
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Andrea"
      >
    >
  >
  children: <
    uid: 101
    attribute: "friend"
  >
>
`, proto.MarshalTextString(gr))
}

func TestSchema(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			debug(_uid_:0x1) {
				_xid_
				name
				gender
				alive
				loc
				friend {
					name
				}
				friend {
				}
			}
		}
  `

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	gr, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	require.EqualValues(t, "debug", gr.Children[0].Attribute)
	require.EqualValues(t, 1, gr.Children[0].Uid)
	require.EqualValues(t, "mich", gr.Children[0].Xid)
	require.Len(t, gr.Children[0].Properties, 4)

	require.EqualValues(t, "Michonne",
		getProperty(gr.Children[0].Properties, "name").GetStrVal())
	var g types.Geo
	x.Check(g.UnmarshalBinary(getProperty(gr.Children[0].Properties, "loc").GetGeoVal()))
	received, err := g.MarshalText()
	require.EqualValues(t, "{'type':'Point','coordinates':[1.1,2]}", string(received))

	require.Len(t, gr.Children[0].Children, 5)

	child := gr.Children[0].Children[0]
	require.EqualValues(t, 23, child.Uid)
	require.EqualValues(t, "friend", child.Attribute)

	require.Len(t, child.Properties, 1)
	require.EqualValues(t, "Rick Grimes",
		getProperty(child.Properties, "name").GetStrVal())
	require.Empty(t, child.Children)

	child = gr.Children[0].Children[1]
	require.EqualValues(t, 24, child.Uid)
	require.EqualValues(t, "friend", child.Attribute)

	require.Len(t, child.Properties, 1)
	require.EqualValues(t, "Glenn Rhee",
		getProperty(child.Properties, "name").GetStrVal())
	require.Empty(t, child.Children)

	child = gr.Children[0].Children[4]
	require.EqualValues(t, 101, child.Uid)
	require.EqualValues(t, "friend", child.Attribute)

	require.Len(t, child.Properties, 0)
}

func TestToProtoFilter(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Andrea"
      >
    >
  >
>
`
	require.EqualValues(t, expectedPb, proto.MarshalTextString(pb))
}

func TestToProtoFilterOr(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea") || anyof("name", "Glenn Rhee")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Glenn Rhee"
      >
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Andrea"
      >
    >
  >
>
`
	require.EqualValues(t, proto.MarshalTextString(pb), expectedPb)
}

func TestToProtoFilterAnd(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend @filter(anyof("name", "Andrea") && anyof("name", "Glenn Rhee")) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
>
`
	require.EqualValues(t, expectedPb, proto.MarshalTextString(pb))
}

// Test sorting / ordering by dob.
func TestToJSONOrder(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(order: dob) {
					name
					dob
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"dob":"1901-01-15","name":"Andrea"},{"dob":"1909-01-10","name":"Daryl Dixon"},{"dob":"1909-05-05","name":"Glenn Rhee"},{"dob":"1910-01-02","name":"Rick Grimes"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

// Test sorting / ordering by dob.
func TestToJSONOrderDesc(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(orderdesc: dob) {
					name
					dob
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"dob":"1910-01-02","name":"Rick Grimes"},{"dob":"1909-05-05","name":"Glenn Rhee"},{"dob":"1909-01-10","name":"Daryl Dixon"},{"dob":"1901-01-15","name":"Andrea"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

// Test sorting / ordering by dob.
func TestToJSONOrderOffset(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(order: dob, offset: 2) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Glenn Rhee"},{"name":"Rick Grimes"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

// Test sorting / ordering by dob.
func TestToJSONOrderOffsetCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(order: dob, offset: 2, first: 1) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	js, err := sg.ToJSON(&l)
	require.NoError(t, err)
	require.EqualValues(t,
		`{"me":[{"friend":[{"name":"Glenn Rhee"}],"gender":"female","name":"Michonne"}]}`,
		string(js))
}

// Test sorting / ordering by dob.
func TestToProtoOrder(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(order: dob) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Andrea"
      >
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Daryl Dixon"
      >
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Glenn Rhee"
      >
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Rick Grimes"
      >
    >
  >
>
`
	require.EqualValues(t, expectedPb, proto.MarshalTextString(pb))
}

// Test sorting / ordering by dob.
func TestToProtoOrderCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(order: dob, first: 2) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Andrea"
      >
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Daryl Dixon"
      >
    >
  >
>
`
	require.EqualValues(t, expectedPb, proto.MarshalTextString(pb))
}

// Test sorting / ordering by dob.
func TestToProtoOrderOffsetCount(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
		{
			me(_uid_:0x01) {
				name
				gender
				friend(order: dob, first: 2, offset: 1) {
					name
				}
			}
		}
	`

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
  properties: <
    prop: "gender"
    value: <
      bytes_val: "female"
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Daryl Dixon"
      >
    >
  >
  children: <
    attribute: "friend"
    properties: <
      prop: "name"
      value: <
        str_val: "Glenn Rhee"
      >
    >
  >
>
`
	require.EqualValues(t, expectedPb, proto.MarshalTextString(pb))
}

func TestSchema1(t *testing.T) {
	require.NoError(t, schema.Parse("test_schema"))

	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	// Alright. Now we have everything set up. Let's create the query.
	query := `
		{
			person(_uid_:0x01) {
				alive
				survival_rate
				friend
			}
		}
	`
	js := processToJSON(t, query)
	require.JSONEq(t,
		`{"person":[{"address":"31, 32 street, Jupiter","age":38,"alive":true,"friend":[{"address":"21, mark street, Mars","age":15,"name":"Rick Grimes"}],"name":"Michonne","survival_rate":98.99}]}`,
		js)
}

func TestGenerator(t *testing.T) {
	dir1, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)
	query := `
    {
      me(anyof("name", "Michonne")) {
        name
        gender
      }
    }
  `
	js := processToJSON(t, query)
	require.JSONEq(t, `{"me":[{"gender":"female","name":"Michonne"}]}`, js)
}

func TestGeneratorMultiRoot(t *testing.T) {
	dir1, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)
	query := `
    {
      me(anyof("name", "Michonne Rick Glenn")) {
        name
      }
    }
  `
	js := processToJSON(t, query)
	require.JSONEq(t, `{"me":[{"name":"Michonne"},{"name":"Rick Grimes"},{"name":"Glenn Rhee"}]}`, js)
}

func TestToProtoMultiRoot(t *testing.T) {
	dir, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir)
	defer os.RemoveAll(dir2)

	query := `
    {
      me(anyof("name", "Michonne Rick Glenn")) {
        name
      }
    }
  `

	gq, _, err := gql.Parse(query)
	require.NoError(t, err)

	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	require.NoError(t, err)

	ch := make(chan error)
	go ProcessGraph(ctx, sg, nil, ch)
	err = <-ch
	require.NoError(t, err)

	var l Latency
	pb, err := sg.ToProtocolBuffer(&l)
	require.NoError(t, err)

	expectedPb := `attribute: "_root_"
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Michonne"
    >
  >
>
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Rick Grimes"
    >
  >
>
children: <
  attribute: "me"
  properties: <
    prop: "name"
    value: <
      str_val: "Glenn Rhee"
    >
  >
>
`
	require.EqualValues(t, expectedPb, proto.MarshalTextString(pb))
}

func TestNearGenerator(t *testing.T) {
	dir1, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)
	query := `{
		me(near("loc", "{'Type':'Point', 'Coordinates':[1.1,2.0]}", "5")) {
			name
			gender
		}
	}`

	js := processToJSON(t, query)
	require.JSONEq(t, `{"me":[{"gender":"female","name":"Michonne"}]}`, string(js))
}

func TestWithinGenerator(t *testing.T) {
	dir1, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)
	query := `{
		me(within("loc",  "{'Type':'Polygon', 'Coordinates':[[[0.0,0.0], [2.0,0.0], [1.5, 3.0], [0.0, 2.0]]]}")) {
			name
		}
	}`

	js := processToJSON(t, query)
	require.JSONEq(t, `{"me":[{"name":"Michonne"}]}`, string(js))
}

func TestContainsGenerator(t *testing.T) {
	dir1, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)
	query := `{
		me(contains("loc",  "{'Type':'Point', 'Coordinates':[2.0, 0.0]}")) {
			name
		}
	}`

	js := processToJSON(t, query)
	require.JSONEq(t, `{"me":[{"name":"Rick Grimes"}]}`, string(js))
}

func TestIntersectsGenerator(t *testing.T) {
	dir1, dir2, _ := populateGraph(t)
	defer os.RemoveAll(dir1)
	defer os.RemoveAll(dir2)
	query := `{
		me(intersects("loc", "{'Type':'Polygon', 'Coordinates':[[[0.0,0.0], [2.0,0.0], [1.5, 3.0], [0.0, 2.0]]]}")) {
			name
		}
	}`

	js := processToJSON(t, query)
	require.JSONEq(t, `{"me":[{"name":"Michonne"}, {"name":"Rick Grimes"}]}`, string(js))
}

func TestMain(m *testing.M) {
	x.Init()
	os.Exit(m.Run())
}

var q1 = `
{
	al(_xid_: alice) {
		status
		_xid_
		follows {
			status
			_xid_
			follows {
				status
				_xid_
				follows {
					_xid_
					status
				}
			}
		}
		status
		_xid_
	}
}
`
var q2 = `
	query queryName {
		me(_uid_:0x0a) {
			friends {
				name
			}
			gender,age
			hometown
		}
	}
`

var q3 = `
{
  debug(_xid_: m.0bxtg) {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        film.film.directed_by {
          type.object.name.en
        }
      }
    }
  }
}
`
var q4 = `
{
  debug(_xid_: m.0c6qh) {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        type.object.name.en
      }
    }
  }
}
`

// queries to benchmark on.
var q5 = `{
  debug(_xid_: m.0f4vbz) {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        type.object.name.en
      }
    }
  }
}`

var q6 = `{
  debug(_xid_: m.06pj8) {
    type.object.name.en
    film.director.film  {
      film.film.genre {
        type.object.name.en
      }
    }
  }
}`

var q7 = `{
  debug(_xid_: m.0c6qh) {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        type.object.name.en
      }
    }
  }
}`

var q8 = `{
  debug(_xid_: m.0bxtg) {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        film.film.directed_by {
          type.object.name.en
        }
      }
    }
  }
}`

func benchmarkQueryParse(q string, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gq, _, err := gql.Parse(q)
		if err != nil {
			b.Error(err)
			return
		}
		ctx := context.Background()
		_, err = ToSubGraph(ctx, gq)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkQueryParse(b *testing.B) {
	b.Run("q1", func(b *testing.B) { benchmarkQueryParse(q1, b) })
	b.Run("q2", func(b *testing.B) { benchmarkQueryParse(q2, b) })
	b.Run("q3", func(b *testing.B) { benchmarkQueryParse(q3, b) })
	b.Run("q4", func(b *testing.B) { benchmarkQueryParse(q4, b) })
}

var aq1 = `
{
	al(_xid_: "alice") {
		status
		_xid_
		follows {
			status
			_xid_
			follows {
				status
				_xid_
				follows {
					_xid_
					status
				}
			}
		}
		status
		_xid_
	}
}
`

var aq2 = `query queryName {
		me(_uid_ : "0x0a") {
			friends {
				name
			}
			gender,age
			hometown
		}
	}
`

var aq3 = `
{
  debug(_xid_: "m.0bxtg") {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        film.film.directed_by {
          type.object.name.en
        }
      }
    }
  }
}
`

var aq4 = `
{
  debug(_xid_: "m.0c6qh") {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        type.object.name.en
      }
    }
  }
}
`

var aq5 = `{
  debug(_xid_: "m.0f4vbz") {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        type.object.name.en
      }
    }
  }
}`

var aq6 = `{
  debug(_xid_: "m.06pj8") {
    type.object.name.en
    film.director.film  {
      film.film.genre {
        type.object.name.en
      }
    }
  }
}`

var aq7 = `{
  debug(_xid_: "m.0c6qh") {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        type.object.name.en
      }
    }
  }
}`

var aq8 = `{
  debug(_xid_: "m.0bxtg") {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        film.film.directed_by {
          type.object.name.en
        }
      }
    }
  }
}`

// visitor ---------------------
type gQVisitor struct {
	*parser.BaseGraphQLPMVisitor
}

type stringValue struct {
	val string
}

func (this *gQVisitor) VisitStringValue(ctx *parser.StringValueContext) interface{} {
	return stringValue{ctx.STRING().GetText()}
}

func (this *gQVisitor) Visit(t antlr.ParseTree) *gql.GraphQuery {
	return nil
}

// visitor ends ---------------------

type gQListener struct {
	*parser.BaseGraphQLPMListener
	gQPropertyMap map[antlr.RuleContext]*gql.GraphQuery
}

func (l *gQListener) setGQ(ctx antlr.RuleContext, gQ *gql.GraphQuery) {
	if ctx != nil {
		l.gQPropertyMap[ctx] = gQ
	}
}
func (l *gQListener) getGQ(ctx antlr.RuleContext) *gql.GraphQuery {
	return l.gQPropertyMap[ctx]
}

// Context is a ParseNode in java/c++ ; how to make gQPropertyMap a map from ParseNode to gql.GraphQuery
// func (l *gQListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
// 	gqNode := new(gql.GraphQuery)
// 	for i := 0; i < ctx.GetChildCount(); i++ {
// 		child := ctx.GetChild(i)
// 		if child.(antlr.RuleNode) != nil {
// 			gqChild := l.getGQ(child.(antlr.RuleNode).GetRuleContext())
// 			if gqChild != nil {
// 				gqNode.Children = append(gqNode.Children, gqChild)
// 			} else {
// 				fmt.Println("nil node here. :(")
// 			}
// 		}
// 	}
// 	l.setGQ(ctx.GetRuleContext(), gqNode)
// }

func newGQListener() *gQListener {
	r := new(gQListener)
	r.gQPropertyMap = make(map[antlr.RuleContext]*gql.GraphQuery)
	return r
}

func TestQueryParse12(t *testing.T) {
	gq, _, err := gql.Parse(q8)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()
	sg, err := ToSubGraph(ctx, gq)
	if err != nil {
		t.Error(err)
		return
	}
	sg.DebugPrint("")
}

func TestQueryParse11(t *testing.T) {
	input := antlr.NewInputStream(aq8)
	lexer := parser.NewGraphQLPMLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGraphQLPMParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	listener := newGQListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	agq := listener.getGQ(tree)

	ctx := context.Background()
	sq, err := ToSubGraph(ctx, agq)
	require.NoError(t, err)
	sq.DebugPrint("")

	// visitor := new(gQVisitor)
	// if visitor != nil {
	// 	fmt.Println("not nil")
	// 	// do not know but giving me null-pointer exception
	// 	// visitor.Visit(tree)
	// }
}

// going top down over DocumentContext tree to generate graphquery would be so much easier..
// if i could call DefinitionContext over d..
// func gengql.GraphQuery(doc parser.DocumentContext) (gq *gql.GraphQuery, err error) {
// 	gq = &gql.GraphQuery{
// 		Args: make(map[string]string),
// 	}
// 	for _, d := range doc.AllDefinition() {
// 		def := d.(parser.DefinitionContext)
// 		selSet := def.SelectionSet()
// 		if selSet == nil {
// 			def.OperationDefinition().SelectionSet()
// 		}
// 	}
// }

// ExitDocument is called when production document is exited.
func (l *gQListener) ExitDocument(ctx *parser.DocumentContext) {
	def := ctx.Definition()
	qChild := l.getGQ(def.GetRuleContext())
	l.setGQ(ctx, qChild.Children[0])
	fmt.Println(len(l.gQPropertyMap))
}

// ExitDefinition is called when production definition is exited.
func (l *gQListener) ExitDefinition(ctx *parser.DefinitionContext) {
	var sset antlr.RuleContext
	sset = ctx.SelectionSet()
	if sset == nil {
		sset = ctx.OperationDefinition()
	}
	l.setGQ(ctx, l.getGQ(sset))
}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (l *gQListener) ExitOperationDefinition(ctx *parser.OperationDefinitionContext) {
	l.setGQ(ctx, l.getGQ(ctx.SelectionSet()))
}

// ExitOperationType is called when production operationType is exited.
func (l *gQListener) ExitOperationType(ctx *parser.OperationTypeContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (l *gQListener) ExitSelectionSet(ctx *parser.SelectionSetContext) {
	q := new(gql.GraphQuery)
	q.Args = make(map[string]string)
	for _, f := range ctx.AllField() {
		fGQ := l.getGQ(f.GetRuleContext())
		q.Children = append(q.Children, fGQ)
	}
	l.setGQ(ctx, q)
}

// ExitField is called when production field is exited.
func (l *gQListener) ExitField(ctx *parser.FieldContext) {
	q := new(gql.GraphQuery)
	q.Args = make(map[string]string)
	q.Attr = ctx.NAME().GetText()
	if args := ctx.Arguments(); args != nil {
		q.Args = l.getGQ(args).Args
		if xid, ok := q.Args["_xid_"]; ok {
			q.XID = xid
		}
		if uid, ok := q.Args["_uid_"]; ok {
			q.UID, _ = strconv.ParseUint(uid, 0, 64)
		}
	}
	if sset := ctx.SelectionSet(); sset != nil {
		q.Children = l.getGQ(sset).Children
	}
	l.setGQ(ctx, q)
}

// ExitArguments is called when production arguments is exited.
func (l *gQListener) ExitArguments(ctx *parser.ArgumentsContext) {
	q := new(gql.GraphQuery)
	q.Args = make(map[string]string)
	for _, arg := range ctx.AllArgument() {
		argGQ := l.getGQ(arg.GetRuleContext())
		for k, v := range argGQ.Args {
			q.Args[k] = v
		}
	}
	l.setGQ(ctx, q)
}

// ExitArgument is called when production argument is exited.
func (l *gQListener) ExitArgument(ctx *parser.ArgumentContext) {
	q := new(gql.GraphQuery)
	q.Args = make(map[string]string)
	q.Args[ctx.NAME().GetText()] = l.getGQ(ctx.Value().GetRuleContext()).XID
	fmt.Println(ctx.NAME().GetText(), l.getGQ(ctx.Value().GetRuleContext()).XID)
	l.setGQ(ctx, q)
}

func (l *gQListener) ExitStringValue(ctx *parser.StringValueContext) {
	fmt.Println(ctx.STRING().GetText())
	q := new(gql.GraphQuery)
	// check if this is uid or xid ?
	q.XID = ctx.STRING().GetText()
	// q.UID = strconv.ParseUint(ctx.STRING().GetText(), 0, 64)
	l.setGQ(ctx, q)
}
