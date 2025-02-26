package nautils

import (
	"fmt"
	"strings"

	"github.com/jomei/notionapi"
)

type HTMLConverter struct{}

func (hc *HTMLConverter) ToHTML(blocks []notionapi.Block) string {
	var sb strings.Builder
	for _, b := range blocks {
		sb.WriteString(hc.Block(b))
	}
	return sb.String()
}

// HeadingToHTML converts a Heading struct to an HTML string.
func (hc *HTMLConverter) commonHeading(e notionapi.Heading, level int) string {
	var sb strings.Builder

	// Default heading tag
	tag := "h" + fmt.Sprint(level)
	if e.IsToggleable {
		tag = "details"
	}

	// Open heading tag
	sb.WriteString(fmt.Sprintf("<%s", tag))

	// Apply color if specified
	if e.Color != "" && e.Color != "default" {
		sb.WriteString(fmt.Sprintf(` style="color:%s;"`, e.Color))
	}

	sb.WriteString(">")

	// Convert RichText to HTML
	for _, rt := range e.RichText {
		sb.WriteString(hc.RichText(rt))
	}

	// Close heading tag
	sb.WriteString(fmt.Sprintf("</%s>", tag))

	return sb.String()
}

func (hc *HTMLConverter) Heading1(e notionapi.Heading1Block) string {
	return hc.commonHeading(e.Heading1, 1)
}

func (hc *HTMLConverter) Heading2(e notionapi.Heading2Block) string {
	return hc.commonHeading(e.Heading2, 2)
}

func (hc *HTMLConverter) Heading3(e notionapi.Heading3Block) string {
	return hc.commonHeading(e.Heading3, 3)
}

// RichTextToHTML converts a RichText struct to an HTML string.
func (hc *HTMLConverter) RichText(rt notionapi.RichText) string {
	if rt.Text == nil {
		return ""
	}

	text := rt.Text.Content
	if rt.Href != "" {
		text = fmt.Sprintf(`<a href="%s">%s</a>`, rt.Href, text)
	}

	// Apply annotations
	if rt.Annotations != nil {
		if rt.Annotations.Bold {
			text = fmt.Sprintf("<b>%s</b>", text)
		}
		if rt.Annotations.Italic {
			text = fmt.Sprintf("<i>%s</i>", text)
		}
		if rt.Annotations.Underline {
			text = fmt.Sprintf("<u>%s</u>", text)
		}
		if rt.Annotations.Strikethrough {
			text = fmt.Sprintf("<s>%s</s>", text)
		}
		if rt.Annotations.Code {
			text = fmt.Sprintf("<code>%s</code>", text)
		}
	}

	return text
}

// ParagraphToHTML converts a Paragraph struct to an HTML string.
func (hc *HTMLConverter) Paragraph(p notionapi.Paragraph) string {
	var sb strings.Builder

	// Open paragraph tag
	sb.WriteString("<p")

	// Apply color if specified
	if p.Color != "" && p.Color != "default" {
		sb.WriteString(fmt.Sprintf(` style="color:%s;"`, p.Color))
	}

	sb.WriteString(">")

	// Convert RichText to HTML
	for _, rt := range p.RichText {
		sb.WriteString(hc.RichText(rt))
	}

	// Close paragraph tag
	sb.WriteString("</p>")

	// Convert and append children if present
	if len(p.Children) > 0 {
		for _, child := range p.Children {
			sb.WriteString(hc.Block(child))
		}
	}

	return sb.String()
}

// BlockToHTML converts a generic Block interface to an HTML string.
func (hc *HTMLConverter) Block(b notionapi.Block) string {
	switch b.GetType() {
	case notionapi.BlockTypeHeading1:
		return hc.Heading1(*b.(*notionapi.Heading1Block))
	case notionapi.BlockTypeHeading2:
		return hc.Heading2(*b.(*notionapi.Heading2Block))
	case notionapi.BlockTypeHeading3:
		return hc.Heading3(*b.(*notionapi.Heading3Block))
	case notionapi.BlockTypeParagraph:
		return hc.Paragraph(b.(*notionapi.ParagraphBlock).Paragraph)
	}
	return ""
}
