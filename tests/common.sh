#!/bin/bash

function create-tmp-dir () {
	TMPDIR="$BASE_OUTDIR/cli/$1/`date|sed 's/[ :]/_/g'`"
	mkdir -p "$TMPDIR"
	echo "$TMPDIR"
}

function count-files () {
	echo $(find $1 -type f -name $2 |wc -l)
}

function count-files-without-extension () {
	echo $(find $1 -type f ! -name "*.*" |wc -l)
}

function num-equals () {
	if [ ! -z "$1" ] && [ ! -z "$1" ] && [ "$1" -eq "$2" ]; then
		echo "  PASS"
		PASSES=$((PASSES+1))
	else
		echo "  FAIL: ($1 != $2)"
		FAIL_SUITE=1
		FAILURES=$((FAILURES+1))
	fi
}

EXAMPLES_COUNT=$(count-files examples/ "*.zsp")
echo
