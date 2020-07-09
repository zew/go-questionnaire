package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

// TreeT stores a tree of nested HandlerInfos
type TreeT struct {
	Node     Info // Info may contain just a title string => not clickable
	Children []TreeT
}

// Tree returns a recursive structure of handler.Info
func Tree() *TreeT {

	return &TreeT{
		Node: Info{Title: "Root Node", Visible: true},
		Children: []TreeT{
			{Node: infos.ByKey("create-anonymous-id")},
			{Node: Info{Title: "Sys admin", Visible: true},
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

}

/*NavHTML - CSS2020 renders navigation tree to ...

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
func (t *TreeT) NavHTML(w io.Writer, r *http.Request, isLogin, isAdmin bool, lvl int) {

	// The root node itself is not rendered - we go straight to the children
	if lvl == 0 {
		for _, child := range t.Children {
			child.NavHTML(w, r, isLogin, isAdmin, lvl+1)
		}
		return
	}

	needsLogout := t.Node.Allow[LoggedOut]
	if needsLogout && isLogin {
		return
	}
	needsLogin := t.Node.Allow[LoggedIn]
	if needsLogin && !isLogin {
		return
	}
	needsAdmin := t.Node.Allow[Admin]
	if needsAdmin && !isAdmin {
		return
	}

	if t.Node.Visible {

		htmlIndent := strings.Repeat(" ", 10*lvl) // just for readability in HTML source
		navURL := ""
		activeClass := "" // style the active nav item
		_ = activeClass
		activeStyle := "" // ...
		preventClck := "" // nav items without URL should not be clickable;

		if len(t.Node.Urls) > 0 {
			navURL = cfg.Pref(t.Node.Urls[0])
		}

		// log.Printf("cmp %v to %v", r.URL.Path, navURL)
		if navURL == r.URL.Path {
			navURL = ""
			activeClass = " is-active "
			activeStyle = " style='font-weight: bold;' "
		}

		if navURL == "" {
			preventClck = " onclick='return false;' "
		}

		if len(t.Children) == 0 {

			/*                <li><a href="#"             >Up</a></li>  */
			fmt.Fprintf(w, "\n%v<li><a href='%v'   %v  %v   >%v</a></li>  \n", htmlIndent, navURL, activeStyle, preventClck, t.Node.Title)

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
			fmt.Fprintf(w, "%v<a href='%v'   %v  %v   >%v</a>  \n", htmlIndent, navURL, activeStyle, preventClck, t.Node.Title)

			fmt.Fprintf(w, "%v     <ul class='mnu-3rd-lvl'>\n", htmlIndent)
			for _, child := range t.Children {
				child.NavHTML(w, r, isLogin, isAdmin, lvl+1)
			}
			fmt.Fprintf(w, "%v     </ul>\n", htmlIndent)

			fmt.Fprintf(w, "%v</li>\n", htmlIndent)

		}

	}

}
