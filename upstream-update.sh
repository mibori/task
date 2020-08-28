#!/bin/sh

git pull upstream master
git rebase upstream/master

echo If you can see this message, it has been finished without conflicts
