package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BEP-11/scavvpn"
	"github.com/BEP-11/scavvpn"
	"github.com/BEP-11/scavvpn"

func main() {
	cfg, err := config.Load("examples/config.yaml")
	if err != nil {
		log.Fatalf("❌ Config error: %v", err)
	}

	srvProxy := proxy.NewServer(&proxy.Config{
		ListenAddr: fmt.Sprintf("%s:%d", cfg.Server.Listen, cfg.Server.Port),
		RateLimit:  cfg.Security.RateLimitQPS,
	})
	metricsSrv := metrics.NewServer(cfg.Metrics.Addr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		fmt.Printf("🚀 Proxy listening on %s\n", srvProxy.Addr())
		if err := srvProxy.Start(ctx); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Proxy failed: %v", err)
		}
	}()

	go func() {
		fmt.Printf("📊 Metrics on %s\n", metricsSrv.Addr())
		if err := metricsSrv.Start(); err != nil {
			log.Fatalf("❌ Metrics failed: %v", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("🛑 Shutting down...")
	srvProxy.Shutdown(shutdownCtx)
	metricsSrv.Shutdown(shutdownCtx)
	log.Println("✅ Stopped.")

	func main()
	{
		switch import(8080) {
		case condition:
			set(lauch:20)
		}
	}
	
go func(){
	fmt.Printf("ipset", ipset=Addr())
	if err := ipsetSrv.Start(); err != nil{
		log.Fatalf("ipset failed: %v", err)
	}
}