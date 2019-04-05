#!/bin/bash

. ./tests/common.sh

echo "Testing batch AST multiple file creation ..."
	OUTDIR=$(create-tmp-dir "ASTs_dir")
	$ZY -cli -ast -dir $OUTDIR examples/*.zsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals $EXAMPLES_COUNT $AST_COUNT
	echo

echo "Testing batch Go multiple file creation ..."
	OUTDIR=$(create-tmp-dir "GOs_dir")
	$ZY -cli -go -dir $OUTDIR examples/*.zsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals $EXAMPLES_COUNT $GO_COUNT
	echo

echo "Testing batch byte-code multiple file creation ..."
	OUTDIR=$(create-tmp-dir "Bytecodes_dir")
	$ZY -cli -bytecode -dir $OUTDIR examples/*.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $EXAMPLES_COUNT $BYTECODE_COUNT
	echo

echo "Testing batch AST explicit output file creation ..."
	OUTDIR=$(create-tmp-dir "AST_file")
	$ZY -cli -ast -o $OUTDIR/factorial.ast examples/factorial.zsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals 1 $AST_COUNT
	echo

echo "Testing batch Go explicit output file creation ..."
	OUTDIR=$(create-tmp-dir "GO_file")
	$ZY -cli -go -o $OUTDIR/factorial.go examples/factorial.zsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals 1 $GO_COUNT
	echo

echo "Testing batch byte-code explicit output file creation ..."
	OUTDIR=$(create-tmp-dir "Bytecode_file")
	$ZY -cli -bytecode -o $OUTDIR/factorial examples/factorial.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals 1 $BYTECODE_COUNT
	echo

echo "Testing batch AST implicit output file (using dir) creation ..."
	OUTDIR=$(create-tmp-dir "AST_dir")
	$ZY -cli -ast -dir $OUTDIR examples/factorial.zsp
	AST_COUNT=$(count-files $OUTDIR "*.ast")
	num-equals 1 $AST_COUNT
	echo

echo "Testing batch Go implicit output file (using dir) creation ..."
	OUTDIR=$(create-tmp-dir "GO_dir")
	$ZY -cli -go -dir $OUTDIR examples/factorial.zsp
	GO_COUNT=$(count-files $OUTDIR "*.go")
	num-equals 1 $GO_COUNT
	echo

echo "Testing batch byte-code implicit output file (using dir) creation ..."
	OUTDIR=$(create-tmp-dir "Bytecode_dir")
	$ZY -cli -bytecode -dir $OUTDIR examples/factorial.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals 1 $BYTECODE_COUNT
	echo

tear-down
print-summary
