#!/bin/bash

today=`date '+%Y-%m-%d'`
filelog="${today}.log"

if [[ ! -e $filelog ]]; then
	touch $filelog
fi

echo "`date '+%Y-%m-%d %T'` log input start" >> $filelog