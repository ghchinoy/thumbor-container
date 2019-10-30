# thumbor-container

[thumbor](http://thumbor.org/) in a Docker container.

Inspired by the [web.dev](https://web.dev/blog) blog post _[How to install the Thumbor image CDN](https://web.dev/install-thumbor/)_ by Katie Hempenius, this details both the process to put thumbor in a container (even though there are better containers out there, such as [minimalcompact/thumbor](https://hub.docker.com/r/minimalcompact/thumbor)) and how to use Google Cloud Platform to build and host thumbor.

## container usage

Use the container directly as published at [ghchinoy/thumbor](https://hub.docker.com/r/ghchinoy/thumbor) or build and deploy it as per below.

```
docker run -d -p 8080:8080 -e ALLOW_UNSAFE_URL=True ghchinoy/thumbor
```

If you want to use the HMAC security, supply your own SECRET_KEY environment variable, sort of like this (and remember it, since you'll have to use that to encode your URLs):

```
SECURITY_KEY=$(head /dev/urandom | tr -dc A-Za-z0-9 | head -c 13)
docker run -d -p 8080:8080 -e SECURITY_KEY=${SECURITY_KEY} ghchinoy/thumbor
```

## building

Use the included Dockerfile to build the container in the standard manner:

```
docker build -t thumbor .
```

Or, use the Google Cloud Platform Cloud Build service to build the container for you. The [first 120 build-minutes per day are free](https://cloud.google.com/cloud-build/pricing), even though you have to enable billing for your GCP account. See also the [container registry](https://cloud.google.com/container-registry/pricing) pricing for storage of and security scanning of containers.

Use the included cloudbuild.yaml which uses an environment variable to specify your project, i.e. do something like this first: `export PROJECT_ID=thumbor-container` (substituting your own project name, of course).

View build logs at your project name, ex. [thubmor-container](https://console.cloud.google.com/cloud-build/builds?project=thumbor-container)

```
# prerequisites
export PROJECT_ID=thumbor-container
# create a gcp project
gcloud projects create $PROJECT_ID
# enable billing, find via `gcloud alpha billing account list`
gcloud alpha billing projects link $PROJECT_ID --billing-account YOUR-ACCT-NUMB
# enable cloud build service
gcloud services enable cloudbuild.googleapis.com

# build step
gcloud builds submit --config cloudbuild.yaml .
```

## deploy

One can run the container from Docker hub registry, as above, or use Cloud Run to serve up the container! 

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

To use Cloud Run from the command line, choose a region and allow unauthenticated access, when asked.

```
# enable the API
gcloud services enable run.googleapis.com

# deploy to cloud run
gcloud beta run deploy --image gcr.io/$PROJECT_ID/thumbor --platform managed --set-env-vars=ALLOW_UNSAFE_URL=True
```

After this is deployed, you'll see something like the following.

```
Deploying container to Cloud Run service [thumbor] in project [thumbor-container] region [us-central1]
✓ Deploying... Done.                                                                                   
  ✓ Creating Revision...                                                                               
  ✓ Routing traffic...                                                                                 
Done.                                                                                                  
Service [thumbor] revision [thumbor-xh4h4] has been deployed and is serving 100 percent of traffic at https://thumbor-3u7t5nnjpq-uc.a.run.app
```

Test with URLs in the blog posts _[Optimize images with Thumbor](https://web.dev/use-thumbor/)_ or _[How to install the Thumbor image CDN](https://web.dev/install-thumbor/)_.