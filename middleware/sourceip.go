package middleware

import (
	"fmt"
	"github.com/daniel-cole/fileserver/config"
	"github.com/daniel-cole/fileserver/util"
	"net"
	"net/http"
)

// CheckSourceIP checks if the IP address that the client is using
func CheckSourceIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		clientIP := util.ClientIPFromRemoteAddr(r.RemoteAddr)
		LogWithContext(ctx).Debugf("Checking if %s is permitted to access resources", clientIP)
		if allowedIP := util.CheckIPInNetwork(net.ParseIP(clientIP), config.FileServer.SourceRanges); !allowedIP {
			LogWithContext(ctx).Info("Attempted to access fileserver on IP address not in provided source range")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(fmt.Sprint("Forbidden. Go away.")))
			return
		}

		LogWithContext(ctx).Debugf("IP Address is permitted: %s", clientIP)

		next.ServeHTTP(w, r)

	})
}
