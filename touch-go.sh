find . -name '*.go' | while read f;do echo "// $(date)" >> $f;N=$((N + 1));if [[ $N -gt 1 ]];then break;fi done
