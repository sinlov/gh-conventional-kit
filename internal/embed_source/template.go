package embed_source

import (
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/bar-counter/slog"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	RenderStatusShow = "success"
	RenderStatusHide = "failure"
)

// Render parses and executes a template, returning the results in string
// format. Trailing or leading spaces or new-lines are not getting truncated. It
// is able to read templates from remote paths, local files or directly from the
// string.
func Render(template string, payload interface{}) (s string, err error) {
	u, err := url.Parse(template)

	if err == nil {
		switch u.Scheme {
		case "http", "https":
			res, errHttp := http.Get(template)

			if errHttp != nil {
				return s, fmt.Errorf("failed to fetch: %w", err)
			}

			defer func(Body io.ReadCloser) {
				errBodyClose := Body.Close()
				if errBodyClose != nil {
					slog.Errorf(errBodyClose, "failed to close body")
				}
			}(res.Body)

			out, errIoRead := io.ReadAll(res.Body)

			if errIoRead != nil {
				return s, fmt.Errorf("failed to read: %w", errIoRead)
			}

			template = string(out)
		case "file":
			out, errReadFile := os.ReadFile(u.Path)

			if errReadFile != nil {
				return s, fmt.Errorf("failed to read: %w", errReadFile)
			}

			template = string(out)
		}
	}

	return raymond.Render(template, payload)
}
