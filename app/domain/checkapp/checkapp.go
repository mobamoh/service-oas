package checkapp

import (
	"github.com/mobamoh/service-oas/foundation/logger"
	"net/http"
	"os"
	"runtime"
)

type app struct {
	build string
	log   *logger.Logger
	//db    *sqlx.DB
}

func newApp(build string, log *logger.Logger) *app {
	return &app{
		build: build,
		log:   log,
		//db:    db,
	}
}

// readiness checks if the database is ready and if not will return a 500 status.
// Do not respond by just returning an error because further up in the call
// stack it will interpret that as a non-trusted error.
func (a *app) readiness(res http.ResponseWriter, req *http.Request) {
	//ctx, cancel := context.WithTimeout(req.Context(), time.Second)
	//defer cancel()

	//if err := sqldb.StatusCheck(ctx, a.db); err != nil {
	//	a.log.Info(ctx, "readiness failure", "ERROR", err)
	//	return errs.New(errs.Internal, err)
	//}
	res.Write([]byte("Ready"))
	//return nil
}

// liveness returns simple status info if the service is alive. If the
// app is deployed to a Kubernetes cluster, it will also return pod, node, and
// namespace details via the Downward API. The Kubernetes environment variables
// need to be set within your Pod/Deployment manifest.
func (a *app) liveness(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	info := Info{
		Status:     "up",
		Build:      a.build,
		Host:       host,
		Name:       os.Getenv("KUBERNETES_NAME"),
		PodIP:      os.Getenv("KUBERNETES_POD_IP"),
		Node:       os.Getenv("KUBERNETES_NODE_NAME"),
		Namespace:  os.Getenv("KUBERNETES_NAMESPACE"),
		GOMAXPROCS: runtime.GOMAXPROCS(0),
	}

	data, contentType, err := info.Encode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.log.Error(r.Context(), "encode error", "ERROR", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", contentType)

	if _, err := w.Write(data); err != nil {
		a.log.Error(r.Context(), "write error", "ERROR", err)
	}
}
