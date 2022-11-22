
# Wordament Solver

This tool is a solver for the popular game wordament

Pre-Req
========
1. Go 1.18 (with generics needed). Otherwise tests don't work
2. Docker installed
3. A linux or WSL setup
4. Local dev/test works on Windows

Local Build Test
================
Local Build using the following in either the wordament-cli or service folder
``` bash
go build .
```

Local Run using direct console app
``` bash
wordament-cli ZRFLPFUALINXAYEM
```

To test the service run
``` bash
service.exe -port=80
```

From another console run
``` bash
curl localhost/?input=SPAVURNYGERSMSBE -v
```

To run the go tests
``` bash
go test ./… -v
```

Web-Service
============
Build
------
This needs to be run inside a linux environment, my tool of choice is WSL 2

Get into WSL 2 on Windows machine. Ensure the pre-reqs mentioned above is present

``` bash
cd ~/github/wordament/service
./build.sh
docker push bonggeek/wordament
```

Deploy to VM
------------
SSH into your Linux VM
``` bash
ssh <user>@server -i MyKey.cer
```

The either run the deployment script
``` bash
sudo bash
curl -fsSL "https://raw.githubusercontent.com/abhinababasu/wordament/main/service/deploy.sh?token=GHSAT0AAAAAAB26S7FFQJA7APXSWXXJW7U6Y32Y2SA" | bash
```

Or manually
``` bash
sudo bash
docker ps --filter="ancestor=wordament:0.1" -q | xargs docker stop
docker pull bonggeek/wordament
docker run -d --restart="always" -p 8090:8090 bonggeek/wordament
```

Test Remote
------------
Run Against Remote
``` bash
curl commonvm1.westus2.cloudapp.azure.com:8090/?input=SPAVURNYGERSMSBE
```

Misc
=====
While there is many more optimization that can be done, at this point the service meets my basic requirement. On a 2 vCPU Azure VM, the service loads the dictionaries in <100ms and solves a wordament game in ~1ms 

TODOs / Known Issues
====================
1. No multiple letter cell support
1. Does get duplicate words
1. Result is sorted by length only
1. Word found is as good as the word list I have (which is not very good)