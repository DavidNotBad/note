#!/bin/bash

if [[ -f ./file1 ]]; then
	echo "file1 is exists!"
else
	touch file1
fi