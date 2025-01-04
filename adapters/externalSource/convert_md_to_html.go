package externalsource

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type ConvertMdToHtml interface {
	ConvertMdToHtml(md []byte) []byte
}

type ConvertMdToHtmlStruct struct {
}

func (c *ConvertMdToHtmlStruct) ConvertMdToHtml(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func NewConvertMdToHtmlStruct() ConvertMdToHtml {
	return &ConvertMdToHtmlStruct{}
}
