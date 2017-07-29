#!/bin/sh

s_start=100 ##初めのサイズ
s_end=1000000 ##終わりのサイズ
s_intv=100 ##インクリメントする間隔

## mkdir result/
if [ ! -e result ]; then
mkdir result/
fi


for size in `seq 100 100 1000000`
do
    go run mergesort.go $size | sed -e "s/N= //g" | sed -e :loop -e 'N; $!b loop' -e 's/\n/,/g' >> result/mergesort\_$s_start\-$s_intv\-$s_end.csv
done
