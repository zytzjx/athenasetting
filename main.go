package main

import (
	//dmc "github.com/zytzjx/anthenacmc/datacentre"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/zytzjx/anthenacmc/companysetting"
	Log "github.com/zytzjx/anthenacmc/loggersys"
	_ "github.com/zytzjx/anthenacmc/loggersys"
)

var lock sync.Mutex

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Marshal is a function that marshals the object into an
	// io.Reader.
	// By default, it uses the JSON marshaller.
	var Marshal = func(v interface{}) (io.Reader, error) {
		b, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(b), nil
	}
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func main() {
	Log.NewLogger("setting")
	Log.NewLogger("version:20.11.30.0  by Jefferyzhang")
	setting, err := companysetting.Download()
	if err != nil {
		Log.Log.Error(err)
		os.Exit(2)
	}
	// Greent no support features, why respone??
	// delete(setting["settings"].(map[string]interface{}), "features")
	if err = Save("setting.json", setting); err != nil {
		Log.Log.Error(err)
		os.Exit(3)
	}

	// if something need be saved to DB, TODO

	//Log.Log.Info("file start")
	//gitdmc.
}
