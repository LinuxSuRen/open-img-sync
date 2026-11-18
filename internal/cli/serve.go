package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/linuxsuren/open-img-sync/internal/server"
)

// serve 启动 Web + API 服务。骨架阶段仅提供 /healthz 与 /api/version，
// 资源接口随 #2/#3/#5 逐步接入。
func serve(args []string) error {
	fs := flag.NewFlagSet("serve", flag.ContinueOnError)
	addr := fs.String("addr", envOr("OIS_ADDR", ":8080"), "监听地址，也可用环境变量 OIS_ADDR")
	if err := fs.Parse(args); err != nil {
		return err
	}
	srv := server.New(server.Options{
		Version:   version,
		Commit:    commit,
		BuildTime: buildTime,
	})
	fmt.Printf("open-img-sync %s listening on %s\n", version, *addr)
	return srv.ListenAndServe(*addr)
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
