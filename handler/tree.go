package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

// TreeT stores nested handler.Info instances
type TreeT struct {
	Node     Info // Info may contain just a title string => not clickable
	Children []TreeT
}

// Tree returns an application specific nav tree of handler.Info
// based on which a navigation can be rendered;
// we dont use a package variable,
// since the nav tree may be modified by requests
func Tree() *TreeT {

	root := &TreeT{

		Node: Info{Title: "Root Node"},
		Children: []TreeT{
			{Node: infos.ByKey("create-anonymous-id")},
			{Node: Info{Title: "Sys admin"},
				Children: []TreeT{
					{Node: infos.ByKey("login-primitive")},
					{Node: infos.ByKey("logout")},
					{Node: infos.ByKey("change-password-primitive")},
					{Node: infos.ByKey("logins-reload")},
					{Node: infos.ByKey("config-reload")},
					{Node: infos.ByKey("templates-reload")},
					{Node: infos.ByKey("slow-hijacked")},
					{Node: infos.ByKey("slow-buffered")},
					{Node: infos.ByKey("session-get")},
					{Node: infos.ByKey("session-put")},
					{Node: infos.ByKey("cloud-store-test")},
					{Node: infos.ByKey("instance-info")},
					// {Node: infos.ByKey("pprof-index"), Children: []TreeT{{Node: infos.ByKey("pprof-symbol")}}},
					{Node: infos.ByKey("pprof-index"), Children: []TreeT{}},
					{Node: infos.ByKey("pprof-symbol")},
					// {Node: infos.ByKey("login-appengine")},
					// {Node: infos.ByKey("google-token-signin")},
				},
			},
		},
	}

	// if cfg.Get().ForwardCopyByCountry {
	//	addenum := infos.ByKey("export-import-directly")
	//	root.AppendAfterByKey("backup", &addenum)
	// }

	return root

}

// ByKey recursively retrieves a node;
// use ByKey() and SetByKey() to modify the nav tree
func (rt *TreeT) ByKey(key string) *TreeT {
	if rt.Node.HasKey(key) {
		return rt
	}
	for _, c := range rt.Children {
		if c.Node.HasKey(key) {
			return &c
		}
		nd := c.ByKey(key)
		if nd != nil {
			return nd
		}
	}
	return nil
}

// SetByKey recursively replaces an existing node;
// use ByKey() and SetByKey() to modify the nav tree
func (rt *TreeT) SetByKey(key string, repl *Info) bool {
	if rt.Node.HasKey(key) {
		rt.Node = *repl
		return true
	}
	for idx := range rt.Children {
		if rt.Children[idx].SetByKey(key, repl) { // we have to access via slice - to really change the underlying instance
			return true
		}
	}
	return false
}

// AppendAfterByKey recursively appends behind an existing node
func (rt *TreeT) AppendAfterByKey(key string, summand *Info) bool {
	// if rt.Node.HasKey(key) {
	//   level 0 root should never get appended
	// }
	for idx := range rt.Children {
		if rt.Children[idx].Node.HasKey(key) {
			ln := len(rt.Children)
			cp := make([]TreeT, 0, ln+1)
			cp = append(cp, rt.Children[0:idx+1]...)
			cp = append(cp, TreeT{Node: *summand})
			cp = append(cp, rt.Children[idx+1:ln]...)
			rt.Children = cp
			return true
		}
		if rt.Children[idx].AppendAfterByKey(key, summand) {
			return true
		}
	}
	return false
}

