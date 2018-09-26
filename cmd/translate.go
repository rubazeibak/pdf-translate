package main

import (
	"github.com/sruba90z/pdf-translate/internal"
	"github.com/sruba90z/pdf-translate/pkg/google"
	"golang.org/x/text/language"
)

func newTranslator(sl, tl string) internal.Translator {
	return &google.Translator{
		Sl: language.Make(sl),
		Tl: language.Make(tl),
	}
}
