FROM golang
ADD . /go/src/github.com/detroitcybersec/cryptexly
WORKDIR /go/src/github.com/detroitcybersec/cryptexly/cryptexlyd
COPY cryptexlyd/sample_config.json config.json
RUN make
EXPOSE 8080
ENTRYPOINT ./cryptexlyd -config config.json -app ../cryptexly
