# securl

Generate a secure url for thumbor.

Pronounced "secure-l", `securl` is a portmanteau of "secure" & "url".

# usage

Provide two environment variables:

* `SECURITY_KEY` - required; a string that will be used to HMAC the url parts.
* `THUMBOR_ORIGIN` - optional; the http origin for the secure version of thumbor. if not present, will use the provided URL's origin.

To generate a safe thumbor URL, provide `securl` an unsafe one, like so:

```
$ SECURITY_KEY=globo_rocks
$ THUMBOR_ORIGIN=https://thumbor
$ securl https://thumbor-3u7l6nnjpq-uc.a.run.app/unsafe/x500/filters:grayscale()/https://web.dev/backdrop-filter/hero.jpg
```

securl will then parse out the thumbor unsafe origin and size, filters, and image along with the provided SECURITY_KEY as inputs to create a HMAC (hash-based message authentication code), and then output a safe url, using the THUMBOR_ORIGIN.


# walkthrough

As an example, here's deploying a safe thumbor container on Cloud Run


```
$ gcloud beta run deploy thumbor-safe --image gcr.io/${PROJECT_ID}/thumbor --platform managed --set-env-vars=SECURITY_KEY=${SECURITY_KEY}

Deploying container to Cloud Run service [thumbor-safe] in project [thumbor-container] region [us-central1]
✓ Deploying new service... Done.                                                                       
  ✓ Creating Revision...                                                                               
  ✓ Routing traffic...                                                                                 
  ✓ Setting IAM Policy...                                                                              
Done.                                                                                                  
Service [thumbor-safe] revision [thumbor-safe-rnxtj] has been deployed and is serving 100 percent of traffic at https://thumbor-safe-3u7l6nnjpq-uc.a.run.app
```

Now there's a safe version of thumbor at the above url.
