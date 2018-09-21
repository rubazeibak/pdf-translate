package pdf

import (
	"fmt"

	"github.com/pkg/errors"
	"rsc.io/pdf"
)

//Reader implements io.Reader
type Reader struct {
	path    string
	content []byte
}

func (r *Reader) Read(p []byte) (n int, err error) {
	pr, err := pdf.Open(r.path)
	if err != nil {
		return 0, errors.Wrapf(err, "open pdf %s", r.path)
	}
	texts, err := r.getContent(pr)
	if err != nil {
		return 0, errors.Wrapf(err, "get content of pdf %s", r.path)
	}
	err = r.concatText(texts)
	return 0, nil
}

func (r *Reader) getContent(pr *pdf.Reader) ([][]pdf.Text, error) {
	texts := make([][]pdf.Text, 0)
	np := pr.NumPage()
	for n := 1; n <= np; n++ {
		page := pr.Page(n)
		if page.V.IsNull() {
			return texts, fmt.Errorf("page %v is null", n)
		}
		content := page.Content()
		texts = append(texts, content.Text)
	}
	return texts, nil
}

func (r *Reader) concatText(texts [][]pdf.Text) error {
	for n, text := range texts {
		 for _, t := range text {
			 content
		 }
	}
}
