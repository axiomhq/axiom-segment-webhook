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
	logger         *zap.Logger
	timestampField = "timestamp"
	datasetPrefix  = "segement-io-%s"
	segmentMethods = map[string]struct{}{
		"page":     {},
		"screen":   {},
		"identify": {},
		"track":    {},
		"alias":    {},
		"group":    {},
	}
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

func NewWebhook(client *axiom.Client, types []string) (*Webhook, error) {
	for _, typ := range types {
		if _, ok := segmentMethods[typ]; !ok {
			return nil, fmt.Errorf("unknown type: %s", typ)
		}
	}

	for _, typ := range types {
		axiomReq := axiom.DatasetCreateRequest{
			Name:        fmt.Sprintf(datasetPrefix, typ),
			Description: fmt.Sprintf("Segment %s events ", typ),
		}
		client.Datasets.Create(context.Background(), axiomReq)
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
	dec := json.NewDecoder(req.Body)
	ev := axiom.Event{}

	if err := dec.Decode(&ev); err != nil {
		return err
	}

	typ, ok := ev["type"]
	if !ok {
		return fmt.Errorf("missing type field")
	}
	delete(ev, "type")

	dataset := typ.(string)

	opts := axiom.IngestOptions{
		TimestampField:  timestampField,
		TimestampFormat: time.RFC3339,
	}

	status, err := m.client.Datasets.IngestEvents(req.Context(), fmt.Sprintf(datasetPrefix, dataset), opts, ev)
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
