# go
Go (Golang) Libraries

Stream:
Downstream has a field that contains an openUpstream() function
Before openUpstream() is called, no goroutines or channels are created.
Therefore, there is no resource leakage if only intermediate operations are called.
The openUpstream() function is not thread-safe --- it must not be called from multiple goroutines concurrently.
