package hydra

import (
	"encoding/json"
	"io"
	"os"

	"golang.org/x/xerrors"
)

func Load(v interface{}, loaders ...Loader) (err error) {
	for i := range loaders {
		if err := loaders[i].Load(v); err != nil {
			return xerrors.Errorf("failed to load %d-th loader: %w", i, err)
		}
	}
	return nil
}

type Loader interface {
	Load(interface{}) error
}

type LoaderFunc func(interface{}) error

func (f LoaderFunc) Load(v interface{}) error {
	return f(v)
}

func JSONLoader(fn string) Loader {
	return LoaderFunc(func(v interface{}) (err error) {
		f, err := os.Open(fn)
		if err != nil {
			return xerrors.Errorf("failed to open file: %w", err)
		}
		defer func() {
			if err2 := f.Close(); err2 != nil && err == nil {
				err = err2
			}
		}()

		return JSONReaderLoader(f).Load(v)
	})
}

func JSONReaderLoader(f io.Reader) Loader {
	return LoaderFunc(func(v interface{}) (err error) {
		if err := json.NewDecoder(f).Decode(v); err != nil {
			return xerrors.Errorf("failed to decode json: %w", err)
		}
		return nil
	})
}
