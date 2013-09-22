Go Media Servers
=

Overview
=

-	Media worker servers are generally managed via our python media-workers cluster etc
-	More optimized tasks are built in go and go here. These should still connect to the same gearman servers, but we can scale these servers a bit easier as a seperate component

Notes
-

-	To enable vim syntax for golang: `cp $GOROOT/misc/vim/go.vim ~/.vim/syntax`



Application structure
-

-	1.) Bootstrap -- should load up the correct configuration for the application
	
	-	config
	-	error handlers -- both application wide and booting 
	-	gearman handlers -- create client / workers etc ... boot workers as well?

-	2.) Encoder -- should have various functions that can encode a video and export the results to the correct location etc
	
	-	initialize an encoder function that can take in a filePath and encode the file properly
	
-	3.) Upload -- should upload the element to an s3 key as specified 

	-	upload an element to s3 given correct parameters and settings etc
	-	should use multipart upload if relevant

-	4.) Utilities -- utilities for various pieces of application

	-	file naming utilities
-	5.) App -- should control the various tasks and create a master task runner that will run on several cores in parallel 
	
	-	should read multiple parameters to determine how many cores and processes can efficiently be run 

Testing Environments 
-

-	`gowatch -test ` (run inside of a directory to continually test application)
-	need to create `, r` commands within local.vimrc to trigger file changes etc 
-	Note: when running `go test bootstrap` there is no output. However, running `go test` from within the bootstrap package directory, it will diesplay output










