The getpad program takes an itsapad URL as an argument and prints the raw text of
the paste on standard output. For example:

	% getpad http://itsapad.appspot.com/348001
	Hello, github

This is useful in a Plan 9 plumber rule such as:

	# open paste urls in edit
	type is text
	data matches 'http://itsapad.appspot.com/([0-9]+)(/[a-z]+)?'
	plumb start rc -c 'getpad '$0' | plumb -i -d edit -a ''action=showdata filename=/paste/'$1''''
