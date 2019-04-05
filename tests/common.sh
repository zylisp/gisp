#!/bin/bash

GREEN='\033[1;32m'
RED='\033[1;31m'
CLEAR_COLOR='\033[0m'

FAIL_SUITE=0
FAILURES=0
PASSES=0
BASE_OUTDIR=/tmp/zylisp
PATH=$PATH:./bin
ZY=zylisp
ZYC=zyc

function create-tmp-dir () {
	TMPDIR="$BASE_OUTDIR/cli/$1/`date|sed 's/[ :]/_/g'`"
	mkdir -p "$TMPDIR"
	echo "$TMPDIR"
}

function count-files () {
	FILEPATH=$1
	MATCH=$2
	echo $(find "$FILEPATH" -type f -name "$MATCH" |wc -l|tr -d ' ')
}

function count-files-without-extension () {
	FILEPATH=$1
	echo $(find "$FILEPATH" -type f ! -name "*.*" |wc -l|tr -d ' ')
}

function num-equals () {
	EXPECTED=$1
	ACTUAL=$2
	if [ ! -z "$ACTUAL" ] && [ ! -z "$ACTUAL" ] && [ "$ACTUAL" -eq "$EXPECTED" ]; then
		echo -e "  ${GREEN}PASS${CLEAR_COLOR}"
		PASSES=$((PASSES+1))
	else
		echo -e "  ${RED}FAIL${CLEAR_COLOR}: (actual value '$ACTUAL' != expected value '$EXPECTED')"
		FAIL_SUITE=1
		FAILURES=$((FAILURES+1))
	fi
}

function exit-zero () {
	EXIT_CODE=$1
	FAIL_MSG=$2
	if [ $EXIT_CODE -eq 0 ]; then
		echo -e "  ${GREEN}PASS${CLEAR_COLOR}"
		PASSES=$((PASSES+1))
	else
		echo -e "  ${RED}FAIL${CLEAR_COLOR}: $FAIL_MSG"
		FAIL_SUITE=1
		FAILURES=$((FAILURES+1))
	fi
}

function tear-down () {
	if [ ! -z "$BASE_OUTDIR" ]; then
		rm -rf $BASE_OUTDIR
	fi
}

function print-summary () {
	echo -e "Tests passed: ${GREEN}$PASSES${CLEAR_COLOR}"
	echo -e "Tests failed: ${RED}$FAILURES${CLEAR_COLOR}"
	echo
	if [ "$FAIL_SUITE" == 1 ]; then
		exit 1
	fi
}

EXAMPLES_COUNT=$(count-files examples/ "*.zsp")
echo