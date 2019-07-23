#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/revel" -importPath github.com/dtmkeng/framework-exmaple/revel -srcPath "$SCRIPTPATH/src" -runMode dev
