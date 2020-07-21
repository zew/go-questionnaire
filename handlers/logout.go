package handlers

import (
	"fmt"
	"net/http"

	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/tpl"
)

// LogoutSiteH calls lgn.Logout and wraps it in site layout
func LogoutSiteH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	body := ""
	err := lgn.LogoutH(w, r)
	if err != nil {
		body = fmt.Sprintf("Error logging out: %v", err)
	} else {
		body = "Logged out"
	}

	mp := map[string]interface{}{
		"HTMLTitle":      "Logout",
		"FormMainAction": "", // self
		"Content":        string(body),
	}

	tpl.Exec(w, r, mp, "main-desktop.html")
}
