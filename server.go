package simpleimageserver

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/xid"
	"go.uber.org/zap"
)

// Server is main package struct
type Server struct {
	prefix string

	logger *zap.Logger
}

// Opt is opt func for server initialization
type Opt func(*Server)

// Prefix sets prefix to server
func Prefix(prefix string) Opt {
	return func(s *Server) {
		s.prefix = prefix
	}
}

// New returns new server
func New(fns ...Opt) (*Server, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.Encoding = "console"
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	s := &Server{logger: logger, prefix: "/"}

	for _, fn := range fns {
		fn(s)
	}
	return s, nil
}

// Run starts web server
func (s *Server) Run() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fs := http.FileServer(http.Dir("./"))
	http.Handle(s.prefix, http.StripPrefix(s.prefix, fs))

	http.HandleFunc("/upload", s.UploadHandler)

	s.logger.Info("start web server", zap.String("port", port))
	http.ListenAndServe(":"+port, nil)
}

// ServeHandler serve uploaded images
// func (s *Server) ServeHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Static
// }

// UploadHandler handle uploads
func (s *Server) UploadHandler(w http.ResponseWriter, r *http.Request) {
	uploadID := xid.New()
	s.logger.Info("handle upload", zap.String("id", uploadID.String()))

	r.ParseMultipartForm(5 << 30)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	f, err := os.OpenFile(fmt.Sprintf("%s.jpg", uploadID), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}