# Fetcher
Web pages fetching and saving them to disk for later retrieval and browsing

# Run application

Build docker image

```
$ docker build -t fetcher .    
```

And then run docker image with next params

```
$ docker run -v $(pwd)/web:/app/web fetcher --metadata --assets https://google.com https://autify.com
```

Passing `-v $(pwd)/web:/app/web` mounts your local web directory to docker container's one, where all fetched files saved.

Pass `--metadata` if you want to see additional information.

Pass `--assets` if you want to load additional assets (css, js and images).


 
