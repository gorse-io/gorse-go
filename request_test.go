// Copyright 2026 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestPreservesNilPatchFields(t *testing.T) {
	for _, tc := range []struct {
		name       string
		categories []string
		labels     any
		want       string
	}{
		{name: "unset", want: `{"IsHidden":null,"Categories":null,"Timestamp":null,"Labels":null,"Comment":null}`},
		{name: "typed nil map", labels: map[string]any(nil), want: `{"IsHidden":null,"Categories":null,"Timestamp":null,"Labels":null,"Comment":null}`},
		{name: "empty", categories: []string{}, labels: map[string]any{}, want: `{"IsHidden":null,"Categories":[],"Timestamp":null,"Labels":{},"Comment":null}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var body []byte
			var readErr error
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, readErr = io.ReadAll(r.Body)
				_, _ = io.WriteString(w, `{"RowAffected":1}`)
			}))
			defer server.Close()

			result, err := NewGorseClient(server.URL, "").UpdateItem(context.Background(), "item", ItemPatch{
				Categories: tc.categories,
				Labels:     tc.labels,
			})
			require.NoError(t, err)
			require.NoError(t, readErr)
			require.Equal(t, 1, result.RowAffected)
			require.JSONEq(t, tc.want, string(body))
		})
	}
}

func TestRequestRejectsInvalidJSON(t *testing.T) {
	for _, tc := range []struct {
		name string
		body string
	}{
		{name: "duplicate names", body: `{"RowAffected":1,"RowAffected":2}`},
		{name: "invalid UTF-8", body: "{\"Comment\":\"\xff\"}"},
		{name: "trailing value", body: `{"RowAffected":1} {}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = io.WriteString(w, tc.body)
			}))
			defer server.Close()

			_, err := request[map[string]any, any](context.Background(), NewGorseClient(server.URL, ""), http.MethodGet, server.URL, nil)
			require.Error(t, err)
		})
	}
}
