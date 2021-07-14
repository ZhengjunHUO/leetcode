# 判断文件是否有10行
if [ $(cat file.txt | wc -l) -lt 10 ]; then echo "target file less than 10 lines!"; exit 0; fi

# 第一种方法
awk 'NR==10' file.txt

# 第二种方法
sed -n 10p file.txt 

# 第三种方法
tail +10 file.txt | head -1

# 第四种方法
head -10 file.txt | tail -1

# 第五种方法
cnt=1
while read LINE; do
	if [ $cnt -eq 10 ]; then
		echo $LINE
		break
	fi

	cnt=$((cnt+1))
done < file.txt 
