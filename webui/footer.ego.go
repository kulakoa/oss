// Generated by ego.
// DO NOT EDIT

//line footer.ego:1

package webui

import "fmt"
import "html"
import "io"
import "context"

import (
	"github.com/contribsys/faktory/client"
	"net/http"
)

func ego_footer(w io.Writer, req *http.Request) {

//line footer.ego:11
	_, _ = io.WriteString(w, "\n<footer class=\"bg-dark\">\n  <div class=\"container-xl text-center\">\n    <div class=\"navbar navbar-fixed-bottom navbar-inverse navbar-dark justify-content-center justify-content-md-start\">\n      <div class=\"navbar-expand-md\">\n        <ul class=\"navbar-nav\">\n          <li class=\"nav-item me-md-4\">\n            <p class=\"navbar-text mb-0 product-version\">")
//line footer.ego:17
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(client.Name)))
//line footer.ego:17
	_, _ = io.WriteString(w, " ")
//line footer.ego:17
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(client.Version)))
//line footer.ego:17
	_, _ = io.WriteString(w, "</p>\n          </li>\n          <li class=\"nav-item me-md-4\">\n            <p class=\"navbar-text mb-0 faktory-url\">")
//line footer.ego:20
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(serverLocation(req))))
//line footer.ego:20
	_, _ = io.WriteString(w, "</p>\n          </li>\n          <li class=\"nav-item me-md-4\">\n            <p class=\"navbar-text mb-0 server-utc-time\">")
//line footer.ego:23
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(serverUtcTime())))
//line footer.ego:23
	_, _ = io.WriteString(w, "</p>\n          </li>\n          <li class=\"nav-item me-md-4\">\n            <p class=\"navbar-text mb-0\">")
//line footer.ego:26
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(ctx(req).Server().Options.Environment)))
//line footer.ego:26
	_, _ = io.WriteString(w, "</p>\n          </li>\n          <li class=\"nav-item me-md-4\">\n            <p class=\"navbar-text mb-0\"><a style=\"color: #666\" href=\"")
//line footer.ego:29
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(relative(req, "/debug"))))
//line footer.ego:29
	_, _ = io.WriteString(w, "\">debug</a></p>\n          </li>\n          <li class=\"nav-item me-md-4\">\n            <p class=\"navbar-text mb-0\"><a style=\"color: #666\" href=\"https://github.com/contribsys/faktory/wiki/\">docs</a></p>\n          </li>\n        </ul>\n      </div>\n    </div>\n  </div>\n</footer>\n")
//line footer.ego:39
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
