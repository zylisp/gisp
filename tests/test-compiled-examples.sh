#!/bin/bash

. ./tests/common.sh

echo "Testing batch byte-code multiple file creation ..."
	OUTDIR=$(create-tmp-dir "Bytecodes_dir")
	$ZY -cli -bytecode -dir $OUTDIR examples/*.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $EXAMPLES_COUNT $BYTECODE_COUNT

	# Let's not test the slow ones; just the ones that return
	# more or less quickly:
	for FILE in $OUTDIR/factorial $OUTDIR/even_fib_terms $OUTDIR/power_digit_sum $OUTDIR/multiples_of_3_5; do
		echo "$(basename $FILE): `$FILE`"
		exit-zero $? "Compiled file did not execute properly"
	done
	echo

tear-down
print-summary
