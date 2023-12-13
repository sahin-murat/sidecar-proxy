package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	var cases = []struct {
		path        string
		expectation bool
	}{
		{"/company", true},
		{"/tenant/sj3co3s4", false},
		{"/company/sd45f768", true},
		{"/account/acc74850", true},
		{"/company/account", true},
		{"/acc734340", true},
		{"/account/acc234234/user", true},
		{"/account/blocked", true}, // Since we have /account/:id route this should be always true, blocked word is read as :id param here
		{"/tenant/account/blocked", true},
		{"/tenant/account/acc23849", false},
		{"/tenant/account/acc23848", false},
	}

	for _, tc := range cases {
		require.Equal(t, tc.expectation, ValidatePath(tc.path), "Test is failing!")
	}
}
