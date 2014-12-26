# vplot2png

## Deveopler Setup

	cd $GOPATH
	git clone git@github.com:ricallinson/vplot2png.git ./src/github.com/ricallinson/vplot2png

## Install

	cd ./src/github.com/ricallinson/vplot2png
	go install

## Use

Create a png from the given vplot file.

	vplot2png ./fixtures/test.vplot ./test.png

To change the size of the pen tip use option `-p`.

	vplot2png -p 10 ./fixtures/test.vplot ./test.png
