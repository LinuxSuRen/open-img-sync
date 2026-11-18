// Package server 提供 Web + API 服务。骨架阶段只包含健康检查与版本接口，
// 资源路由、鉴权、SSE 等随 #2 - #6、#10 逐步接入。
package server

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

// Options 聚合服务启动所需的注入项。
type Options struct {
	Version   string
	Commit    string
	BuildTime string
}

// Server 持有路由与依赖。
type Server struct {
	opt  Options
	mux  *http.ServeMux
	http *http.Server
}

// New 构造服务实例。
func New(opt Options) *Server {
	s := &Server{opt: opt, mux: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.handleHealthz)
	s.mux.HandleFunc("GET /api/version", s.handleVersion)
}

// Handler 暴露路由供测试与上层复用。
func (s *Server) Handler() http.Handler { return s.mux }

// ListenAndServe 阻塞监听。
func (s *Server) ListenAndServe(addr string) error {
	s.http = &http.Server{
		Addr:              addr,
		Handler:           s.mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.http.Serve(ln)
}

func (s *Server) handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) handleVersion(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"version":    s.opt.Version,
		"commit":     s.opt.Commit,
		"build_time": s.opt.BuildTime,
	})
}
