#!/bin/bash

basedir="/tmp/spycode"

for filename in $(find $HOME/ispycode.com/GO -type f); do

    newname=$(echo ${filename##*/} | sed 's/.html/.go/');

    foldername="$(echo ${filename%/*} | sed 's,$HOME/ispycode.com/GO,,g')"
    if ! test -d "$basedir/$foldername"; then
        mkdir -p "$basedir/$foldername"
    fi

    cat "$filename" |\
        sed -e 's/<[^>]\+>//g' "$filename" |\
        sed -n '/package main/,/Output:/p' |\
        sed 's/\&nbsp;//g'         |\
        sed 's/\&quot;/\"/g'       |\
        sed 's/\&lt;/\</g'         |\
        sed 's/\&amp;/\&/g'        |\
        sed 's/\&gt;/\</g'         |\
        grep -v "Output:"           \
            > "$basedir/$foldername/$newname" &&\
            if [[ $(stat --printf=%s  $basedir/$foldername/$newname) -eq 0 ]]; then
                rm "$basedir/$foldername/$newname"
            fi
done

cd $basedir && go fmt ./...