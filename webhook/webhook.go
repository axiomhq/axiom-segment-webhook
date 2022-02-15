package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/axiomhq/axiom-go/axiom"
	"go.uber.org/zap"
)

var (
	timestampField = "timestamp"
	datasetName    = "axiom_segment_webhook"
)

// Webhook handles a segments.io webhook payload.
type Webhook struct {
	logger *zap.Logger
	client *axiom.Client
}

// NewWebhook creates a new Webhook.
func NewWebhook(ctx context.Context, logger *zap.Logger, client *axiom.Client) (*Webhook, error) {
	if _, err := client.Datasets.Create(ctx, axiom.DatasetCreateRequest{
		Name:        datasetName,
		Description: "Segment events",
	}); err != nil && !errors.Is(err, axiom.ErrExists) {
		return nil, fmt.Errorf("create dataset %q: %w", datasetName, err)
	}

	return &Webhook{
		logger: logger,
		client: client,
	}, nil
}

func (w *Webhook) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		return
	}

	if err := w.sendEvent(req); err != nil {
		w.logger.Error(err.Error())
	}
}

func (w *Webhook) sendEvent(req *http.Request) error {
	var event axiom.Event
	if err := json.NewDecoder(req.Body).Decode(&event); err != nil {
		return err
	}

	opts := axiom.IngestOptions{
		TimestampField:  timestampField,
		TimestampFormat: time.RFC3339,
	}

	status, err := w.client.Datasets.IngestEvents(req.Context(), datasetName, opts, event)
	if err != nil {
		w.logger.Error("error ingesting event", zap.Error(err))
		return err
	}

	w.logger.Info("ingested successfully",
		zap.Uint64("ingested", status.Ingested),
		zap.Uint64("failed", status.Failed),
		zap.Uint64("processed_bytes", status.ProcessedBytes),
		zap.Uint32("blocks_created", status.BlocksCreated),
		zap.Uint32("wal_length", status.WALLength),
	)
	for k, v := range status.Failures {
		w.logger.Warn("failed to ingest event",
			zap.Int("id", k+1),
			zap.Time("timestamp", v.Timestamp),
			zap.String("error", v.Error),
		)
	}

	return nil
}
