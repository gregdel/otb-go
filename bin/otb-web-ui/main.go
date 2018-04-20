package main

import (
	"log"
	"net/http"

	"github.com/gregdel/otb-go/lib/openwrt/ubus"
	"github.com/gregdel/otb-go/lib/openwrt/uci"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// TODO: remove
var rend = render.New()

// GetConfigHandler displays the configuration
func GetConfigHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	configName := ps.ByName("config")

	configuration, err := ubus.UciGetConfig(uci.ConfigType(configName))
	if err != nil {
		rend.JSON(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}

	rend.JSON(w, http.StatusOK, configuration)
}

func main() {
	router := httprouter.New()
	router.GET("/configs/:config", GetConfigHandler)

	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
