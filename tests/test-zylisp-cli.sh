#!/bin/bash

FAIL_SUITE=0
FAILURES=0
PASSES=0
BASE_OUTDIR=/tmp/zylisp

. ./tests/common.sh

echo "Batch AST multiple file creation"
	OUTDIR=$(create-tmp-dir "ASTs_dir")
	zylisp -cli -ast -dir $OUTDIR examples/*.gsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals $AST_COUNT $EXAMPLES_COUNT
	echo

echo "Batch Go multiple file creation"
	OUTDIR=$(create-tmp-dir "GOs_dir")
	zylisp -cli -go -dir $OUTDIR examples/*.gsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals $GO_COUNT $EXAMPLES_COUNT
	echo

echo "Batch byte-code multiple file creation"
	OUTDIR=$(create-tmp-dir "Bytecodes_dir")
	zylisp -cli -bytecode -dir $OUTDIR examples/*.gsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $BYTECODE_COUNT $EXAMPLES_COUNT
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

echo "Batch byte-code explicit output file creation"
	OUTDIR=$(create-tmp-dir "Bytecode_file")
	zylisp -cli -bytecode -o $OUTDIR/factorial examples/factorial.gsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $BYTECODE_COUNT 1
	echo

echo "Batch AST implicit output file (using dir) creation"
	OUTDIR=$(create-tmp-dir "AST_dir")
	zylisp -cli -ast -dir $OUTDIR examples/factorial.gsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals $AST_COUNT 1
	echo

echo "Batch Go implicit output file (using dir) creation"
	OUTDIR=$(create-tmp-dir "GO_dir")
	zylisp -cli -go -dir $OUTDIR examples/factorial.gsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals $GO_COUNT 1
	echo

echo "Batch byte-code implicit output file (using dir) creation"
	OUTDIR=$(create-tmp-dir "Bytecode_dir")
	zylisp -cli -bytecode -dir $OUTDIR examples/factorial.gsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $BYTECODE_COUNT 1
	echo

echo "Tests passed: $PASSES"
echo "Tests failed: $FAILURES"
echo

if [ ! -z "$BASE_OUTDIR" ]; then
	rm -rf $BASE_OUTDIR
fi

if [ "$FAIL_SUITE" == 1 ]; then
	exit 1
fi
