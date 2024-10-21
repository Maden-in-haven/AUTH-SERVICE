// Code generated by ogen, DO NOT EDIT.

package gen

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// APIAuthLoginPost invokes POST /api/auth/login operation.
	//
	// Авторизует пользователя с использованием учетных
	// данных (логин/пароль).
	//
	// POST /api/auth/login
	APIAuthLoginPost(ctx context.Context, request *APIAuthLoginPostReq) (APIAuthLoginPostRes, error)
	// APIAuthRefreshPost invokes POST /api/auth/refresh operation.
	//
	// Обновляет токен доступа с использованием
	// действующего токена обновления.
	//
	// POST /api/auth/refresh
	APIAuthRefreshPost(ctx context.Context, request *APIAuthRefreshPostReq) (APIAuthRefreshPostRes, error)
	// APIAuthVerifyPost invokes POST /api/auth/verify operation.
	//
	// Проверяет действительность переданного токена JWT.
	//
	// POST /api/auth/verify
	APIAuthVerifyPost(ctx context.Context, request *APIAuthVerifyPostReq) (APIAuthVerifyPostRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// APIAuthLoginPost invokes POST /api/auth/login operation.
//
// Авторизует пользователя с использованием учетных
// данных (логин/пароль).
//
// POST /api/auth/login
func (c *Client) APIAuthLoginPost(ctx context.Context, request *APIAuthLoginPostReq) (APIAuthLoginPostRes, error) {
	res, err := c.sendAPIAuthLoginPost(ctx, request)
	return res, err
}

func (c *Client) sendAPIAuthLoginPost(ctx context.Context, request *APIAuthLoginPostReq) (res APIAuthLoginPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/auth/login"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIAuthLoginPost",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/auth/login"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAPIAuthLoginPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIAuthLoginPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIAuthRefreshPost invokes POST /api/auth/refresh operation.
//
// Обновляет токен доступа с использованием
// действующего токена обновления.
//
// POST /api/auth/refresh
func (c *Client) APIAuthRefreshPost(ctx context.Context, request *APIAuthRefreshPostReq) (APIAuthRefreshPostRes, error) {
	res, err := c.sendAPIAuthRefreshPost(ctx, request)
	return res, err
}

func (c *Client) sendAPIAuthRefreshPost(ctx context.Context, request *APIAuthRefreshPostReq) (res APIAuthRefreshPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/auth/refresh"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIAuthRefreshPost",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/auth/refresh"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAPIAuthRefreshPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIAuthRefreshPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIAuthVerifyPost invokes POST /api/auth/verify operation.
//
// Проверяет действительность переданного токена JWT.
//
// POST /api/auth/verify
func (c *Client) APIAuthVerifyPost(ctx context.Context, request *APIAuthVerifyPostReq) (APIAuthVerifyPostRes, error) {
	res, err := c.sendAPIAuthVerifyPost(ctx, request)
	return res, err
}

func (c *Client) sendAPIAuthVerifyPost(ctx context.Context, request *APIAuthVerifyPostReq) (res APIAuthVerifyPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/auth/verify"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIAuthVerifyPost",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/auth/verify"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAPIAuthVerifyPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIAuthVerifyPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
