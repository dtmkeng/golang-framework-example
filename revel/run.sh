#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/revel" -importPath github.com/dtmkeng/golang-framework-example/revel -srcPath "$SCRIPTPATH/src" -runMode dev
