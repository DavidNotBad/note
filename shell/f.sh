#!/bin/bash
tot=0

for i in $*; do
	tot=$(($tot+$i))
done

echo $tot