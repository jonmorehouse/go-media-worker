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
-	2.) Encoder -- should have various functions that can encode a video and export the results to the correct location etc
-	3.) Upload -- should upload the element to an s3 key as specified 
-	4.) Utilities -- utilities for various pieces of application

Setup Commands
-

-	`go get -u launchpad.net/gocheck`




