Go Media Servers
=

Overview
=

-	Media worker servers are generally managed via our python media-workers cluster etc
-	More optimized tasks are built in go and go here. These should still connect to the same gearman servers, but we can scale these servers a bit easier as a seperate component

Application structure
-

-	1.) Bootstrap -- should load up json config and parse environment properly... should be importable etc
-	2.) Worker -- should be able to register various tasks and call the correct shared modules with correct settings etc -- should be able to parse json request etc
-	3.) Tests -- should have unit tests, integration tests, stubs, factories etc



