#!/bin/bash

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac

branch_name=$(git branch | sed -n -e 's/^\* \(.*\)/\1/p')

if [ "$machine" == "Mac" ]; then

    printf "%b" "\e[1;34m Staging to git... \e[0m\n"

    git add .

    read -r -p "Commit Message: " commit_msg 

    git commit -m "${commit_msg}"

    git push origin $branch_name

    printf "%b" "\e[1;32m Finished pushing to github! \e[0m\n"
else
    echo -e "\e[1;36m Staging to git... \e[0m"

    git add .

    read -r -p "Commit Message: " commit_msg 

    git commit -m "${commit_msg}"

    git push origin $branch_name

    echo -e "\e[1;32m Finished pushing to github! \e[0m"
fi

