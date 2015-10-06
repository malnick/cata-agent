package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	//Alert       string
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
	Route{
		"HostDisk",
		"GET",
		"/host/disk",
		routeHostDisk,
	},
	Route{
		"HostNetio",
		"GET",
		"/host/netio",
		routeHostNetio,
	},
	Route{
		"HostNetcon",
		"GET",
		"/host/netcon",
		routeHostNetcon,
	},
	Route{
		"HostDocker",
		"GET",
		"/host/docker",
		routeHostDocker,
	},
}
