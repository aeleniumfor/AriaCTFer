FROM golang

WORKDIR /go/src/AriaCTFer
COPY ./ ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN go install -v ./
CMD ["AriaCTFer"]