/*NavHTML - renders a navigation tree into
the HTML base snippet below;
it is called from templates via template func {{nav .Req}}
being mapped to tpl.fcNav

<nav>

    <!-- logo text or image are part of the design; they are set via CSS -->
    <div class="logo">
    </div>

    <!-- burger checkbox - out of sight - still can be tabbed and get focus -->
    <input type="checkbox" id="mnu-1st-lvl-toggler" class="mnu-1st-lvl-toggler">

    <!-- burger label -->
    <label for="mnu-1st-lvl-toggler" class="burger">
        <div class="line1"></div>
        <div class="line2"></div>
        <div class="line3"></div>
    </label>

    <!-- menu second level-->
    <ul class="mnu-2nd-lvl">

        {{ .Content}}

        <li><a href="#">Home</a></li>
        <li><a href="#">Blog</a></li>
        <li><a href="#">Contact</a></li>

        <li class='nde-2nd-lvl'>
            <a href='#' onclick='return false;'>Sys Admin</a>

            <ul class='mnu-3rd-lvl'>
                <li><a href='/login-primitive' class=''>Login app</a></li>
                <li><a href='/login-appengine' class=''>Login appengine</a></li>
                <li><a href='/google-token-signin' class=''>Google token sign-in</a></li>
            </ul>
        </li>
        <li class='nde-2nd-lvl'>
            <a href="#">About</a>

            <ul class='mnu-3rd-lvl'>
                <li><a href='/login-primitive' class=''>More</a></li>
                <li><a href='/login-primitive' class=''>Subitems</a></li>
            </ul>

        </li>

    </ul>


</nav>


*/
func (rt *TreeT) NavHTML(w io.Writer, r *http.Request, isLogin, isAdmin bool, lvl int) {

	// The root node itself is not rendered - we go straight to the children
	if lvl == 0 {
		for _, child := range rt.Children {
			child.NavHTML(w, r, isLogin, isAdmin, lvl+1)
		}
		return
	}

	needsLogout := rt.Node.Allow[LoggedOut]
	if needsLogout && isLogin {
		return
	}
	needsLogin := rt.Node.Allow[LoggedIn]
	if needsLogin && !isLogin {
		return
	}
	needsAdmin := rt.Node.Allow[Admin]
	if needsAdmin && !isAdmin {
		return
	}

	// if rt.Node.InNav {
	{

		htmlIndent := strings.Repeat(" ", 10*lvl) // just for readability in HTML source
		navURL := ""
		activeClass := "" // style the active nav item
		_ = activeClass
		activeStyle := "" // ...
		preventClck := "" // nav items without URL should not be clickable;

		accessKey := ""
		if rt.Node.ShortCut != "" {
			accessKey = fmt.Sprintf(" accesskey='%v' ", rt.Node.ShortCut)
			rt.Node.Title += fmt.Sprintf(" <span title='Keyboard shortcut SHIFT+ALT+%v' >(%v)</span>", rt.Node.ShortCut, rt.Node.ShortCut)
		}

		if len(rt.Node.Urls) > 0 {
			navURL = cfg.Pref(rt.Node.Urls[0])
		}

		// log.Printf("cmp %v to %v", strings.TrimSuffix(r.URL.Path, "/"), navURL)
		if navURL == strings.TrimSuffix(r.URL.Path, "/") {
			navURL = ""
			activeClass = " is-active "
			activeStyle = " style='font-weight: bold;' "
		}

		if navURL == "" {
			preventClck = " onclick='return false;' "
		}

		if len(rt.Children) == 0 {

			fmt.Fprintf(w, "\n%v<li><a href='%v'   %v  %v  %v  >%v</a></li>  \n",
				htmlIndent, navURL, activeStyle, preventClck, accessKey, rt.Node.Title,
			)

		} else {

			/*
				<li class='nde-2nd-lvl'>
					<a href='#' onclick='return false;'>Sys Admin</a>

					<ul class='mnu-3rd-lvl'>
						<li><a href='/login-primitive' class=''>Login app</a></li>
						...
						<li><a href='/google-token-signin' class=''>Google token sign-in</a></li>
					</ul>
				</li>

			*/

			fmt.Fprintf(w, "%v<li class='nde-2nd-lvl'>\n", htmlIndent)

			// same as above - without enclosing <li>
			fmt.Fprintf(w, "%v<a href='%v'   %v  %v  %v  >%v</a>  \n",
				htmlIndent, navURL, activeStyle, preventClck, accessKey, rt.Node.Title,
			)

			fmt.Fprintf(w, "%v     <ul class='mnu-3rd-lvl'>\n", htmlIndent)
			for _, child := range rt.Children {
				child.NavHTML(w, r, isLogin, isAdmin, lvl+1)
			}
			fmt.Fprintf(w, "%v     </ul>\n", htmlIndent)

			fmt.Fprintf(w, "%v</li>\n", htmlIndent)

		}

	}

}
