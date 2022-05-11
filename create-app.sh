#!/bin/sh

docker build -t anselbrandt/go-notes:latest .
docker push anselbrandt/go-notes:latest
ssh root@anselbrandt.dev << HERE
dokku apps:create go-notes
dokku domains:set go-notes anselbrandt.dev
dokku proxy:ports-set go-notes http:80:8080
dokku certs:add go-notes < cert-key.tar
docker pull anselbrandt/go-notes:latest
docker tag anselbrandt/go-notes:latest dokku/go-notes
dokku tags:deploy go-notes
docker system prune -a
y
HERE

exit 0