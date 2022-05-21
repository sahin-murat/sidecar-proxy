### Sidecar Proxy ###
Implementation of sidecar pattern

##### To Run the application
```

//Developmennt
docker compose up

//Production - Build
 docker build -t sidecar-proxy-production . --target production
 
 //Production - Run
 docker run -p 80:3333 --name sidecar-proxy-production sidecar-proxy-production

```


##### To Run the test

```

go test ./...

```