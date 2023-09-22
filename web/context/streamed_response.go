package context

import (
	"io"
	"net/http"
)

type StreamedResponse struct {
	*Response
	step func(w io.Writer) bool
}

func (streamed *StreamedResponse) SetStep(step func(w io.Writer) bool) *StreamedResponse {
	streamed.step = step
	return streamed
}

func (streamed *StreamedResponse) Send(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
	default:
		for {
			if !streamed.step(w) {
				break
			}
		}
		w.(http.Flusher).Flush()
	}
}
