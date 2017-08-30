#!/bin/bash

set -euo pipefail

scriptDir=$(dirname "$(readlink -f "$0")")

while [[ $# -gt 1 ]]; do
	key="$1"
	case $key in
		--tmp)
			tmp="$2"
			shift
			;;
		*)
			echo "unknown option $1"
			exit 1
			;;
	esac
	shift
done

tmp=${tmp:-/tmp}

go install github.com/dgraph-io/dgraph/cmd/bulkloader
go install github.com/dgraph-io/dgraph/cmd/dgcmp

function run_test {

	[[ $# == 3 ]] || { echo "bad args"; exit 1; }
	schema=$1
	dataDownload=$2
	wantHash=$3

	# Get the data file.
	dataFile="$scriptDir/$(basename "$dataDownload")"
	dataFile=${dataFile%?raw=true}
	if [[ ! -f "$dataFile" ]]; then
		wget "$dataDownload" -P "$scriptDir" -O "$dataFile"
	fi

	tmpDir=$(mktemp -p $tmp -d --suffix="_bulk_loader_speed_test")
	echo "$schema" > $tmpDir/sch.schema

	# Run bulk loader.
	bulkloader -tmp "$tmp" -b "$tmpDir" -s "$tmpDir/sch.schema" -r "$dataFile"

	# Get hash.
	gotHash=$(dgcmp -h "$tmpDir")

	# Cleanup
	rm -rf $tmpDir

	# Result
	if [[ $gotHash == $wantHash ]]; then
		echo "hash match"
	else
		echo "hash mismatch"
	fi
	echo "want: $wantHash"
	echo "got:  $gotHash"
}

echo "========================="
echo " 1 million data set      "
echo "========================="

run_test '
director.film:        uid @reverse @count .
genre:                uid @reverse .
initial_release_date: dateTime @index(year) .
name:                 string @index(term) .
starring:             uid @count .
' https://github.com/dgraph-io/tutorial/blob/master/resources/1million.rdf.gz?raw=true 100C4F8F0FFB0BF1
