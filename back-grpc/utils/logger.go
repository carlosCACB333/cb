package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	startTime := time.Now()
	resp, err := handler(ctx, req)
	endTime := time.Now()

	duration := endTime.Sub(startTime)
	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	l := log.Info()
	if err != nil {
		l = log.Error().Err(err)
	}

	l.
		Str("protocol", "GRPC").
		Str("method", info.FullMethod).
		Str("duration", duration.String()).
		Int("status", int(statusCode)).
		Str("status_text", statusCode.String()).
		Msg("GRPC Request")

	return resp, err

}

type ResponseW struct {
	http.ResponseWriter
	Status int
	Body   []byte
}

func (r *ResponseW) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *ResponseW) Write(b []byte) (int, error) {
	r.Body = b
	return r.ResponseWriter.Write(b)
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rw := &ResponseW{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}
		handler.ServeHTTP(rw, r)
		endTime := time.Now()

		duration := endTime.Sub(startTime)

		l := log.Info()

		if rw.Status != http.StatusOK {
			l = log.Error().Bytes("body", rw.Body)
		}

		l.
			Str("protocol", "HTTP").
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("duration", duration.String()).
			Int("status", rw.Status).
			Str("status_text", http.StatusText(rw.Status)).
			Msg("HTTP Request")
	})
}
