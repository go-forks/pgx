package pgtype_test

import (
	"math"
	"testing"

	"github.com/jackc/pgx/pgtype"
)

func TestInt8Transcode(t *testing.T) {
	testSuccessfulTranscode(t, "int8", []interface{}{
		pgtype.Int8{Int: math.MinInt64, Status: pgtype.Present},
		pgtype.Int8{Int: -1, Status: pgtype.Present},
		pgtype.Int8{Int: 0, Status: pgtype.Present},
		pgtype.Int8{Int: 1, Status: pgtype.Present},
		pgtype.Int8{Int: math.MaxInt64, Status: pgtype.Present},
		pgtype.Int8{Int: 0, Status: pgtype.Null},
	})
}

func TestInt8ConvertFrom(t *testing.T) {
	type _int8 int8

	successfulTests := []struct {
		source interface{}
		result pgtype.Int8
	}{
		{source: int8(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: int16(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: int32(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: int64(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: int8(-1), result: pgtype.Int8{Int: -1, Status: pgtype.Present}},
		{source: int16(-1), result: pgtype.Int8{Int: -1, Status: pgtype.Present}},
		{source: int32(-1), result: pgtype.Int8{Int: -1, Status: pgtype.Present}},
		{source: int64(-1), result: pgtype.Int8{Int: -1, Status: pgtype.Present}},
		{source: uint8(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: uint16(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: uint32(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: uint64(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: "1", result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
		{source: _int8(1), result: pgtype.Int8{Int: 1, Status: pgtype.Present}},
	}

	for i, tt := range successfulTests {
		var r pgtype.Int8
		err := r.ConvertFrom(tt.source)
		if err != nil {
			t.Errorf("%d: %v", i, err)
		}
	}
}