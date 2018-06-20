#!/bin/bash
#menu.sh

title="My Menu"
url="www.davidnotbad.com"
time=`date +%Y-%m-%d`

function menu()
{
	clear
cat << eof
###########################################################
			`echo -e "\033[32m$title\033[0m"`
###########################################################
`echo -e "\033[31m*\033[0m"`	0)quit
`echo -e "\033[31m*\033[0m"`	1)return to main menu
`echo -e "\033[31m*\033[0m"`	2)add a user
`echo -e "\033[31m*\033[0m"`	3)set password for user
`echo -e "\033[31m*\033[0m"`	4)delete a user
`echo -e "\033[31m*\033[0m"`	5)print disk space
`echo -e "\033[31m*\033[0m"`	6)print mem space
###########################################################
$url			$time
###########################################################
eof
	echo
}

