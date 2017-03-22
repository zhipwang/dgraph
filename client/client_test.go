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

package client

import (
	"testing"

	"github.com/dgraph-io/dgraph/protos/graphp"
	"github.com/stretchr/testify/assert"
)

func graphValue(x string) *graphp.Value {
	return &graphp.Value{&graphp.Value_StrVal{x}}
}

func TestCheckNQuad(t *testing.T) {
	s := graphValue("Alice")
	if err := checkNQuad(graphp.NQuad{
		Predicate:   "name",
		ObjectValue: s,
	}); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad(graphp.NQuad{
		Subject:     "alice",
		ObjectValue: s,
	}); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad(graphp.NQuad{
		Subject:   "alice",
		Predicate: "name",
	}); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad(graphp.NQuad{
		Subject:     "alice",
		Predicate:   "name",
		ObjectValue: s,
		ObjectId:    "id",
	}); err == nil {
		t.Fatal(err)
	}
}

func TestSetMutation(t *testing.T) {
	req := Req{}

	s := graphValue("Alice")
	if err := req.AddMutation(graphp.NQuad{
		Subject:     "alice",
		Predicate:   "name",
		ObjectValue: s,
	}, SET); err != nil {
		t.Fatal(err)
	}

	s = graphValue("rabbithole")
	if err := req.AddMutation(graphp.NQuad{
		Subject:     "alice",
		Predicate:   "falls.in",
		ObjectValue: s,
	}, SET); err != nil {
		t.Fatal(err)
	}

	if err := req.AddMutation(graphp.NQuad{
		Subject:     "alice",
		Predicate:   "falls.in",
		ObjectValue: s,
	}, DEL); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(req.gr.Mutation.Set), 2, "Set should have 2 entries")
	assert.Equal(t, len(req.gr.Mutation.Del), 1, "Del should have 1 entry")
}
