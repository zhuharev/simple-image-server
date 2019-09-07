# Simple image upload and store server

This server can handle upload images, resize to max allowed size and store on local filesystem. All images converted to jpg.

# Usage 

`simple-image-server`

```curl
curl \
  -F "file=@/path/to/image/image.jpg" \
  localhost:8080/upload
```

You will get receive result such as: `{"url":"https://example.com/image.jpg"}`, now you can access this url from browser.

# Endpoints

`/upload` - send to it images as multipart or just slice of bytes. Result in json sample: `{"url":"https://example.com/image.jpg"}`