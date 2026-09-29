package listener

import (
	"context"
	"log/slog"
	"time"
)

type EventPayload struct {
	EscrowID       uint64
	ProfileHashHex string
	PayoutAmount   int64
	ContractID     string
}

type EventSubscriber struct {
	rpcURL     string
	contractID string
	logger     *slog.Logger
}

func NewEventSubscriber(rpcURL, contractID string, logger *slog.Logger) *EventSubscriber {
	return &EventSubscriber{
		rpcURL:     rpcURL,
		contractID: contractID,
		logger:     logger,
	}
}

// PollEvents polls the Soroban RPC for DisbursementAuthorized events
func (s *EventSubscriber) PollEvents(ctx context.Context, startLedger uint32, eventChan chan<- EventPayload) error {
	s.logger.Info("Starting Soroban RPC event polling", "contract_id", s.contractID, "start_ledger", startLedger)

	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Polling routine executes here
		}
	}
}
