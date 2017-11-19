FROM golang:1

ADD . /go/src/github.com/tobyjsullivan/resume-api
RUN  go install github.com/tobyjsullivan/resume-api/companies-db
CMD /go/bin/companies-db
