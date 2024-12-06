package checkapp

import "github.com/go-faster/jx"

// Info represents information about the service.
type Info struct {
	Status     string `json:"status,omitempty"`
	Build      string `json:"build,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	PodIP      string `json:"podIP,omitempty"`
	Node       string `json:"node,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
}

// Encode implements the encoder interface.
func (app Info) Encode() ([]byte, string, error) {
	// Create a new encoder
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	// Start encoding the object
	e.ObjStart()

	// Encode each field if it's not empty
	if app.Status != "" {
		e.FieldStart("status")
		e.Str(app.Status)
	}
	if app.Build != "" {
		e.FieldStart("build")
		e.Str(app.Build)
	}
	if app.Host != "" {
		e.FieldStart("host")
		e.Str(app.Host)
	}
	if app.Name != "" {
		e.FieldStart("name")
		e.Str(app.Name)
	}
	if app.PodIP != "" {
		e.FieldStart("podIP")
		e.Str(app.PodIP)
	}
	if app.Node != "" {
		e.FieldStart("node")
		e.Str(app.Node)
	}
	if app.Namespace != "" {
		e.FieldStart("namespace")
		e.Str(app.Namespace)
	}
	if app.GOMAXPROCS != 0 {
		e.FieldStart("GOMAXPROCS")
		e.Int(app.GOMAXPROCS)
	}

	// End the object
	e.ObjEnd()

	// Get the encoded bytes
	data := e.Bytes()

	return data, "application/json", nil
}
