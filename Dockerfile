FROM python:2-buster

LABEL maintainer="ghchinoy@gmail.com"

# ssl packages
# computer vision packages
# image format packages
RUN apt-get update -y && apt-get upgrade -y && \
    apt-get install -y libcurl4-openssl-dev libssl-dev  \
    python-opencv libopencv-dev  \
    libjpeg-dev libpng-dev libwebp-dev webp

# install thumbor
RUN pip install thumbor

CMD ["thumbor", "--port", "8080"]

EXPOSE 8080 8080