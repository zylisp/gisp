#!/bin/bash

GREEN='\033[1;32m'
RED='\033[1;31m'
CLEAR_COLOR='\033[0m'

FAIL_SUITE=0
FAILURES=0
PASSES=0
BASE_OUTDIR=/tmp/zylisp

function create-tmp-dir () {
	TMPDIR="$BASE_OUTDIR/cli/$1/`date|sed 's/[ :]/_/g'`"
	mkdir -p "$TMPDIR"
	echo "$TMPDIR"
}

function count-files () {
	echo $(find "$1" -type f -name "$2" |wc -l|tr -d ' ')
}

function count-files-without-extension () {
	echo $(find "$1" -type f ! -name "*.*" |wc -l|tr -d ' ')
}

function num-equals () {
	if [ ! -z "$1" ] && [ ! -z "$1" ] && [ "$1" -eq "$2" ]; then
		echo -e "  ${GREEN}PASS${CLEAR_COLOR}"
		PASSES=$((PASSES+1))
	else
		echo -e "  ${RED}FAIL${CLEAR_COLOR}: ($1 != $2)"
		FAIL_SUITE=1
		FAILURES=$((FAILURES+1))
	fi
}

EXAMPLES_COUNT=$(count-files examples/ "*.zsp")
echo
