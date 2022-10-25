package api

import "net/http"

func (a *App) dashboardRoutes() {
	a.HandleFunc("/api/flags/{id}", a.H.GetFlag).Methods("GET")
	a.HandleFunc("/api/flags/{id}/toggle", a.H.ToggleFlag).Methods("PATCH")
	a.HandleFunc("/api/flags/{id}", a.H.UpdateFlag).Methods("PATCH")
	a.HandleFunc("/api/flags/{id}", a.H.DeleteFlag).Methods("DELETE")
	a.HandleFunc("/api/flags", a.H.GetAllFlags).Methods("GET")
	a.HandleFunc("/api/flags", a.H.CreateFlag).Methods("POST")

	a.HandleFunc("/api/audiences/{id}", a.H.UpdateAudience).Methods("PATCH")
	a.HandleFunc("/api/audiences/{id}", a.H.GetAudience).Methods("GET")
	a.HandleFunc("/api/audiences/{id}", a.H.DeleteAudience).Methods("DELETE")
	a.HandleFunc("/api/audiences", a.H.GetAllAudiences).Methods("GET")
	a.HandleFunc("/api/audiences", a.H.CreateAudience).Methods("POST")

	a.HandleFunc("/api/attributes", a.H.GetAllAttributes).Methods("GET")
	a.HandleFunc("/api/attributes/{id}", a.H.GetAttribute).Methods("GET")
	a.HandleFunc("/api/attributes", a.H.CreateAttribute).Methods("POST")
	a.HandleFunc("/api/attributes/{id}", a.H.DeleteAttribute).Methods("DELETE")

	a.HandleFunc("/api/auditlogs", a.H.GetAuditLogs).Methods("GET")

	a.HandleFunc("/api/sdkkeys", a.H.GetSdkKeys).Methods("GET")
	a.HandleFunc("/api/sdkkeys/{id}", a.H.RegenSDKkey).Methods("DELETE")
}

func (a *App) providerRoutes() {
	a.HandleFunc("/flagset", a.H.GetFlagset).Methods("GET")
}

func (a *App) staticRoutes() {
	a.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static/")))).Methods("GET")
	a.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	}).Methods("GET")
}
