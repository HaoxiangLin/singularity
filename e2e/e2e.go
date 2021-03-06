// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package e2e

import (
	"os"
	"testing"
)

// RunE2ETests is the main func to trigger the test suite
func RunE2ETests(t *testing.T) {
	if os.Getenv("SINGULARITY_E2E") == "" {
		t.Skip("Skipping e2e tests, SINGULARITY_E2E not set")
	} else {
		t.Log("Running E2E tests for Singularity")
		Run(t)
	}
}
