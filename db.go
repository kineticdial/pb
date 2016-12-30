package pb

import (
	"encoding/gob"
	"fmt"
	"os"
)

func Put(k string, o interface{}) error {
	f, err := os.Create(filename(k))
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(f)
	err = enc.Encode(o)
	if err != nil {
		return err
	}

	f.Close()
	return nil
}

func Get(k string, o interface{}) error {
	f, err := os.Open(filename(k))
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(f)
	err = dec.Decode(o)
	if err != nil {
		return err
	}

	f.Close()
	return nil
}

func filename(k string) string {
	return fmt.Sprintf("./.pb/objects/%s", k)
}
