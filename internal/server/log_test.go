package server_test

import (
	"reflect"
	"testing"

	"github.com/gabeno/logspace/internal/server"
)

func TestLog(t *testing.T) {
	t.Run("append record to empty Log", func(t *testing.T) {
		log := server.NewLog()
		record := server.Record{
			Value: []byte("alpha"),
		}

		log.Append(record)

		got, err := log.Read(0)

		if err != nil {
			t.Errorf("did not expect an error")
		}
		if !reflect.DeepEqual(got, record) {
			t.Errorf("got %v, want %v", got, record)
		}

	})

	t.Run("read from an empty log", func(t *testing.T) {
		log := server.NewLog()

		_, err := log.Read(0)

		if err == nil {
			t.Errorf("expected an error")
		}

		if err != server.ErrOffsetNotFound {
			t.Errorf("got %s expected %s", err, server.ErrOffsetNotFound)
		}
	})
}
