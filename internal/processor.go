package internal

import (
	"io"

	"github.com/pkg/errors"
)

//Translator implements the Translate Method for translating the bytes from source to target language.
type Translator interface {
	Translate(b []byte) ([]byte, error)
}

//Processor embeds the different interfaces needed for processing a file.
type Processor struct {
	io.Reader
	Translator
}

//Run runs the process (as defined in the Processor struct) on the file in path.
//Here it reads file from path and translates it from source language (sl) to target language(tl)
func (p Processor) Run() ([]byte, error) {
	b := make([]byte, 0)
	_, err := p.Read(b)
	if err != nil {
		return nil, errors.Wrap(err, "reading file")
	}
	tb, err := p.Translate(b)
	if err != nil {
		return nil, errors.Wrapf(err, "translate file")
	}
	return tb, nil
}
