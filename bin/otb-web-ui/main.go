package main

import (
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gregdel/otb-go/lib/openwrt/ubus"
	"github.com/unrolled/render"
)

// Use packr to statically add the templates in the go binary
//go:generate packr

func main() {
	templateBox := packr.NewBox("./templates")
	r := render.New(render.Options{
		Layout: "layout",
		Asset: func(name string) ([]byte, error) {
			n := strings.Replace(name, "templates/", "", -1)
			return templateBox.MustBytes(n)
		},
		AssetNames: func() []string {
			return []string{
				"templates/layout.tmpl",
				"templates/error.tmpl",
				"templates/system-info.tmpl",
				"templates/network-info.tmpl",
			}
		},
	})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		info, err := ubus.SystemBoard()
		if err != nil {
			r.HTML(w, http.StatusOK, "error", err.Error())
			return
		}

		r.HTML(w, http.StatusOK, "system-info", info)
	})

	mux.HandleFunc("/network", func(w http.ResponseWriter, req *http.Request) {
		info, err := ubus.NetworkDump()
		if err != nil {
			r.HTML(w, http.StatusOK, "error", err.Error())
			return
		}

		r.HTML(w, http.StatusOK, "network-info", info)
	})

	http.ListenAndServe("0.0.0.0:8080", mux)
}
