#!/bin/sh
## mkdir result/
if [ ! -e result ];
then
    mkdir result/
else
    rm -r result/
fi
