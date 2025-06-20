package main

import (
	acmetest "github.com/cert-manager/cert-manager/test/acme"
	"os"
	"testing"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.
	//

	// Uncomment the below fixture when implementing your custom DNS provider
	fixture := acmetest.NewFixture(&stratoDNSProviderSolver{},
		acmetest.SetResolvedZone(zone),
		acmetest.SetResolvedFQDN("_acme-challenge."+zone),
		acmetest.SetAllowAmbientCredentials(false),
		acmetest.SetManifestPath("testdata/strato"),
	)

	fixture.RunConformance(t)
}
