set wildignore+=src/github.com/*
set wildignore+=src/bitbucket.org/*
set wildignore+=src/launchpad.net/*
set wildignore+=src/code.google.com/*

fu! Shell(command)
	
	let filePath = a:command . "/.restart" 
	execute "! touch " . filePath

endfunction 

map <Leader>r :call Shell($PACKAGE)
map <Leader>rr :call Shell($PACKAGE)



