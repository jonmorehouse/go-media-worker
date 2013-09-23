Gearman Encoding Worker Binary
=

High Level
-

-	1.) Should initialize workers with the correct functions linked up and run them concurrently
-	2.) Each gearman-worker should be its own go-routine
-	3.) Each worker should be callable and should return its own go-routine etc
-	4.) Upon calling the app.Worker() it should initialize and shoot off its own go-routine

Questions
-

-	1.) How do we kill this goroutine in testing?
	
	-	"runtime"


