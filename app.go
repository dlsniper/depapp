package main

import "github.com/dlsniper/depone"
import "github.com/dlsniper/deptwo"
import "github.com/dlsniper/depthree"
import "github.com/sirupsen/logrus"

type Helloer interface {
	Hello()
}

func holla(h Helloer) {
	logrus.Println("called holla()")
	h.Hello()
}

func main() {
	mdo := depone.MyDepOne("mdo")
	mdt := deptwo.MyDepTwo("mdt")
	mdtt := depthree.MyDepThree("mdtt")
	holla(mdo)
	holla(mdt)
	holla(mdtt)
}
