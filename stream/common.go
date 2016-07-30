package stream

func prepareOutChannels() (out, outBackup chan<- interface{}, outForReturn <-chan interface{}) {
	outCommon := make(chan interface{})
	return nil, outCommon, outCommon
}

func prepareIntOutChannels() (out, outBackup chan<- int, outForReturn <-chan int) {
	outCommon := make(chan int)
	return nil, outCommon, outCommon
}
