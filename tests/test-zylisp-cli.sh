#!/bin/bash

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
	echo $(find $1 -type f -name $2 |wc -l)
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

EXAMPLES_COUNT=$(count-files examples/ "*.gsp")
echo

echo "Batch AST multiple file creation"
	OUTDIR=$(create-tmp-dir "ASTs")
	zylisp -cli -ast -dir $OUTDIR examples/*.gsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals $AST_COUNT $EXAMPLES_COUNT
	echo

echo "Batch Go multiple file creation"
	OUTDIR=$(create-tmp-dir "GOs")
	zylisp -cli -go -dir $OUTDIR examples/*.gsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals $GO_COUNT $EXAMPLES_COUNT
	echo

echo "Batch AST explicit output file creation"
	OUTDIR=$(create-tmp-dir "AST_file")
	zylisp -cli -ast -o $OUTDIR/factorial.ast examples/factorial.gsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals $AST_COUNT 1
	echo

echo "Batch Go explicit output file creation"
	OUTDIR=$(create-tmp-dir "GO_file")
	zylisp -cli -go -o $OUTDIR/factorial.go examples/factorial.gsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals $GO_COUNT 1
	echo

echo "Batch AST implicit output file (using dir) creation"
	OUTDIR=$(create-tmp-dir "AST_dir")
	zylisp -cli -ast -dir $OUTDIR examples/factorial.gsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals $AST_COUNT 1
	echo

echo "Batch Go explicit output file (using dir) creation"
	OUTDIR=$(create-tmp-dir "GO_dir")
	zylisp -cli -go -dir $OUTDIR examples/factorial.gsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals $GO_COUNT 1
	echo

echo "Tests passed: $PASSES"
echo "Tests failed: $FAILURES"

if [ ! -z "$BASE_OUTDIR" ]; then
	rm -rf $BASE_OUTDIR
fi

if [ "$FAIL_SUITE" == 1 ]; then
	exit 1
fi
