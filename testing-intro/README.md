Commands to test manually with curl:
```bash
# get homepage
curl http://localhost:8000

# list all articles
curl http://localhost:8000/articles

# list specific article (here 3)
curl http://localhost:8000/article/3

# add new article
curl -X POST -d \
'{"Id":"5","Title":"ulno-net","author":"ulno","link":"https://www.ulno.net"}' \
http://localhost:8000/article

# delete article with index 2
curl -X DELETE http://localhost:8000/article/2
```

Run tests with `go test ./... -count=1`