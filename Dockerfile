FROM iron/go:dev

WORKDIR /app

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/c0r0nel/api_jwt/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o api_jwt; cp api_jwt /app/

ENTRYPOINT ["./api_jwt"]
