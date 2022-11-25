package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func algOne() {

	//l := &Logger{out: out, prefix: prefix, flag: flag}
	//if out == io.Discard {
	//	l.isDiscard = 1
	//}
	//return l
	l := log.New(os.Stdout, "hello", log.LstdFlags)
	s := fmt.Sprintf("hey %v", l)

	_ = s
	str := "diwakar"
	nr := bytes.NewReader([]byte(str))

	_, _ = str, nr
}

/*
	main() {sum()}

	sum() {
	when work finishes and the function returns the memory would be cleaned automatically (stack is self cleaning)
}

	log() *something {
	this is the large func
 	&something{} -> we are going to refer this out of the current function scope so this can't be allocated on the stack , it has to go on heap
}


*/
