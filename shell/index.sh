#!/bin/bash
#index.sh

. menu.sh

menu

while true; do
	read -p "please input a option: " option
	case $option in
		0 )
			echo -e "\033[32mquit successfully\033[0m"
			echo
			break
			;;
		1 )
			menu
			;;
		2 )
			read -p "add a user: " name
			useradd $name &>/dev/null

			if [[ $? -eq 0 ]]; then
				echo -e "\033[32muser ${name} created successfully\033[0m"
			else
				echo -e "\033[31muser ${name} created failly\033[0m"
			fi
			;;
		3 )
			read -p "input the user: " name
			read -p "set pass for ${name}: " pass
			echo $pass | passwd --stdin $name &>/dev/null

			if [[ $? -eq 0 ]]; then
				echo -e "\033[32msuccessful\033[0m"
			else
				echo -e "\033[31mfail\033[0m"
			fi
			;;
		4 )
			read -p "delete a user: " name
			userdel -r $name &>/dev/null

			if [[ $? -eq 0 ]]; then
				echo -e "\033[32msuccessful\033[0m"
			else
				echo -e "\033[31mfail\033[0m"
			fi
			;;
		5 )
			df -Th
			;;
		6 )
			free -m
			;;
	
	esac
	echo
done