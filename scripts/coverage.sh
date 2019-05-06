#!/bin/bash

go get github.com/go-playground/overalls && go get github.com/mattn/goveralls

overalls -project=github.com/deelawn/convert -covermode=count
goveralls -coverprofile=overalls.coverprofile -service travis-ci