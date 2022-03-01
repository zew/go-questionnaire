package stream

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pbberlin/dbg"
	adminapi "google.golang.org/api/appengine/v1"
	// adminapi "google.golang.org/api/appengine/v1beta5"  // empty responses
)

func getInfos(ctx context.Context, w io.Writer) {

	appsID := os.Getenv("GAE_APPLICATION")
	// appsID = "rent-o-mat"
	if len(appsID) > 2 {
		// chopping of g~ or h~ ...
		tokens := strings.Split(appsID, "~")
		if len(tokens) > 1 {
			appsID = tokens[1]
		}
	}

	aeS, err := adminapi.NewService(ctx) // appengine service
	if err != nil {
		fmt.Fprintf(w, "error obtaining service: %v\n", err)
	}

	{
		/*
			AuthorizedCertificates	*AppsAuthorizedCertificatesService
			AuthorizedDomains		*AppsAuthorizedDomainsService
			DomainMappings			*AppsDomainMappingsService
			Firewall				*AppsFirewallService
			Locations				*AppsLocationsService
			Operations				*AppsOperationsService
			Services				*AppsServicesService
		*/
		apps := aeS.Apps
		call := apps.Get(appsID)
		app, err := call.Do()
		if err != nil {
			fmt.Fprintf(w, "error obtaining app info for %v: %v\n", appsID, err)
		}
		fmt.Fprintf(w, "\napp info:  %v\n\n", dbg.Dump2String(app))
	}

	//
	{
		als := adminapi.NewAppsLocationsService(aeS)
		call := als.List(appsID)
		app, err := call.Do()
		if err != nil {
			fmt.Fprintf(w, "error obtaining apps location service for %v: %v\n", appsID, err)
		}
		fmt.Fprintf(w, "\napps locations:  %v\n\n", dbg.Dump2String(app))
	}

	//
	{
		ais := adminapi.NewAppsServicesVersionsInstancesService(aeS)
		call := ais.List(appsID, "service-xx", "version-xx")
		app, err := call.Do()
		if err != nil {
			fmt.Fprintf(w, "error obtaining apps service versions instances for %v: %v\n", appsID, err)
		}
		fmt.Fprintf(w, "\napps service versions instances:  %v\n\n", dbg.Dump2String(app))
	}

}

/*InstanceInfo queries via app engine admin service API

https://pkg.go.dev/google.golang.org/api@v0.14.0/appengine/v1beta5?tab=doc#APIService

*/
func InstanceInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	env(w)

	ctx := context.Background()
	getInfos(ctx, w)

	// ctx = r.Context()  // getting the context from the request does not change anything
	// getInfos(ctx, w)

}
