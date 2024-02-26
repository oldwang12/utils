default:
	git add .
	git commit -m "`date '+%Y/%m/%d %H:%M:%S'`"
	git push origin master
	git tag -a "$(tag) -m "`date '+%Y/%m/%d %H:%M:%S'`"
	git push origin "$(tag)"
