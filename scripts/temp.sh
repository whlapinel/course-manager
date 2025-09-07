#!/bin/bash

cp -r hugosites/lapinel-6/public/. hugosites/lapinel-6/docs
cd hugosites/lapinel-6
git add .
git commit -m "auto-commit"
git push
rm -r hugosites/lapinel-6/docs
