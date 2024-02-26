default:
	git add .
	git commit -m "`date '+%Y/%m/%d %H:%M:%S'`"
	git push origin master
	git tag -a v1.0.9 -m "no log"
	git push origin v1.0.9
