# Web Info
This web application gives the information about a webpage

## Running this application
```
docker build -t "server:web-info" .

docker run --publish 3001:3001 -i -t "server:web-info"
```

### Running test cases

```
go test -race -cover -v -mod=vendor ./...   
```

## API
1. `GET /summary`

### Sample Request 
```json 
		{
			"url": "https://www.home24.com/websites/homevierundzwanzig/English/0/home24.html"
		}
```

### Sample Response
```json
	{
	    "htmlVersion": "HTML 5",
	    "title": "home24 | The online destination for home and living.",
	    "headings": {
	        "heading1Count": 0,
	        "heading2Count": 0,
	        "heading3Count": 3,
	        "heading4Count": 4,
	        "heading5Count": 0,
	        "heading6Count": 0
	    },
	    "internalLinksCount": 13,
	    "externalLinks": 2,
        "inAccessibleLinks": 0,
        "isLogIn": false
    }
```

Curl request:
```
curl -X GET \
  -H "Content-type: application/json" \
  -H "Accept: application/json" \
  -d '{"url": "https://www.home24.com/websites/homevierundzwanzig/English/0/home24.html"}' \
  "http://localhost:3001/summary"
```

2. `GET /info`

### Sample Request 
```json 
		{
			"url": "https://www.home24.com/websites/homevierundzwanzig/English/0/home24.html"
		}
```

### Sample Response
```json
	{
	    "htmlVersion": "HTML 5",
	    "title": "home24 | The online destination for home and living.",
	    "headings": {
	        "heading1": null,
	        "heading2": null,
	        "heading3": null,
	        "heading4": [
	            "get in touch with us",
	            "Home24",
	            "International Shops",
	            "Get in Touch with us"
	        ]
	    },
	    "internalLinks": [
	    "https://www.home24.com/websites/homevierundzwanzig/English/1/homepage.html",
        "https://www.home24.com/websites/homevierundzwanzig/English/7000/contact.html",
        "https://www.home24.com/websites/homevierundzwanzig/German/0/home24.html"
    	],
	    "externalLinks": [
		"https://twitter.com/home24_de",
		"https://www.facebook.com/home24.de",
		"https://www.linkedin.com/company/home24/",
		"https://www.instagram.com/home24_de/"
        ],
        "inAccessibleLinks": [],
        "isLogIn": false
    }
```

```
curl -X GET \
  -H "Content-type: application/json" \
  -H "Accept: application/json" \
  -d '{"url": "https://www.home24.com/websites/homevierundzwanzig/English/0/home24.html"}' \
  "http://localhost:3001/info"

````

### Application architecture

The application crawls the information for the given URL. Multiple go routines are spun to check the accessiblity of the links.