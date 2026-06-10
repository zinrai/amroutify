package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/prometheus/alertmanager/dispatch"
	"github.com/zinrai/amroutify/internal/config"
	"github.com/zinrai/amroutify/internal/routing"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	configFile  = flag.String("config", "alertmanager.yml", "Path to Alertmanager configuration file")
	testsFile   = flag.String("tests", "routing_tests.yml", "Path to routing test cases file")
	showVersion = flag.Bool("version", false, "Show version information")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("amroutify %s (commit: %s, built: %s)\n", version, commit, date)
		return
	}

	amConfig, err := config.LoadAlertmanagerConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load Alertmanager config: %v", err)
	}

	testCases, err := config.LoadTestCases(*testsFile)
	if err != nil {
		log.Fatalf("Failed to load test cases: %v", err)
	}

	route := dispatch.NewRoute(amConfig.Route, nil)

	results := routing.RunTests(route, testCases)

	failedTests := printResults(results)

	if failedTests > 0 {
		os.Exit(1)
	}
}

// Displays the test results and returns the number of failed tests
func printResults(results []routing.TestResult) int {
	failedTests := 0

	for _, result := range results {
		if result.Success {
			fmt.Printf("PASS: %s\n", result.Name)
			fmt.Printf("  Receivers: %v\n", result.Actual)
		} else {
			fmt.Printf("FAIL: %s\n", result.Name)
			fmt.Printf("  Expected: %v\n", result.Expected)
			fmt.Printf("  Actual:   %v\n", result.Actual)
			failedTests++
		}
	}

	fmt.Printf("\nResults: %d passed, %d failed, %d total\n",
		len(results)-failedTests, failedTests, len(results))

	return failedTests
}
