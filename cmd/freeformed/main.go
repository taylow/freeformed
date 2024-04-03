package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/taylow/freeformed/config"
	"github.com/taylow/freeformed/data"
	"github.com/taylow/freeformed/errors"
	"github.com/taylow/freeformed/form"
	"github.com/taylow/freeformed/random"
	"github.com/taylow/freeformed/storage"
)

var (
	logger    *slog.Logger
	logLogger *log.Logger

	Port     = 8080
	LogLevel = slog.LevelDebug
	DSN      = "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable"

	FormID = random.GenerateRandomAlphaString(8)
)

// init initialises the program on start
func init() {

	var logLevel string
	flag.IntVar(&Port, "port", Port, "define the port to run the server on")
	flag.StringVar(&logLevel, "log-level", logLevel, "define the log level (debug, info, warn, error)")
	flag.StringVar(&DSN, "dsn", DSN, "define the data source name for the database")
	flag.StringVar(&FormID, "form-id", FormID, "define a form ID to use for the server")
	flag.Parse()

	if logLevel != "" {
		if err := LogLevel.UnmarshalText([]byte(logLevel)); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	logger, logLogger = initLogger()
}

// main is the entrypoint to the program
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	log.Println("bye bye 👋")
}

// run runs the program
func run() error {
	ctx := context.Background()

	logger.Debug("setting up interrupts")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	logger.Debug("creating local data repository")
	dataConfig := data.NewLocalDataConfig(
		data.WithDataFilePath("formdata.json"),
		data.WithFormatJSON(true),
	)
	dataRepository, err := data.NewLocalDataRepository(dataConfig)
	if err != nil {
		return err
	}
	defer dataRepository.Close()

	// logger.Debug("creating local file repository")
	// storageConfig := storage.NewLocalFileConfig(storage.WithRootPath("uploaded"))
	// fileRepository, err := storage.NewLocalFileRepository(storageConfig)
	logger.Debug("creating blob file repository")
	storageConfig := config.NewBlobFileConfig(
		config.WithBucket("files"),
		// config.WithForceCreateBucket(true),
		config.WithStaticCredentials("64j4hm6cdvf5ri1EQu7A", "1affBiWOiquxdr7KUFhlfexfkumqW1qhULOSsGii"),
	)
	fileRepository, err := storage.NewBlobFileRepository(ctx, storageConfig)
	if err != nil {
		return err
	}
	defer fileRepository.Close()

	logger.Debug("creating form handler")
	handlerConfig := form.NewProcessorConfig(
		form.WithMaxMemory(32<<20),
		form.WithMaxFiles(10),
		form.WithStaticFormID(FormID),
	)
	handler := form.NewHandler(handlerConfig, nil, dataRepository, fileRepository)

	logger.Debug("setting up server")
	router := http.NewServeMux()
	router.Handle("/", errors.Handle404())
	router.HandleFunc("/submit", errors.HandleWithError(handler.HandleForm))
	router.HandleFunc("/submit/{id}", errors.HandleWithError(handler.HandleForm))
	// router.Handle("/", http.FileServer(http.Dir("web")))

	addr := fmt.Sprintf(":%d", Port)
	server := http.Server{
		Handler:  router,
		Addr:     addr,
		ErrorLog: logLogger,
	}

	logger.Info("starting server", "addr", addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server encountered an error while listening and serving", "error", err)
		}
	}()
	fmt.Println("\n\nform available at http://localhost:8080/submit/" + FormID)

	<-sigs
	logger.Info("shutting down server")

	if err := server.Shutdown(context.Background()); err != nil {
		return errors.Wrap(err, "unable to shutdown server")
	}

	return nil
}

// initLogger initialises the logger
func initLogger() (*slog.Logger, *log.Logger) {
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: LogLevel,
	})
	logger := slog.New(handler)
	// logLogger := slog.NewLogLogger(handler, LogLevel)
	slog.SetDefault(logger)

	return logger, slog.NewLogLogger(handler, LogLevel)
}
