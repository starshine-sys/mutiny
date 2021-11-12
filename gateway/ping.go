package gateway

import "emperror.dev/errors"

type ping struct {
	Type string `json:"type"`
	Data int64  `json:"data"`
}

type wserror struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

// ErrNotAuthenticatedEvent is returned when the first event from the websocket isn't the Authenticated or Error event.
const ErrNotAuthenticatedEvent = errors.Sentinel("first event wasn't Authenticated")

// Authentication errors
const (
	ErrUnlabeled             = errors.Sentinel("unlabeled error")
	ErrInternalError         = errors.Sentinel("internal error")
	ErrInvalidSession        = errors.Sentinel("invalid session")
	ErrOnboardingNotFinished = errors.Sentinel("onboarding not finished")
	ErrAlreadyAuthenticated  = errors.Sentinel("already authenticated")
)
