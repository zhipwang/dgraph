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
	"fmt"
	"time"

	"github.com/dgraph-io/dgraph/protos/graphp"
	"github.com/dgraph-io/dgraph/types"
	geom "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

func ValueFromGeoJson(json string, nq *graphp.NQuad) error {
	var g geom.T
	// Parse the json
	err := geojson.Unmarshal([]byte(json), &g)
	if err != nil {
		return err
	}

	geo, err := types.ObjectValue(types.GeoID, g)
	if err != nil {
		return err
	}

	nq.ObjectValue = geo
	nq.ObjectType = int32(types.GeoID)
	return nil
}

func Date(date time.Time, nq *graphp.NQuad) error {
	d, err := types.ObjectValue(types.DateID, date)
	if err != nil {
		return err
	}
	nq.ObjectValue = d
	nq.ObjectType = int32(types.DateID)
	return nil
}

func Datetime(date time.Time, nq *graphp.NQuad) error {
	d, err := types.ObjectValue(types.DateTimeID, date)
	if err != nil {
		return err
	}
	nq.ObjectValue = d
	nq.ObjectType = int32(types.DateTimeID)
	return nil
}

func Str(val string, nq *graphp.NQuad) error {
	v, err := types.ObjectValue(types.StringID, val)
	if err != nil {
		return err
	}
	nq.ObjectValue = v
	nq.ObjectType = int32(types.StringID)
	return nil
}

func Int(val int32, nq *graphp.NQuad) error {
	v, err := types.ObjectValue(types.Int32ID, val)
	if err != nil {
		return err
	}
	nq.ObjectValue = v
	nq.ObjectType = int32(types.Int32ID)
	return nil

}

func Float(val float64, nq *graphp.NQuad) error {
	v, err := types.ObjectValue(types.FloatID, val)
	if err != nil {
		return err
	}
	nq.ObjectValue = v
	nq.ObjectType = int32(types.FloatID)
	return nil

}

func Bool(val bool, nq *graphp.NQuad) error {
	v, err := types.ObjectValue(types.BoolID, val)
	if err != nil {
		return err
	}
	nq.ObjectValue = v
	nq.ObjectType = int32(types.BoolID)
	return nil
}

// Uid converts an uint64 to a string, which can be used as part of
// Subject and ObjectId fields in the graphp.NQuad
func Uid(uid uint64) string {
	return fmt.Sprintf("%#x", uid)
}

func Password(val string, nq *graphp.NQuad) error {
	v, err := types.ObjectValue(types.PasswordID, val)
	if err != nil {
		return err
	}
	nq.ObjectValue = v
	nq.ObjectType = int32(types.PasswordID)
	return nil
}
