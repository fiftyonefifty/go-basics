package main

import (
	"context"
	"flag"
	"fmt"
	"hello-cdk/utils"
	"log"
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"go.opencensus.io/trace"
	"gocloud.dev/blob"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
)

type cliFlags struct {
	bucket          string
	dbHost          string
	dbName          string
	dbUser          string
	dbPassword      string
	motdVar         string
	motdVarWaitTime time.Duration

	// GCP only.
	cloudSQLRegion    string
	runtimeConfigName string
}

var envFlag string

func main() {
	cf := new(cliFlags)
	flag.StringVar(&envFlag, "env", "local", "environment to run under (gcp, aws, azure, or local)")
	addr := flag.String("listen", ":8080", "port to listen for HTTP on")
	flag.StringVar(&cf.bucket, "bucket", "", "bucket name")
	flag.StringVar(&cf.dbHost, "db_host", "", "database host or Cloud SQL instance name")
	flag.StringVar(&cf.dbName, "db_name", "guestbook", "database name")
	flag.StringVar(&cf.dbUser, "db_user", "guestbook", "database user")
	flag.StringVar(&cf.dbPassword, "db_password", "", "database user password")
	flag.StringVar(&cf.motdVar, "motd_var", "", "message of the day variable location")
	flag.DurationVar(&cf.motdVarWaitTime, "motd_var_wait_time", 5*time.Second, "polling frequency of message of the day")
	flag.StringVar(&cf.cloudSQLRegion, "cloud_sql_region", "", "region of the Cloud SQL instance (GCP only)")
	flag.StringVar(&cf.runtimeConfigName, "runtime_config", "", "Runtime Configurator config resource (GCP only)")
	flag.Parse()

	fmt.Println(*addr)

	ctx := context.Background()
	var srv *server.Server
	var cleanup func()
	var err error
	switch envFlag {
	case "local":
		// The default MySQL instance is running on localhost
		// with this root password.
		if cf.dbHost == "" {
			cf.dbHost = "localhost"
		}
		if cf.dbPassword == "" {
			cf.dbPassword = "xyzzy"
		}
		srv, cleanup, err = setupLocal(ctx, cf)
	default:
		log.Fatalf("unknown -env=%s", envFlag)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// Listen and serve HTTP.
	log.Printf("Running, connected to %q cloud", envFlag)
	log.Fatal(srv.ListenAndServe(*addr))
}

// applicationSet is the Wire provider set for the Guestbook application that
// does not depend on the underlying platform.
var applicationSet = wire.NewSet(
	newApplication,
	appHealthChecks,
	trace.AlwaysSample,
	newRouter,
	wire.Bind(new(http.Handler), new(*mux.Router)),
)

// application is the main server struct for Guestbook. It contains the state of
// the most recently read message of the day.
type application struct {
	bucket  *blob.Bucket
	motdVar *runtimevar.Variable
}

// newApplication creates a new application struct based on the backends and the message
// of the day variable.
func newApplication(bucket *blob.Bucket, motdVar *runtimevar.Variable) *application {
	return &application{
		bucket:  bucket,
		motdVar: motdVar,
	}
}
func newRouter(app *application) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", app.index)
	return r
}
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, utils.Message(true, "hello world!"))
}

// appHealthChecks returns a health check for the database. This will signal
// to Kubernetes or other orchestrators that the server should not receive
// traffic until the server is able to connect to its database.
func appHealthChecks() ([]health.Checker, func()) {

	list := []health.Checker{}
	return list, func() {

	}
}
