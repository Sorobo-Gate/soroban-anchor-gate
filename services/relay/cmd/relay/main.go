package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/soroban-anchor-gate/relay/internal/listener"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	contractID := os.Getenv("SOROBAN_CONTRACT_ID")
	rpcURL := os.Getenv("SOROBAN_RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://soroban-testnet.stellar.org"
	}

	logger.Info("SorobanAnchor Gate Relay daemon starting", "rpc_url", rpcURL, "contract_id", contractID)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	eventChan := make(chan listener.EventPayload, 100)
	sub := listener.NewEventSubscriber(rpcURL, contractID, logger)

	go func() {
		if err := sub.PollEvents(ctx, 1, eventChan); err != nil && err != context.Canceled {
			logger.Error("Event poller failure", "error", err)
		}
	}()

	<-ctx.Done()
	logger.Info("Shutting down relay daemon gracefully")
}
