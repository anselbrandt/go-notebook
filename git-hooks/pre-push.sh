#!/bin/sh

docker build -t anselbrandt/go-notes:latest .
docker push anselbrandt/go-notes:latest
ssh root@anselbrandt.dev << HERE
docker pull anselbrandt/go-notes:latest
docker tag anselbrandt/go-notes:latest dokku/go-notes
dokku tags:deploy go-notes
docker system prune -a
y
HERE

exit 0