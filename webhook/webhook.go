package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/axiomhq/axiom-go/axiom"
	"go.uber.org/zap"
)

var (
	logger           *zap.Logger
	timestampField   = "timestamp"
	datasetName      = "axiom_segement_webhook"
	alreadyExistsErr = fmt.Errorf("API error 409: entity exists: entity exists")
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// type check on axiom.Event incase it's ever not a map[string]interface{}
	// so we can use unsafe.Pointer for a quick type conversion instead of allocating a new slice
	if !reflect.TypeOf(axiom.Event{}).ConvertibleTo(reflect.TypeOf(map[string]interface{}{})) {
		panic("axiom.Event is not a map[string]interface{}, please contact support")
	}
}

type Webhook struct {
	client *axiom.Client
}

func NewWebhook(client *axiom.Client) (*Webhook, error) {
	axiomReq := axiom.DatasetCreateRequest{
		Name:        datasetName,
		Description: "Segment events",
	}
	_, err := client.Datasets.Create(context.Background(), axiomReq)
	if err != nil && err.Error() != alreadyExistsErr.Error() {
		return nil, err
	}

	return &Webhook{
		client: client,
	}, nil
}

func (m *Webhook) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	if err := m.sendEvent(req); err != nil {
		logger.Error(err.Error())
	}
}

func (m *Webhook) sendEvent(req *http.Request) error {
	opts := axiom.IngestOptions{
		TimestampField:  timestampField,
		TimestampFormat: time.RFC3339,
	}

	var event axiom.Event
	if err := json.NewDecoder(req.Body).Decode(&event); err != nil {
		return err
	}

	status, err := m.client.Datasets.IngestEvents(
		context.Background(),
		datasetName,
		opts,
		event,
	)

	if err != nil {
		logger.Error("error ingesting event", zap.Error(err))
		return err
	}

	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(status); err != nil {
		return err
	}

	logger.Info(buf.String())
	return nil
}
