package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/prometheus/alertmanager/dispatch"
	"github.com/zinrai/amroutify/pkg/config"
	"github.com/zinrai/amroutify/pkg/routing"
)

var (
	configFile = flag.String("config", "alertmanager.yml", "Path to Alertmanager configuration file")
	testsFile  = flag.String("tests", "routing_tests.yml", "Path to routing test cases file")
	verbose    = flag.Bool("verbose", false, "Show detailed output for tests")
)

func main() {
	flag.Parse()

	amConfig, err := config.LoadAlertmanagerConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load Alertmanager config: %v", err)
	}

	testCases, err := config.LoadTestCases(*testsFile)
	if err != nil {
		log.Fatalf("Failed to load test cases: %v", err)
	}

	route := dispatch.NewRoute(amConfig.Route, nil)

	results := routing.RunTests(route, testCases, *verbose)

	failedTests := printResults(results, *verbose)

	if failedTests > 0 {
		os.Exit(1)
	}
}

// Displays the test results and returns the number of failed tests
func printResults(results []routing.TestResult, verbose bool) int {
	failedTests := 0

	for _, result := range results {
		if result.Success {
			fmt.Printf("PASS: %s\n", result.Name)
			if verbose {
				fmt.Printf("  Matched receivers: %v\n", result.Actual)
				for i, r := range result.Routes {
					fmt.Printf("  Route %d: %s\n", i+1, r.RouteOpts.Receiver)
				}
			}
		} else {
			fmt.Printf("FAIL: %s\n", result.Name)
			fmt.Printf("  Expected receivers: %v\n", result.Expected)
			fmt.Printf("  Actual receivers:   %v\n", result.Actual)
			failedTests++
		}
	}

	fmt.Printf("\nResults: %d passed, %d failed, %d total\n",
		len(results)-failedTests, failedTests, len(results))

	return failedTests
}
