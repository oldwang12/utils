default:
	git add .
	git commit -m "`date '+%Y/%m/%d %H:%M:%S'`"
	git push origin master
	git tag -a "$(tag) -m "no log"
	git push origin "$(tag)"
