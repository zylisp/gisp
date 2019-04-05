#!/bin/bash

. ./tests/common.sh

echo "Batch byte-code multiple file creation for zyc"
	OUTDIR=$(create-tmp-dir "Bytecodes_dir")
	zyc -dir $OUTDIR examples/*.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $BYTECODE_COUNT $EXAMPLES_COUNT

	# Let's not test the slow ones; just the ones that return
	# more or less quickly:
	for FILE in $OUTDIR/factorial $OUTDIR/even_fib_terms $OUTDIR/power_digit_sum $OUTDIR/multiples_of_3_5; do
		echo "$(basename $FILE): `$FILE`"
		if [ $? -eq 0 ]; then
			echo -e "  ${GREEN}PASS${CLEAR_COLOR}"
			PASSES=$((PASSES+1))
		else
			echo -e "  ${RED}FAIL${CLEAR_COLOR}: Compiled file did not execute properly"
			FAIL_SUITE=1
			FAILURES=$((FAILURES+1))
		fi
	done
	echo

echo "Batch byte-code explicit output file creation for zyc"
	OUTDIR=$(create-tmp-dir "Bytecode_file")
	zyc -o $OUTDIR/factorial examples/factorial.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $BYTECODE_COUNT 1
	echo

echo "Batch byte-code implicit output file (using dir) creation for zyc"
	OUTDIR=$(create-tmp-dir "Bytecode_dir")
	zyc -dir $OUTDIR examples/factorial.zsp
	BYTECODE_COUNT=$(count-files-without-extension $OUTDIR)
	num-equals $BYTECODE_COUNT 1
	echo

echo -e "Tests passed: ${GREEN}$PASSES${CLEAR_COLOR}"
echo -e "Tests failed: ${RED}$FAILURES${CLEAR_COLOR}"
echo

if [ ! -z "$BASE_OUTDIR" ]; then
	rm -rf $BASE_OUTDIR
fi

if [ "$FAIL_SUITE" == 1 ]; then
	exit 1
fi
