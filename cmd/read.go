package main

import (
	"io"

	"github.com/sruba90z/pdf-translate/pkg/pdf"
)

func newReader(path string) io.Reader {
	return &pdf.Reader{path: path}
}
