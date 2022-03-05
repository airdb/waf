package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"

	"github.com/airdb/sailor/deployutil"
	"github.com/airdb/sailor/faas"
	"github.com/airdb/sailor/version"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/ipipdotnet/ipdb-go"
)

var idb = &ipdb.City{}

func main() {
	var err error

	idb, err = ipdb.NewCity("ipv4_en.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeHTML))

	r.Get("/", faas.HandleVersion)
	r.Get("/ip/", faas.HandleVersion)
	r.Get("/ip/info", HandleRemoteIP)

	project := "waf"
	path := filepath.Join("/", project, "dl", "/")

	fs := http.FileServer(http.Dir("dl"))
	r.Handle(path+"*", http.StripPrefix(path, fs))
	r.Handle("/dl/*", http.StripPrefix("/tmp/", fs))

	fmt.Println("hello", deployutil.GetDeployStage())

	faas.RunTencentChiWithSwagger(r, "waf")
}

const IPIPEN = "EN"

// RootHandler - Returns all the available APIs
// @Summary Root handler.
// @Description Tells if the root APIs are working or not.
// @Tags root
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Header 200 {string} Token "jwt"
// @Failure 400,404 {string} response "4xx"
// @Failure 500 {string} response "500"
// @Failure default {string} response "default"
// @Router / [get]
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("return version info"))
	fmt.Println(version.GetBuildInfo())
	// ret := map[string]interface{}{"feed": "11", "success": true}

	// version.Init()
	// ret := version.GetBuildInfo()
	// render.JSON(w, r, ret)
	render.JSON(w, r, version.GetBuildInfo())
	w.WriteHeader(http.StatusOK)
}

// HandleItmeCreate - create item.
// @Summary Create item.
// @Description Create item.
// @Tags item
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router /chi/item/create [post]
func HandleItemCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create item"))
	w.WriteHeader(http.StatusOK)
}

// HandleItmeQuery - query item.
// @Summary Query item.
// @Description Query item api by id or name.
// @Tags item
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router /ip/query [get]
func HandleItemQuery(w http.ResponseWriter, r *http.Request) {
	db, err := ipdb.NewCity("ipv4_en.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	ip := r.URL.Query().Get("ip")

	// ip := "129.226.148.218"
	dbmap, _ := db.FindMap(ip, IPIPEN)
	dbmap["ip"] = ip
	mJson, _ := json.Marshal(dbmap)

	// msg = dbmap["idc"]
	// w.Write([]byte("welcome hello"))
	// w.Write([]byte(msg))
	w.Write(mJson)
	w.WriteHeader(http.StatusOK)
}

// HandleItmeQuery - query item.
// @Summary Query item.
// @Description Query item api by id or name.
// @Tags item
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router /ip/{ip} [get]
func HandleQueryIP(w http.ResponseWriter, r *http.Request) {
	/*
		db, err := ipdb.NewCity("ipv4_en.ipdb")
		if err != nil {
			log.Fatal(err)
		}
	*/

	ip := chi.URLParam(r, "ip")

	if ip == "" {
		cip, _, _ := net.SplitHostPort(r.RemoteAddr)
		ip = cip
	}

	dbmap, _ := idb.FindMap(ip, IPIPEN)
	dbmap["ip"] = ip
	mJson, _ := json.Marshal(dbmap)

	// msg = dbmap["idc"]
	// w.Write([]byte("welcome hello"))
	// w.Write([]byte(msg))
	w.Write(mJson)
	w.WriteHeader(http.StatusOK)
}

// HandleItmeQuery - query item.
// @Summary Query item.
// @Description Query item api by id or name.
// @Tags item
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router /ip [get]
func HandleRemoteIP(w http.ResponseWriter, r *http.Request) {
	/*
		db, err := ipdb.NewCity("ipv4_en.ipdb")
		if err != nil {
			log.Fatal(err)
		}
	*/

	//fmt.Println(r.)

	cip, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := cip
	fmt.Println(ip)
	//w.Write([]byte(ip))

	dbmap, _ := idb.FindMap(ip, IPIPEN)
	dbmap["ip"] = ip
	mJson, _ := json.Marshal(dbmap)

	// msg = dbmap["idc"]
	// w.Write([]byte("welcome hello"))
	// w.Write([]byte(msg))
	w.Write(mJson)
	w.WriteHeader(http.StatusOK)
}
