run:
	buffalo dev

action:
	echo r= root directory s= subdirectory m= method
	echo example: make action r=auth s=signup m=POST
	buffalo g actions ${r} ${s} --skip-template --method ${m}
	echo generated ${r}/${s} Method: ${m}