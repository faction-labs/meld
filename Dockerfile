FROM alpine:latest
COPY build/meld /bin/meld
ADD build/public /bin/public
WORKDIR /bin
ENTRYPOINT ["/bin/meld"]
EXPOSE 8080
