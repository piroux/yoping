package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"

	"piroux.dev/yoping/api/pkg/apps/main/control/rest/rest/controllers"
	"piroux.dev/yoping/api/pkg/apps/main/domain/services"
	repo_ping "piroux.dev/yoping/api/pkg/apps/main/persistence/repos/ping"
	repo_user "piroux.dev/yoping/api/pkg/apps/main/persistence/repos/user"
	"piroux.dev/yoping/api/pkg/apps/main/ports/portnotify_megaring"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func Run() {

	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		TimeFieldFormat:  time.RFC3339,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	// Service
	rtr := chi.NewRouter()

	rtr.Use(httplog.RequestLogger(logger))
	rtr.Use(middleware.Heartbeat("/ping"))

	rtr.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			httplog.LogEntrySetField(ctx, "user", slog.StringValue("user1"))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	rtrTest := chi.NewRouter()

	rtrTest.Get("/log/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	rtrTest.Get("/log/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oh no")
	})

	rtrTest.Get("/log/info", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		w.Header().Add("Content-Type", "text/plain")
		oplog.Info("info here")
		w.Write([]byte("info here"))
	})

	rtrTest.Get("/log/warn", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		oplog.Warn("warn here")
		w.WriteHeader(400)
		w.Write([]byte("warn here"))
	})

	rtrTest.Get("/log/err", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		oplog.Error("msg here", "err", errors.New("err here"))
		w.WriteHeader(500)
		w.Write([]byte("oops, err"))
	})

	rtr.Mount("/_/test", rtrTest)

	api := humachi.New(
		rtr,
		huma.DefaultConfig("Yoping API", "0.1.0"),
	)

	// Register GET /greeting/{name}
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	dbURL := os.Getenv("DB_URL")
	if len(dbURL) == 0 {
		log.Fatal("Environment variable DB_URL is empty")
	}

	ctrPing := &controllers.ControllerPing{
		ServicePing: services.NewPingService(
			&portnotify_megaring.NotifierMegaring{},
			&repo_user.RepoDB{},
			repo_ping.NewRepoDB(dbURL),
		),
	}

	// Register GET /greeting/{name}
	huma.Register(api, huma.Operation{
		OperationID: "ping-ex",
		Method:      http.MethodPost,
		//Path:        "/users/:phoneFrom/ping/:phone",
		Path:        "/ping/ex/{phoneFrom}/{phoneTo}",
		Summary:     "Send a Ping",
		Description: "Send a Ping from a Phone Number to another Phone Number",
		Tags:        []string{"Ping"},
	}, ctrPing.PingEx)

	huma.Register(api, huma.Operation{
		OperationID: "ping-in",
		Method:      http.MethodGet,
		//Path:        "/users/:phoneFrom/ping/:phone",
		Path:        "/ping/in/{phoneFrom}/{phoneTo}",
		Summary:     "Receive a Ping",
		Description: "Receive a Ping from a Phone Number to another Phone Number",
		Tags:        []string{"Ping"},
	}, ctrPing.PingIn)

	// Start the server!
	slog.Debug("API Server running ...")
	fmt.Println("API Server running ...")
	http.ListenAndServe("0.0.0.0:8855", rtr)
}
