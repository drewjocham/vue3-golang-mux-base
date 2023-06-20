package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"fullstackguru/internal/config"
	"fullstackguru/internal/courses"
	"fullstackguru/pkg"
	_ "github.com/lib/pq"
	"strings"
	"sync"

	clogger "fullstackguru/pkg/logger"
)

type application struct {
	config     *config.Config
	logger     *clogger.Logger
	courses    *courses.Courses
	tokenRepo  *pkg.TokenRepository
	middleware pkg.MiddleWare
	helper     pkg.Helper
	db         *sql.DB
	wg         sync.WaitGroup
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clog := clogger.GetLoggerFromContext(ctx)

	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	clog.Info("Starting Application")

	flag.IntVar(&cfg.Port, "port", cfg.Port, "API server port")
	flag.StringVar(&cfg.Env, "env", cfg.Env, "Environment (development|staging|production)")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.Cors.TrustedOrigins = strings.Fields(val)

		return nil
	})

	// postgres database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DatabaseConfig.Host, cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.DataName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		clog.ErrorCtx(err, clogger.Ctx{
			"msg": "Unable to connect to the database",
		})
	}

	// Courses
	courseRepo := courses.NewCourseRepository(db)
	courseService := courses.NewCoursesService(courseRepo)

	//Tokens
	tokenRepo := pkg.NewTokenRepository(db)

	app := &application{
		config:    cfg,
		logger:    clog,
		db:        db,
		courses:   courseService,
		tokenRepo: tokenRepo,
	}

	err = app.serve()
	if err != nil {
		clog.Error(err)
	}

}
