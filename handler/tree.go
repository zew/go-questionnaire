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
func Tree(lc string) *TreeT {

	root := &TreeT{

		Node: Info{Title: "Root Node", Keys: []string{"root"}},
		Children: []TreeT{
			// {Node: Info{Title: "Language", Keys: []string{"language"}},
			// 	Children: []TreeT{
			// 		{Node: Info{Title: "Deutsch"}},
			// 		{Node: Info{Title: "English"}},
			// 	},
			// },
			{Node: Info{
				Title: cfg.Get().Mp["user"].Tr(lc),
				Keys:  []string{"loginlogout"},
			},
				Children: []TreeT{
					{Node: infos.ByKey("login-primitive")},
					{Node: infos.ByKey("logout")},
					{Node: infos.ByKey("change-password-primitive")},
					{Node: infos.ByKey("create-anonymous-id")},
				},
			},
			{Node: Info{
				Title: "&nbsp;" + cfg.Get().Mp["about"].Tr(lc) + "&nbsp;&nbsp",
				Keys:  []string{"about"},
			},
				Children: []TreeT{
					{Node: infos.ByKeyTranslated("imprint", lc)},
				},
			},
			{Node: Info{
				Title: "Sys admin",
				Keys:  []string{"admin"},
				Allow: map[Privilege]bool{Admin: true},
			},
				Children: []TreeT{
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

	return root

}

// ByKey recursively retrieves a node;
// use ByKey() and SetByKey() to modify the nav tree
func (tr *TreeT) ByKey(key string) *TreeT {
	if tr.Node.HasKey(key) {
		return tr
	}
	for _, c := range tr.Children {
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
func (tr *TreeT) SetByKey(key string, repl *Info) bool {
	if tr.Node.HasKey(key) {
		tr.Node = *repl
		return true
	}
	for idx := range tr.Children {
		if tr.Children[idx].SetByKey(key, repl) { // we have to access via slice - to really change the underlying instance
			return true
		}
	}
	return false
}

// AppendAfterByKey recursively appends behind an existing node
// asChild == false => append as sibling - same level
// asChild == true  => append as child   - one  level deeper
func (tr *TreeT) AppendAfterByKey(key string, summand *Info, asChild ...bool) bool {

	asCh := false
	if len(asChild) > 0 {
		asCh = asChild[0]
	}

	//
	//
	//		tr.Node                    - new child to root
	// 			or
	//		tr.Children                - new sibling to child
	// 			or
	// 		tr.Children[idx].Children  - new child to child
	if key == "root" && tr.Node.HasKey(key) {
		// level 0 - root
		// special case - should never get siblings - only children
		// log.Printf("appending node %-12q - %-12v as child   to root", summand.Title, summand.Keys[0])
		ln := len(tr.Children)
		cp := make([]TreeT, 0, ln+1)
		cp = append(cp, TreeT{Node: *summand}) // first slot
		cp = append(cp, tr.Children...)        // previous children
		tr.Children = cp
		return true
	}
	for idx := range tr.Children {
		if tr.Children[idx].Node.HasKey(key) {
			if !asCh {
				// log.Printf("appending node %-12q - %-12v as sibling to child", summand.Title, summand.Keys[0])
				ln := len(tr.Children)
				cp := make([]TreeT, 0, ln+1)
				cp = append(cp, tr.Children[0:idx+1]...)
				cp = append(cp, TreeT{Node: *summand})
				cp = append(cp, tr.Children[idx+1:ln]...)
				tr.Children = cp
			} else {
				// log.Printf("appending node %-12q - %-12v as child   to child", summand.Title, summand.Keys[0])
				ln := len(tr.Children[idx].Children)
				cp := make([]TreeT, 0, ln+1)
				cp = append(cp, TreeT{Node: *summand})        // first slot
				cp = append(cp, tr.Children[idx].Children...) // previous children
				tr.Children[idx].Children = cp
			}
			return true

		}
		found := tr.Children[idx].AppendAfterByKey(key, summand, asCh) // this is no oversight; return only on true
		if found {
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
func (tr *TreeT) NavHTML(w io.Writer, r *http.Request, isLogin, isAdmin bool, lvl int) {

	// the root node itself is not rendered - we go straight to the children
	if lvl == 0 {
		for _, child := range tr.Children {
			child.NavHTML(w, r, isLogin, isAdmin, lvl+1)
		}
		return
	}

	if lvl == 1 {
		fmt.Fprint(w, "\n") // pretty indentations in HTML source
	}

	preventClick := " onclick='return false;' " // nav items without URL should not be clickable;

	needsLogout := tr.Node.Allow[LoggedOut]
	if needsLogout && isLogin {
		return
	}
	needsLogin := tr.Node.Allow[LoggedIn]
	if needsLogin && !isLogin {
		return
	}
	needsAdmin := tr.Node.Allow[Admin]
	if needsAdmin && !isAdmin {
		return
	}

	// if tr.Node.InNav {
	{

		htmlIndent := strings.Repeat(" ", 10*lvl) // just for readability in HTML source
		key := ""
		navURL := ""
		activeClass := "" // style the active nav item
		onClick := ""
		accessKey := ""

		if len(tr.Node.Keys) > 0 {
			key = tr.Node.Keys[0]
		}

		if len(tr.Node.Urls) > 0 {
			navURL = cfg.Pref(tr.Node.Urls[0])
		}

		/*
			regular case for determining is-active

			notice, that
			Node.Urls dont have suffix '/'
			except home page, which is '/'
		*/
		if navURL == strings.TrimSuffix(r.URL.Path, "/") || tr.Node.Active {
			navURL = ""
			activeClass = " is-active "
		}

		/*
			special case for determining is-active

			some parent nodes have *no* Node.Urls,
			and can never be active

			nodesNoURLs := []string{
				"loginlougout",
				"about",
				"admin",
				"language",
				"quest-pages",
			}
		*/
		if len(tr.Node.Urls) == 0 {
			activeClass = ""
		}

		if navURL == "" {
			onClick = preventClick
		}

		if tr.Node.OnClick != "" {
			onClick = tr.Node.OnClick
		}

		if tr.Node.ShortCut != "" {
			accessKey = fmt.Sprintf(" accesskey='%v' ", tr.Node.ShortCut)
			tr.Node.Title += fmt.Sprintf(" <span title='Keyboard shortcut SHIFT+ALT+%v' >(%v)</span>", tr.Node.ShortCut, tr.Node.ShortCut)
		}

		if len(tr.Children) == 0 {

			fmt.Fprintf(w, "%v<li id='%v' >\n", htmlIndent, key)
			fmt.Fprintf(w, "%v  <a href='%v' class='%v'  %v  %v  >%v</a>\n",
				htmlIndent, navURL, activeClass, onClick, accessKey, tr.Node.Title,
			)
			fmt.Fprintf(w, "%v</li>\n", htmlIndent)

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

			fmt.Fprintf(w, "%v<li  class='nde-2nd-lvl' id='%v' >\n", htmlIndent, key)

			// same as above - without enclosing <li>
			fmt.Fprintf(w, "%v  <a href='%v' class='%v'  %v  %v  >%v</a>  \n",
				htmlIndent, navURL, activeClass, onClick, accessKey, tr.Node.Title,
			)

			fmt.Fprintf(w, "%v     <ul class='mnu-3rd-lvl'>\n", htmlIndent)
			for _, child := range tr.Children {
				child.NavHTML(w, r, isLogin, isAdmin, lvl+1)
			}
			fmt.Fprintf(w, "%v     </ul>\n", htmlIndent)

			fmt.Fprintf(w, "%v</li>\n", htmlIndent)

		}

	}

}
