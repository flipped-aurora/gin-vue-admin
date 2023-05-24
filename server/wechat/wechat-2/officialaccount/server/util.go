package server

import (
	"encoding/xml"
	"net/http"
)

var xmlContentType = []string{"application/xml; charset=utf-8"}
var plainContentType = []string{"text/plain; charset=utf-8"}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// Render render from bytes
func (srv *Server) Render(bytes []byte) {
	// debug
	// fmt.Println("response msg = ", string(bytes))
	srv.Writer.WriteHeader(200)
	_, err := srv.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

// String render from string
func (srv *Server) String(str string) {
	writeContextType(srv.Writer, plainContentType)
	srv.Render([]byte(str))
}

// XML render to xml
func (srv *Server) XML(obj interface{}) {
	writeContextType(srv.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	srv.Render(bytes)
}

// Query returns the keyed url query value if it exists
func (srv *Server) Query(key string) string {
	value, _ := srv.GetQuery(key)
	return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (srv *Server) GetQuery(key string) (string, bool) {
	req := srv.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}
