package chucknorris

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testResponse = &ChuckJoke{
		Categories: nil,
		CreatedAt:  "test",
		IconUrl:    "www.test.com",
		Id:         "109i092u3r08",
		Url:        "www.fkdpojfpsodj.com",
		Value:      "Chuck Norris",
	}
)
var errBadRequest = errors.New("bad request")

func TestRandomJoke(t *testing.T) {
	testCases := []struct {
		description    string
		http           HTTPTestClient
		expectedOutput *ChuckJoke
		category       string
		expectedErr    error
	}{
		{
			description: "Success: a joke is returned with no category",
			category:    "",
			http: HTTPTestClient{
				GetData: &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader(`{"categories":[],"created_at":"test","icon_url":"www.test.com","id":"109i092u3r08","url":"www.fkdpojfpsodj.com","value":"Chuck Norris"}`)),
				},
			},
			expectedOutput: testResponse,
			expectedErr:    nil,
		},
		{
			description: "Failure: Bad Request",
			category:    "",
			http: HTTPTestClient{
				GetErr: errBadRequest,
			},
			expectedOutput: nil,
			expectedErr:    errBadRequest,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Log(tc.description)

			ncc := NewChuckNorrisService(&tc.http)
			encodes, err := ncc.RandomJoke(tc.category)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NotNil(t, encodes)
				assert.NoError(t, err, tc.description)
				require.NoError(t, err)
				assert.Equal(t, encodes, tc.expectedOutput)
			}
		})
	}
}
