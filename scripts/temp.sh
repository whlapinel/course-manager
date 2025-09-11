#!/bin/bash

mv hugosites/lapinel-6/public hugosites/lapinel-6/docs
cd hugosites/lapinel-6
touch docs/CNAME && echo "python-1.info" >docs/CNAME
git add .
git commit -m "auto-commit"
git push
mv docs public
