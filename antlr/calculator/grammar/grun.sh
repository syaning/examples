#!/bin/bash

alias antlr4-parse='java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'
antlr4-parse *.g4 prog -gui
