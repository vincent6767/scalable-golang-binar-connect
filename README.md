# scalable-golang-binar-connect
The demonstration source code used in the Binar Connect #5.

Quick tutorial to run simpleaccelbytescrapper inside the Docker. 

1. Install docker. You can find the official tutorial [here](https://docs.docker.com/engine/installation/).
2. Go to the **dockerexample** directory.
3. Run `CGO_ENABLED=0 CGOOS=linux go build -o main simpleaccelbytescrapper.go`
4. Run `docker build -t simpleaccelbytescrapper -f Dockerfile.simpleaccelbytescrapper .`
5. Run `docker run -it simpleaccelbytescrapper`

Mutex and Channel performance test can be found [here](https://gist.github.com/cyfdecyf/4562635).
