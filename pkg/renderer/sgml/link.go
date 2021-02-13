package sgml

import (
	"html"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderLink(ctx *renderer.Context, l types.InlineLink) (string, error) { //nolint: unparam
	result := &strings.Builder{}
	location := l.Location.Stringify()
	text := ""
	class := ""
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render link")
	}
	// TODO; support `mailto:` positional attributes
	if t, exists := l.Attributes[types.AttrInlineLinkText]; exists {
		switch t := t.(type) {
		case string:
			text = t
		case []interface{}:
			var err error
			if text, err = r.renderInlineElements(ctx, t); err != nil {
				return "", errors.Wrap(err, "unable to render link")
			}
		}
		class = roles // can be empty (and it's fine)
	} else {
		text = html.EscapeString(location)
		if len(roles) > 0 {
			class = "bare " + roles
		} else {
			class = "bare"
		}
	}
	err = r.link.Execute(result, struct {
		URL    string
		Text   string
		Class  string
		Target string
	}{
		URL:    location,
		Text:   text,
		Class:  class,
		Target: l.Attributes.GetAsStringWithDefault(types.AttrInlineLinkTarget, ""),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render link")
	}
	// log.Debugf("rendered link: %s", result.String())
	return result.String(), nil
}
