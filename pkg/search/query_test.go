package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trueBool = true

func TestQueryString(t *testing.T) {
	tests := []struct {
		name  string
		query Query
		out   string
	}{
		{
			name: "converts query to string",
			query: Query{
				Keywords: []string{"some", "keywords"},
				Qualifiers: Qualifiers{
					Archived:         &trueBool,
					Created:          "created",
					Followers:        "1",
					Fork:             "true",
					Forks:            "2",
					GoodFirstIssues:  "3",
					HelpWantedIssues: "4",
					In:               []string{"readme"},
					Language:         []string{"language"},
					License:          []string{"license"},
					Mirror:           &trueBool,
					Org:              "org",
					Pushed:           "updated",
					Repo:             "repo",
					Size:             "5",
					Stars:            "6",
					Topic:            []string{"topic"},
					Topics:           "7",
					User:             "user",
					Is:               "public",
				},
			},
			out: "some keywords archived:true created:created followers:1 fork:true forks:2 good-first-issues:3 help-wanted-issues:4 in:readme is:public language:language license:license mirror:true org:org pushed:updated repo:repo size:5 stars:6 topic:topic topics:7 user:user",
		},
		{
			name: "quotes keywords",
			query: Query{
				Keywords: []string{"quote keywords"},
			},
			out: "\"quote keywords\"",
		},
		{
			name: "quotes qualifiers",
			query: Query{
				Qualifiers: Qualifiers{
					Topic: []string{"quote qualifier"},
				},
			},
			out: "topic:\"quote qualifier\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.query.String())
		})
	}
}

func TestQualifiersMap(t *testing.T) {
	tests := []struct {
		name       string
		qualifiers Qualifiers
		out        map[string]string
	}{
		{
			name: "changes qualifiers to map",
			qualifiers: Qualifiers{
				Archived:         &trueBool,
				Created:          "created",
				Followers:        "1",
				Fork:             "true",
				Forks:            "2",
				GoodFirstIssues:  "3",
				HelpWantedIssues: "4",
				In:               []string{"readme"},
				Language:         []string{"language"},
				License:          []string{"license"},
				Mirror:           &trueBool,
				Org:              "org",
				Pushed:           "updated",
				Repo:             "repo",
				Size:             "5",
				Stars:            "6",
				Topic:            []string{"topic"},
				Topics:           "7",
				User:             "user",
				Is:               "public",
			},
			out: map[string]string{
				"archived":           "true",
				"created":            "created",
				"followers":          "1",
				"fork":               "true",
				"forks":              "2",
				"good-first-issues":  "3",
				"help-wanted-issues": "4",
				"in":                 "readme",
				"is":                 "public",
				"language":           "language",
				"license":            "license",
				"mirror":             "true",
				"org":                "org",
				"pushed":             "updated",
				"repo":               "repo",
				"size":               "5",
				"stars":              "6",
				"topic":              "topic",
				"topics":             "7",
				"user":               "user",
			},
		},
		{
			name: "excludes unset qualifiers from map",
			qualifiers: Qualifiers{
				Org:    "org",
				Pushed: "updated",
				Repo:   "repo",
				Size:   "5",
				Stars:  "6",
			},
			out: map[string]string{
				"org":    "org",
				"pushed": "updated",
				"repo":   "repo",
				"size":   "5",
				"stars":  "6",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.qualifiers.Map())
		})
	}
}
