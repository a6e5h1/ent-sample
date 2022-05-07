package graphql

import (
	"io"
	"strconv"
	"time"

	"golang.org/x/xerrors"

	"github.com/99designs/gqlgen/graphql"
)

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	case time.Time:
		return v, nil
	default:
		return time.Now(), xerrors.New("DateTime is invalid")
	}
}

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(t.Format(time.RFC3339))))
	})
}
