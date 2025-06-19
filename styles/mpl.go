package styles

import (
	. "github.com/crowyy03/chroma/v2"
)

// MPL style.
var MPL = Register(MustNewStyle("mpl", StyleEntries{
    Background:       "#23272e",
    Text:             "#e0e6f0",
    Error:            "#ff5370",
    Comment:          "#5c6370",
    CommentPreproc:   "#82aaff",
    Keyword:          "#c792ea",
    KeywordType:      "#ffcb6b",
    Operator:         "#89ddff",
    Punctuation:      "#89ddff",
    Name:             "#e0e6f0",
    NameBuiltin:      "#f78c6c",
    NameClass:        "#ffcb6b",
    NameFunction:     "#82aaff",
    NameNamespace:    "#b2ccd6",
    NameVariable:     "#c3e88d",
    LiteralString:    "#c3e88d",
    LiteralNumber:    "#f78c6c",
    LiteralNumberHex: "#ff5370",
    TextWhitespace:   "#444a57",
}))
