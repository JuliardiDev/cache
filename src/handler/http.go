package src

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mod "github.com/cache/src/module"
)

// Write to write person on disk and memory
func Write(w http.ResponseWriter, r *http.Request) {
	p := mod.Person{
		Name:    r.FormValue("name"),
		HP:      r.FormValue("hp"),
		Address: r.FormValue("address"),
	}

	err := mod.WriteFile(p)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(map[string]interface{}{
		"success": true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("kirim handler write")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Read to read person on disk and memory
func Read(w http.ResponseWriter, r *http.Request) {

	// Sanitize form input
	req := func(rArg *http.Request) mod.PersonReq {
		sid := rArg.FormValue("sid")
		sname := rArg.FormValue("sname")
		saddress := rArg.FormValue("saddress")
		shp := rArg.FormValue("shp")
		page, _ := strconv.Atoi(rArg.FormValue("page"))

		s := func(s string) int {
			switch s {
			case "1":
				return 1
			case "2":
				return 2
			default:
				return 0
			}
		}

		return mod.PersonReq{
			Sort: mod.Rsort{
				ID:      s(sid),
				Name:    s(sname),
				Address: s(saddress),
				HP:      s(shp),
			},
			Page: page,
		}
	}

	data, err := mod.ReadFiles(req(r))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(map[string]interface{}{
		"data": data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
