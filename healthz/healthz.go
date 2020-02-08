package healthz

import (
	"encoding/json"
	"github.com/daniel-cole/fileserver/config"
	"github.com/daniel-cole/fileserver/middleware"
	"net/http"
)

var Version string

type healthz struct {
	Health  bool                    `json:"health"`
	Version string                  `json:"version"`
	Message string                  `json:"message"`
	Config  config.FileServerConfig `json:"config"`
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	message := "all is well and good with the world"

	healthz := &healthz{
		Health:  true,
		Version: Version,
		Message: message,
		Config:  *config.FileServer,
	}

	health, err := json.Marshal(healthz)
	if err != nil {
		middleware.LogWithContext(ctx).Errorf("/healthz failed to respond: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(health)
}
