set wildignore+=src/github.com/*
set wildignore+=src/bitbucket.org/*
set wildignore+=src/launchpad.net/*
set wildignore+=src/code.google.com/*

fu! Shell(command)
	
	execute "! " . a:command

endfunction 

"map <Leader>r :call Shell($PACKAGE)
"map <Leader>rr :call Shell($PACKAGE)
nunmap <Leader>r
map <Leader>r :call Shell("printf \"\033c\" && go test")<CR>

