sed 's@[[:space:]]\{1,\}@\n@g' words.txt | sort | uniq -c | sort -k 1 -n -r | awk '{ print $2,$1}'
