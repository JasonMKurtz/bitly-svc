package main

import (
	"encoding/json"
	"net/http"
	"strings"

	bitlinks "./bitlinks"
	metrics "./metrics"
	server "./server"
	userinfo "./userinfo"
)

var app server.Routes

func main() {
	app = server.Routes{Port: "5000"}
	app.Routes = []server.Route{
		server.Route{Route: "/mystats", Handler: StatsEndpoint},
	}

	app.Listen()
}

type StatResponse struct {
	Stats []metrics.MetricResponse
}

func StatsEndpoint(w http.ResponseWriter, req *http.Request, route server.Route) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		panic("Required Authorization header missing.")
	}

	token := strings.Fields(authHeader)[1]

	groupId := userinfo.GetUserinfo(token).Group
	links := bitlinks.GetBitlinks(groupId, token)

	var statsResp []metrics.MetricResponse
	for _, link := range links {
		avg := metrics.GetAverageClickPerCountry(token, link)
		if len(avg.Clicks) < 1 {
			continue
		}
		statsResp = append(statsResp, avg)
	}

	stats := StatResponse{Stats: statsResp}
	resp, _ := json.Marshal(stats)

	w.Write(resp)
}
