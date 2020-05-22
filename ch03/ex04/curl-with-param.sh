#!/bin/sh
cd `dirname $0`

# parameter = width / height / color
# color = 文字列 red or blue or etc
curl 'http://localhost:8000/svg' > 1default.svg
curl 'http://localhost:8000/svg?color=purple' > 1colorparam.svg
curl 'http://localhost:8000/svg?width=1280&height=640' > 1whparam.svg
open .
