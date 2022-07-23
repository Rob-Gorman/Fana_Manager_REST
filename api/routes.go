package api

func (s *Server) dashboardRoutes() {
	s.HandleFunc("/api/flags/{id}", s.H.GetFlag).Methods("GET")
	s.HandleFunc("/api/flags/{id}/toggle", s.H.ToggleFlag).Methods("PATCH")
	s.HandleFunc("/api/flags/{id}", s.H.UpdateFlag).Methods("PATCH")
	s.HandleFunc("/api/flags/{id}", s.H.DeleteFlag).Methods("DELETE")
	s.HandleFunc("/api/flags", s.H.GetAllFlags).Methods("GET")
	s.HandleFunc("/api/flags", s.H.CreateFlag).Methods("POST")

	s.HandleFunc("/api/audiences/{id}", s.H.UpdateAudience).Methods("PATCH")
	s.HandleFunc("/api/audiences/{id}", s.H.GetAudience).Methods("GET")
	s.HandleFunc("/api/audiences/{id}", s.H.DeleteAudience).Methods("DELETE")
	s.HandleFunc("/api/audiences", s.H.GetAllAudiences).Methods("GET")
	s.HandleFunc("/api/audiences", s.H.CreateAudience).Methods("POST")

	s.HandleFunc("/api/attributes", s.H.GetAllAttributes).Methods("GET")
	s.HandleFunc("/api/attributes/{id}", s.H.GetAttribute).Methods("GET")
	s.HandleFunc("/api/attributes", s.H.CreateAttribute).Methods("POST")
	s.HandleFunc("/api/attributes/{id}", s.H.DeleteAttribute).Methods("DELETE")

	s.HandleFunc("/api/auditlogs", s.H.GetAuditLogs).Methods("GET")

	s.HandleFunc("/api/sdkkeys", s.H.GetSdkKeys).Methods("GET")
	s.HandleFunc("/api/sdkkeys/{futurework}", s.H.RegenSDKkey).Methods("POST")
}

func (s *Server) providerRoutes() {
	s.HandleFunc("/flagset", s.H.GetFlagset).Methods("GET")
}
