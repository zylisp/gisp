#!/bin/bash

. ./tests/common.sh

echo "Testing batch byte-code multiple file creation for zyc ..."
	OUTDIR=$(create-tmp-dir "Bytecodes_dir")
	$ZYC -dir $OUTDIR examples/*.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $EXAMPLES_COUNT $BYTECODE_COUNT

	# Let's not test the slow ones; just the ones that return
	# more or less quickly:
	for FILE in $OUTDIR/factorial $OUTDIR/even_fib_terms $OUTDIR/power_digit_sum $OUTDIR/multiples_of_3_5; do
		echo "$(basename $FILE): `$FILE`"
		exit-zero $? "Compiled file did not execute properly"
	done
	echo

echo "Testing batch byte-code explicit output file creation for zyc ..."
	OUTDIR=$(create-tmp-dir "Bytecode_file")
	$ZYC -o $OUTDIR/factorial examples/factorial.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals 1 $BYTECODE_COUNT
	echo

echo "Testing batch byte-code implicit output file (using dir) creation for zyc ..."
	OUTDIR=$(create-tmp-dir "Bytecode_dir")
	$ZYC -dir $OUTDIR examples/factorial.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals 1 $BYTECODE_COUNT
	echo

tear-down
print-summary
