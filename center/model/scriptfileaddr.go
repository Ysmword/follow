package model

import "errors"

type ScriptFileAddr struct {
	Addr        string `toml:"addr"`
	RunFilename string `toml:"runFilename"`
}

func (s *ScriptFileAddr) Check() error {
	if s.Addr == "" {
		return errors.New("addr expect")
	}
	if s.RunFilename == "" {
		return errors.New("run filename expect")
	}
	return nil
}
