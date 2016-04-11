#!/bin/bash
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  go build -v -o ${distdir}/mr-burns
}

go_build