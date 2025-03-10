package driver

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	FS interface {
		List(context.Context, string, ...Meta) ([]File, error)
		Move(context.Context, string, string) error
		Copy(context.Context, string, string) error
		Rename(context.Context, string, string) error
		Remove(context.Context, string) error
		MakeDir(context.Context, string) error
		Get(context.Context, string) (File, error)
		Open(string) (FileReader, error)
		Create(string) (FileWriter, error)
		Close() error
	}

	Option interface {
		NewFS() (FS, error)
	}
	OptionCreator func() Option
)

type Base struct{}

func (Base) List(context.Context, string, ...Meta) ([]File, error) { return nil, ErrNotSupport }
func (Base) Move(context.Context, string, string) error            { return ErrNotSupport }
func (Base) Copy(context.Context, string, string) error            { return ErrNotSupport }
func (Base) Rename(context.Context, string, string) error          { return ErrNotSupport }
func (Base) Remove(context.Context, string) error                  { return ErrNotSupport }
func (Base) MakeDir(context.Context, string) error                 { return ErrNotSupport }
func (Base) Get(context.Context, string) (File, error)             { return nil, ErrNotSupport }
func (Base) Open(string) (FileReader, error)                       { return nil, ErrNotSupport }
func (Base) Create(string) (FileWriter, error)                     { return nil, ErrNotSupport }
func (Base) Close() error                                          { return nil }

var allOptions map[string]OptionCreator

func DriverFS(driver string, option string) (FS, error) {
	creator, ok := allOptions[driver]
	if !ok {
		return nil, fmt.Errorf("The driver %s not found", driver)
	}
	opt := creator()
	if err := json.Unmarshal([]byte(option), opt); err != nil {
		return nil, err
	}
	return opt.NewFS()
}

func Verify(driver string, option string) error {
	creator, ok := allOptions[driver]
	if !ok {
		return fmt.Errorf("The driver %s not found", driver)
	}
	opt := creator()
	if err := json.Unmarshal([]byte(option), opt); err != nil {
		return err
	}
	return VerifyOption(opt)
}

func Exists(driver string) bool {
	_, ok := allOptions[driver]
	return ok
}

func Register(typ string, creator OptionCreator) {
	allOptions[typ] = creator
}

func init() {
	allOptions = make(map[string]OptionCreator)
}
