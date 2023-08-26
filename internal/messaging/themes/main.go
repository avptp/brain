package themes

import (
	"embed"
	"fmt"
)

//go:embed *.html *.txt
var fs embed.FS

type Theme string

func (t *Theme) Name() string {
	if t == nil {
		return ""
	}

	return string(*t)
}

func (t *Theme) HTMLTemplate() string {
	return t.read("html")
}

func (t *Theme) PlainTextTemplate() string {
	return t.read("txt")
}

func (t *Theme) read(e string) string {
	file, err := fs.ReadFile(
		fmt.Sprintf("%s.%s", t.Name(), e),
	)

	if err != nil {
		panic(err) // unrecoverable situation
	}

	return string(file)
}
