#!/bin/bash

# Test script for running all unit tests in the project

# Default options
VERBOSE=false
COVERAGE=false
COVERAGE_OUTPUT=""
PACKAGES="./..."

# Show help message
show_help() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -v, --verbose          Run tests in verbose mode"
    echo "  -c, --coverage         Show test coverage"
    echo "  -o, --output <file>    Write coverage profile to file (requires -c)"
    echo "  -p, --packages <path>  Test specific packages (default: ./...)"
    echo "  -h, --help             Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                     # Run all tests"
    echo "  $0 -v                  # Run tests in verbose mode"
    echo "  $0 -c                  # Run tests with coverage"
    echo "  $0 -c -o coverage.out  # Generate coverage profile"
    echo "  $0 -p ./pkg/gslice     # Test specific package"
}

# Parse command line arguments
parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -v|--verbose)
                VERBOSE=true
                shift
                ;;
            -c|--coverage)
                COVERAGE=true
                shift
                ;;
            -o|--output)
                if [[ -z "$2" ]]; then
                    echo "Error: -o requires a filename" >&2
                    exit 1
                fi
                COVERAGE_OUTPUT="$2"
                COVERAGE=true
                shift 2
                ;;
            -p|--packages)
                if [[ -z "$2" ]]; then
                    echo "Error: -p requires a package path" >&2
                    exit 1
                fi
                PACKAGES="$2"
                shift 2
                ;;
            -h|--help)
                show_help
                exit 0
                ;;
            *)
                echo "Unknown option: $1" >&2
                show_help
                exit 1
                ;;
        esac
    done
}

# Build test command
build_test_cmd() {
    local cmd="go test"
    
    if [[ "$VERBOSE" == true ]]; then
        cmd="$cmd -v"
    fi
    
    if [[ "$COVERAGE" == true ]]; then
        cmd="$cmd -cover"
        if [[ -n "$COVERAGE_OUTPUT" ]]; then
            cmd="$cmd -coverprofile=$COVERAGE_OUTPUT"
        fi
    fi
    
    cmd="$cmd $PACKAGES"
    
    echo "$cmd"
}

# Main function
main() {
    parse_args "$@"
    
    echo "Running tests..."
    echo "Packages: $PACKAGES"
    if [[ "$VERBOSE" == true ]]; then
        echo "Mode: Verbose"
    fi
    if [[ "$COVERAGE" == true ]]; then
        echo "Coverage: Enabled"
        if [[ -n "$COVERAGE_OUTPUT" ]]; then
            echo "Coverage output: $COVERAGE_OUTPUT"
        fi
    fi
    echo ""
    
    # Build and execute test command
    local test_cmd=$(build_test_cmd)
    eval "$test_cmd"
    local exit_code=$?
    
    echo ""
    if [[ $exit_code -eq 0 ]]; then
        echo "✅ All tests passed"
        
        if [[ -n "$COVERAGE_OUTPUT" ]]; then
            echo ""
            echo "Coverage profile saved to: $COVERAGE_OUTPUT"
            echo "View coverage report:"
            echo "  go tool cover -html=$COVERAGE_OUTPUT"
        fi
    else
        echo "❌ Tests failed"
        exit $exit_code
    fi
}

main "$@"

