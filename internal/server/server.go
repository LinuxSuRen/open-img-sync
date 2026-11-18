// Package server 提供 Web + API 服务。
// 骨架阶段提供 /healthz、/api/version 与系统信息接口；
// 鉴权中间件（#10）已就绪，资源路由随 #2/#3/#5 接入。
package server

import (
	"log/slog"
	"net"
	"net/http"
	"time"
)

// Options 聚合服务启动所需的注入项。
type Options struct {
	Version   string
	Commit    string
	BuildTime string
	// Token 非空时对 /api/* 启用 Bearer 鉴权（/healthz 豁免）
	Token string
	// DataDir 数据目录（系统信息展示用）
	DataDir string
	// MasterKeySource 主密钥来源：env / file
	MasterKeySource string
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
	s.mux.HandleFunc("GET /api/system/info", s.handleSystemInfo)
}

// Handler 暴露路由（含中间件）供测试与上层复用。
func (s *Server) Handler() http.Handler {
	return s.withAuth(withAccessLog(s.mux))
}

// ListenAndServe 阻塞监听。
func (s *Server) ListenAndServe(addr string) error {
	s.http = &http.Server{
		Addr:              addr,
		Handler:           s.Handler(),
		ReadHeaderTimeout: 10 * time.Second,
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	slog.Info("listening", "addr", addr, "version", s.opt.Version)
	return s.http.Serve(ln)
}

func (s *Server) handleHealthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleVersion(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"version":    s.opt.Version,
		"commit":     s.opt.Commit,
		"build_time": s.opt.BuildTime,
	})
}

func (s *Server) handleSystemInfo(w http.ResponseWriter, _ *http.Request) {
	authEnabled := s.opt.Token != ""
	writeJSON(w, http.StatusOK, map[string]any{
		"version":           s.opt.Version,
		"data_dir":          s.opt.DataDir,
		"master_key_source": s.opt.MasterKeySource,
		"auth_enabled":      authEnabled,
	})
}
