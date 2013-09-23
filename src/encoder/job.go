package encoder

type Job {

	var FilePath, OutputPath, EncodingType string

	// whether 
	var Finished bool

	var Channel chan byte[]
}

//func (self *EncodingJob) New () *EncodingJob {




//}

