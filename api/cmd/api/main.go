package main

import (
	"context"
	"expvar"
	"flag"
	"github.com/interviews/internal/config"
	"github.com/interviews/internal/data"
	"github.com/interviews/internal/jsonlog"
	"github.com/interviews/internal/vcs"
	log "github.com/interviews/utils/logger"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	version = vcs.Version()
)

type application struct {
	config *config.Config
	logger *jsonlog.Logger
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clog := log.GetLoggerFromContext(ctx)

	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	clog.Info("Starting Application")

	// not a huge fan of this...
	port, err := strconv.Atoi(strconv.Itoa(cfg.Port))
	if err != nil {
		return
	}

	flag.IntVar(&cfg.Port, "port", port, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	expvar.NewString("version").Set(version)

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.Cors.TrustedOrigins = strings.Fields(val)

		return nil
	})

	app := &application{
		config: cfg,
		logger: logger,
		//models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

}
