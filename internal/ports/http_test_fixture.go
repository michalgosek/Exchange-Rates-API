package ports

import (
	"context"
	"encoding/json"
	"exchange-rates-api/internal/app"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

type HTTPTestFixture struct {
	t    *testing.T
	http *HTTP
}

func (h *HTTPTestFixture) NewResponseRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func (h *HTTPTestFixture) DecodeResponse(res *http.Response, dst any) {
	defer res.Body.Close()

	h.t.Helper()
	err := json.NewDecoder(res.Body).Decode(&dst)
	require.NoError(h.t, err)
}

func (h *HTTPTestFixture) NewRequestWithParams(ctx context.Context, method, url string, params url.Values) *http.Request {
	h.t.Helper()
	req := h.NewRequestWithContext(ctx, method, url, nil)
	req.URL.RawQuery = params.Encode()

	return req
}

func (h *HTTPTestFixture) NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) *http.Request {
	h.t.Helper()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	require.Nil(h.t, err)
	require.NotNil(h.t, req)

	return req
}

func (h *HTTPTestFixture) ServeHTTP(req *http.Request, rec *httptest.ResponseRecorder) {
	h.t.Helper()
	h.http.eng.ServeHTTP(rec, req)
}

func NewHTTPServerTestFixture(t *testing.T, app *app.Application) *HTTPTestFixture {
	if t == nil {
		panic("testing parameter is required")
	}

	return &HTTPTestFixture{t: t, http: NewHTTP(app)}
}
