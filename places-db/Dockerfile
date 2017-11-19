FROM golang:1

ADD . /go/src/github.com/tobyjsullivan/resume-api
RUN  go install github.com/tobyjsullivan/resume-api/places-db
CMD /go/bin/places-db
