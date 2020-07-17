run:
	buffalo dev

action:
	buffalo g actions ${root} ${sub}  --skip-template
	echo generated route: ${root}/${sub}