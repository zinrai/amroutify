package routing

import (
	"sort"

	"github.com/prometheus/alertmanager/dispatch"
	"github.com/prometheus/common/model"
	"github.com/zinrai/amroutify/pkg/config"
)

type TestResult struct {
	Name     string
	Success  bool
	Expected []string
	Actual   []string
	Routes   []*dispatch.Route // Only used when verbose=true
}

// Executes all the test cases against the given routing tree and returns results
func RunTests(route *dispatch.Route, testCases []config.TestCase, verbose bool) []TestResult {
	results := make([]TestResult, 0, len(testCases))

	for _, tc := range testCases {
		result := runTest(route, tc, verbose)
		results = append(results, result)
	}

	return results
}

// Executes a single test case
func runTest(route *dispatch.Route, tc config.TestCase, verbose bool) TestResult {
	result := TestResult{
		Name:     tc.Name,
		Expected: tc.ExpectedReceivers,
	}

	// Build label set
	labelSet := model.LabelSet{}
	for k, v := range tc.Labels {
		labelSet[model.LabelName(k)] = model.LabelValue(v)
	}

	// Find matching routes
	matchingRoutes := route.Match(labelSet)

	// Store routes for verbose output
	if verbose {
		result.Routes = matchingRoutes
	}

	// Extract receivers
	var receivers []string
	for _, r := range matchingRoutes {
		receivers = append(receivers, r.RouteOpts.Receiver)
	}
	result.Actual = receivers

	// Compare (order-independent)
	result.Success = CompareReceivers(tc.ExpectedReceivers, receivers)

	return result
}

// Compares two slices of receivers in an order-independent way
func CompareReceivers(expected, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}

	expectedCopy := make([]string, len(expected))
	actualCopy := make([]string, len(actual))

	copy(expectedCopy, expected)
	copy(actualCopy, actual)

	sort.Strings(expectedCopy)
	sort.Strings(actualCopy)

	for i, v := range expectedCopy {
		if actualCopy[i] != v {
			return false
		}
	}

	return true
}
