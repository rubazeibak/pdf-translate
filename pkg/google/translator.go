package google

import (
	"context"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

//Translator implements the internal translator interface
type Translator struct {
	Sl language.Tag
	Tl language.Tag
}

//Translate translates from source lang sl to target lang
func (t *Translator) Translate(content []byte) ([]byte, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "new translator client")
	}
	translations, err := client.Translate(ctx,
		[]string{string(content)}, t.Tl,
		&translate.Options{
			Source: t.Sl,
			Format: translate.Text,
		})
	if err != nil {
		return nil, errors.Wrap(err, "translate content")
	}

	return concatTranslations(translations), nil
}

func concatTranslations(trs []translate.Translation) []byte {
	translated := make([]string, 0)
	for _, translation := range trs {
		translated = append(translated, translation.Text)
	}
	concated := strings.Join(translated, " ")
	return []byte(concated)
}
