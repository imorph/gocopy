#!/bin/bash

mkdir -p tmp

fail() {
    echo "Result => FAIL"
    exit 1
}

ok() {
    echo "Result => OK"
}

echo "Test #01: Input file not exist."
./gocopy -from tmp/no_file -to /tmp/no_file > /dev/null 2>&1 && fail || ok

echo "Test #02: No input file in arguments."
./gocopy -to /tmp/no_file > /dev/null 2>&1 && fail || ok

echo "Test #03: No output file in arguments."
./gocopy -from tmp/no_file > /dev/null 2>&1 && fail || ok

echo "Test #04: Offset + limit is more than input file size"
dd if=/dev/zero of=tmp/from_file04 bs=100 count=1  > /dev/null 2>&1 && \
./gocopy -from tmp/from_file04 -to tmp/to_file04 -offset 234234 -limit 12342345 > /dev/null 2>&1 && fail || ok

echo "Test #05: Normal copy with zero offset"
dd if=/dev/zero of=tmp/from_file05 bs=100 count=1  > /dev/null 2>&1 && \
./gocopy -from tmp/from_file05 -to tmp/to_file05 -offset 0 -limit 100 > /dev/null 2>&1 && \
test $(du -b tmp/to_file05 | awk '{print $1}') -eq 100 && ok || fail

echo "Test #06: Normal copy with non-zero offset"
dd if=/dev/zero of=tmp/from_file06 bs=220 count=1  > /dev/null 2>&1 && \
./gocopy -from tmp/from_file06 -to tmp/to_file06 -offset 100 -limit 120 > /dev/null 2>&1 && \
test $(du -b tmp/to_file06 | awk '{print $1}') -eq 120 && ok || fail


rm -rf ./tmp
