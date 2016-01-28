package rest

func (h *handler) handleRoot() error {
	response := map[string]interface{}{
		"status": "running",
	}

	return h.writeJSON(response)
}

func (h *handler) handleAuth() error {
	return nil
}
