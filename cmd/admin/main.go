package main

import httputil "multiapi/pkg/http"

func main() {
	conf := NewConfig()
	adminAPI := NewAdminAPI(conf)
	httputil.StartServer(adminAPI.server)
}
