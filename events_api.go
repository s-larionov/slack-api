package slack

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fastjson"

	"slack/events"
	"slack/events/base"
)

type SubscribeCallback func(event interface{})

type EventsAPI struct {
	server          *http.Server
	subscriptions   map[base.Type][]chan events.Wrapper
	subscriptionsMu sync.RWMutex
}

func NewEventsAPI(listen, path string) *EventsAPI {
	eventsAPI := &EventsAPI{
		subscriptions: make(map[base.Type][]chan events.Wrapper),
	}

	handler := mux.NewRouter()
	handler.HandleFunc(path, eventsAPI.handle)
	eventsAPI.server = &http.Server{
		Addr:    listen,
		Handler: handler,
	}

	return eventsAPI
}

func (a *EventsAPI) ListenAndServe() error {
	err := a.server.ListenAndServe()

	if err == nil || err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (a *EventsAPI) Shutdown(ctx context.Context) error {
	err := a.server.Shutdown(ctx)

	if err == nil || err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (a *EventsAPI) Subscribe(event base.Type) <-chan events.Wrapper {
	a.subscriptionsMu.Lock()
	defer a.subscriptionsMu.Unlock()

	list, ok := a.subscriptions[event]
	if !ok {
		list = make([]chan events.Wrapper, 0)
	}

	ch := make(chan events.Wrapper, 1)

	a.subscriptions[event] = append(list, ch)

	return ch
}

func (a *EventsAPI) handle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	requestType := events.RequestType(fastjson.GetString(body, "type"))
	switch requestType {
	case events.RequestTypeVerification:
		a.handleVerificationRequest(body, w)
		return
	case events.RequestTypeWrapper:
		a.handleEventRequest(body, w)
		return
	case events.RequestTypeAppRateLimited:
		a.handleRateLimitedEvent(body, w)
		return
	default:
		log.Warn("unknown request type")
	}
}

func (a *EventsAPI) handleEventRequest(body []byte, w http.ResponseWriter) {
	var request events.Wrapper
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.WithError(err).Error("unable to parse event_callback request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.subscriptionsMu.RLock()
	event, ok := request.Event.(base.EventInterface)
	if !ok {
		log.Error("unable to parse event")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	subscriptions, ok := a.subscriptions[event.GetType()]
	if !ok {
		log.Error("unable to parse event")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, ch := range subscriptions {
		ch <- request
	}

	a.subscriptionsMu.RUnlock()
}

func (a *EventsAPI) handleVerificationRequest(body []byte, w http.ResponseWriter) {
	var request events.Verification
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.WithError(err).Error("unable to parse url_verification request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, _ = w.Write([]byte(request.Challenge))
	log.WithFields(log.Fields{
		"token":     request.Token,
		"challenge": request.Challenge,
	}).Info("verification request was processed")
}

func (a *EventsAPI) handleRateLimitedEvent(body []byte, w http.ResponseWriter) {
	var request events.AppRateLimited
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.WithError(err).Error("unable to parse app_rate_limited request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.WithField("minute_rate_limited", request.MinuteRateLimited).Warn("git a rate limited warning")
}
