#!/bin/sh

for i in $(seq 1 100)
do
    grep w:uiPriority='"'$i'"' styles.xml >> demo.xml
done
