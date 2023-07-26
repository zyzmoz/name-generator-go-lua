package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	lua "github.com/yuin/gopher-lua"
)

type ResponseJSON struct {
	Name string
}

func main() {
	serverPort := 8080

	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			fmt.Printf("Default port")
		} else {
			serverPort = port
			fmt.Printf("Default port %d:", port)

		}
	}
	// server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("server %s /\n", r.Method)

		L := lua.NewState()
		defer L.Close()
		if err := L.DoFile("get_file_name.lua"); err != nil {
			panic(err)
		}
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("get_name"),
			NRet:    1,
			Protect: true,
		}, lua.LString("")); err != nil {
			panic(err)
		}
		ret := L.Get(-1) // returned value
		L.Pop(1)         // remove received value
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseJSON{
			Name: ret.String(),
		})
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error running server: \n", err)
		}
	}

	time.Sleep(100 * time.Millisecond)
}
