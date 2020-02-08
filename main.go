package main

import (
	"context"
	"flag"
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"github.com/daniel-cole/fileserver/config"
	"github.com/daniel-cole/fileserver/healthz"
	"github.com/daniel-cole/fileserver/middleware"
	"github.com/daniel-cole/fileserver/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Version = "undefined"

func main() {
	var port int
	var address string
	var sourceRanges string
	var directory string
	var htpasswdFile string
	var logLevel string

	flag.IntVar(&port, "port", 9000, "port to listen on")
	flag.StringVar(&address, "address", "0.0.0.0", "address to bind to")
	flag.StringVar(&sourceRanges, "sourceRanges", "0.0.0.0/0,::/0", "source addresses to allow connectivity from")
	flag.StringVar(&directory, "directory", ".", "directory to serve files from")
	flag.StringVar(&htpasswdFile, "htpasswdFile", "htpasswd", "htpasswd file to use for authenticating users")
	flag.StringVar(&logLevel, "logLevel", "INFO", "set the log level [INFO|WARN|ERROR|DEBUG]")
	flag.Parse()

	parsedSourceRanges, err := util.ParseSourceRanges(sourceRanges)
	if err != nil {
		logrus.Fatalf("Unable to parse source ranges provided. %v", err)
	}

	_, err = os.Stat(htpasswdFile)
	if err != nil {
		logrus.Fatalf("Unable to find htpasswd file '%s' %v", htpasswdFile, err)
	}

	setLogLevel(logLevel)

	listenAddress := fmt.Sprintf("%s:%d", address, port)

	config.FileServer = &config.FileServerConfig{
		Port:         port,
		Address:      address,
		SourceRanges: parsedSourceRanges,
		Directory:    directory,
		HTPasswdFile: htpasswdFile,
	}

	logrus.Infof("Fileserver port: %d", port)
	logrus.Infof("Fileserver address: %s", address)
	for _, sourceRange := range parsedSourceRanges {
		logrus.Infof("Fileserver source range: %s", sourceRange)
	}
	logrus.Infof("Fileserver directory: %s", directory)
	logrus.Infof("Fileserver htpasswd file: %s", htpasswdFile)

	healthz.Version = Version
	healthzHandler := http.HandlerFunc(healthz.HealthzHandler)
	http.Handle("/healthz", middleware.Logger(healthzHandler))

	// ensure that any requests to the fileserver are authenticated
	secret := auth.HtpasswdFileProvider(htpasswdFile)
	authenticator := auth.NewBasicAuthenticator("", secret)

	http.Handle("/", middleware.Logger(middleware.CheckSourceIP(
		authenticator.Wrap(func(res http.ResponseWriter, req *auth.AuthenticatedRequest) {
			middleware.LogWithContext(req.Context()).Infof("Processing request for user: %s", req.Username)
			http.FileServer(http.Dir(directory)).ServeHTTP(res, &req.Request)
		}))))

	server := http.Server{
		Addr:         listenAddress,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		graceTime := 60 * time.Second
		logrus.Infof("Server is shutting down... grace period: %s", graceTime.String())

		ctx, cancel := context.WithTimeout(context.Background(), graceTime)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logrus.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logrus.Infof("File server is ready to handle requests at %s", listenAddress)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Could not listen on %s: %v\n", listenAddress, err)
	}

	<-done

	logrus.Info("Server stopped")

}

func setLogLevel(logLevel string) {

	switch logLevel {
	case "": // default to INFO level logging
		logrus.SetLevel(logrus.InfoLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	}
}
