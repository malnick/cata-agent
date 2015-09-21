package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"HostIndex",
		"GET",
		"/host",
		indexHost,
	},
	Route{
		"HostHostname",
		"GET",
		"/host/hostname",
		routeHostname,
	},
	Route{
		"HostMemory",
		"GET",
		"/host/memory",
		routeHostMemory,
	},
	Route{
		"HostCpu",
		"GET",
		"/host/cpu",
		routeHostCpu,
	},
	Route{
		"HostLoad",
		"GET",
		"/host/load",
		routeHostLoad,
	},
}
