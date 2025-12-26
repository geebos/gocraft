#!/bin/bash

# Version number prefix
VERSION_PREFIX="v"

# Show help message
show_help() {
    echo "Usage: $0 [-a | -b | -c] [-m <message>] [-h]"
    echo ""
    echo "Options:"
    echo "  -a          Increment major version (A), reset B and C to 0"
    echo "  -b          Increment minor version (B), keep A unchanged, reset C to 0"
    echo "  -c          Increment patch version (C), keep A and B unchanged (default)"
    echo "  -m <msg>    Specify the tag annotation message"
    echo "  -h          Show help message"
    echo ""
    echo "Examples:"
    echo "  $0              # Increment patch version (default)"
    echo "  $0 -a           # Increment major version"
    echo "  $0 -b           # Increment minor version"
    echo "  $0 -c -m 'fix'" # Increment patch version with annotation"
}

# Get the latest version number
get_latest_version() {
    # Get all tags matching the pattern, sort by version number, take the latest
    git tag -l "${VERSION_PREFIX}*" | \
        sed "s/^${VERSION_PREFIX}//" | \
        sort -t. -k1,1n -k2,2n -k3,3n | \
        tail -n 1
}

# Parse version number
parse_version() {
    local version=$1
    if [[ -z "$version" ]]; then
        echo "0 0 0"
        return
    fi
    
    local major=$(echo "$version" | cut -d. -f1)
    local minor=$(echo "$version" | cut -d. -f2)
    local patch=$(echo "$version" | cut -d. -f3)
    
    # Use default values if parsing fails
    major=${major:-0}
    minor=${minor:-0}
    patch=${patch:-0}
    
    echo "$major $minor $patch"
}

# Main logic
main() {
    local bump_type="c"  # Default: increment patch version
    local message=""
    
    # Parse command line arguments
    while getopts "abcm:h" opt; do
        case $opt in
            a)
                bump_type="a"
                ;;
            b)
                bump_type="b"
                ;;
            c)
                bump_type="c"
                ;;
            m)
                message="$OPTARG"
                ;;
            h)
                show_help
                exit 0
                ;;
            \?)
                echo "Invalid option: -$OPTARG" >&2
                show_help
                exit 1
                ;;
        esac
    done
    
    # Check if current directory is a git repository
    if ! git rev-parse --git-dir > /dev/null 2>&1; then
        echo "Error: Current directory is not a git repository" >&2
        exit 1
    fi
    
    # Get latest version number
    local latest_version=$(get_latest_version)
    
    if [[ -z "$latest_version" ]]; then
        echo "No tag matching pattern '${VERSION_PREFIX}X.Y.Z' found, starting from 0.0.0"
        latest_version="0.0.0"
    fi
    
    echo "Current latest version: ${VERSION_PREFIX}${latest_version}"
    
    # Parse version number
    read major minor patch <<< $(parse_version "$latest_version")
    
    # Calculate new version based on bump_type
    case $bump_type in
        a)
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        b)
            minor=$((minor + 1))
            patch=0
            ;;
        c)
            patch=$((patch + 1))
            ;;
    esac
    
    local new_version="${major}.${minor}.${patch}"
    local new_tag="${VERSION_PREFIX}${new_version}"
    
    echo "New version: $new_tag"
    
    # Check if tag already exists
    if git rev-parse "$new_tag" >/dev/null 2>&1; then
        echo "Error: Tag '$new_tag' already exists" >&2
        exit 1
    fi
    
    # Set default message
    if [[ -z "$message" ]]; then
        message="Release $new_version"
    fi
    
    # Create tag
    git tag -a "$new_tag" -m "$message"
    
    if [[ $? -eq 0 ]]; then
        echo "✅ Successfully created tag: $new_tag"
        echo ""
        echo "Push to remote repository:"
        echo "  git push origin $new_tag"
    else
        echo "❌ Failed to create tag" >&2
        exit 1
    fi
}

main "$@"
