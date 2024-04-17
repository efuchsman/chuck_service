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
		Categories: []string{},
		CreatedAt:  "test",
		IconUrl:    "www.test.com",
		Id:         "109i092u3r08",
		Url:        "www.fkdpojfpsodj.com",
		Value:      "Chuck Norris",
	}

	testResponse2 = &ChuckJoke{
		Categories: []string{"food"},
		CreatedAt:  "test",
		IconUrl:    "www.test.com",
		Id:         "109i093u3r08",
		Url:        "www.fkdpojfpsodj.com",
		Value:      "Chuck Norris",
	}

	testCategories = &[]string{
		"animal",
		"career",
		"celebrity",
		"dev",
		"explicit",
		"fashion",
		"food",
		"history",
		"money",
		"movie",
		"music",
		"political",
		"religion",
		"science",
		"sport",
		"travel",
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
			description: "Success: a joke is returned with category",
			category:    "food",
			http: HTTPTestClient{
				GetData: &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader(`{"categories":["food"],"created_at":"test","icon_url":"www.test.com","id":"109i093u3r08","url":"www.fkdpojfpsodj.com","value":"Chuck Norris"}`)),
				},
			},
			expectedOutput: testResponse2,
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
		{
			description: "Failure: Bad Request causing an unmarshal failure",
			category:    "{}",
			http: HTTPTestClient{
				GetData: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       ioutil.NopCloser(strings.NewReader(`"BAD REQUEST"`)),
				},
			},
			expectedOutput: nil,
			expectedErr:    errors.New("json: cannot unmarshal string into Go value of type chucknorris.ChuckJoke"),
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Log(tc.description)

			ncc := NewChuckNorrisService(&tc.http)
			joke, err := ncc.RandomJoke(tc.category)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr.Error(), err.Error(), "error message should match")
			} else {
				assert.NotNil(t, joke)
				assert.NoError(t, err, tc.description)
				require.NoError(t, err)
				assert.Equal(t, joke, tc.expectedOutput)
			}
		})
	}
}

func TestCategories(t *testing.T) {
	testCases := []struct {
		description    string
		http           HTTPTestClient
		expectedOutput *[]string
		expectedErr    error
	}{
		{
			description: "Success: categories are returned",
			http: HTTPTestClient{
				GetData: &http.Response{
					StatusCode: http.StatusOK,
					Body: ioutil.NopCloser(strings.NewReader(`["animal",
					"career",
					"celebrity",
					"dev",
					"explicit",
					"fashion",
					"food",
					"history",
					"money",
					"movie",
					"music",
					"political",
					"religion",
					"science",
					"sport",
					"travel"]`)),
				},
			},
			expectedOutput: testCategories,
			expectedErr:    nil,
		},
		{
			description: "Failure: Bad Request",
			http: HTTPTestClient{
				GetErr: errBadRequest,
			},
			expectedOutput: nil,
			expectedErr:    errBadRequest,
		},
		{
			description: "Failure: Bad Request",
			http: HTTPTestClient{
				GetData: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       ioutil.NopCloser(strings.NewReader(`"BAD REQUEST"`)),
				},
			},
			expectedOutput: nil,
			expectedErr:    errors.New("json: cannot unmarshal string into Go value of type []string"),
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Log(tc.description)

			ncc := NewChuckNorrisService(&tc.http)
			categories, err := ncc.Categories()
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr.Error(), err.Error(), "error message should match")
			} else {
				assert.NotNil(t, categories)
				assert.NoError(t, err, tc.description)
				require.NoError(t, err)
				assert.Equal(t, categories, tc.expectedOutput)
			}
		})
	}
}
