package global

import "github.com/ucarion/urlpath"

// AllowedRoutes are the paths that sidecar-proxy application supports
var AllowedRoutes = []urlpath.Path{
	urlpath.New("/company"),
	urlpath.New("/company/:id"),
	urlpath.New("/company/account"),
	urlpath.New("/account"),
	urlpath.New("/account/:id"),
	urlpath.New("/:id"),
	urlpath.New("/account/:id/user"),
	urlpath.New("/tenant/account/blocked"),
}
