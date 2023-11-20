FROM golang:1.21-alpine

ENV GOSRC=${GOPATH}/src

WORKDIR ${GOSRC}/miniproject

COPY . .

RUN go mod tidy
RUN go install

CMD [ "miniproject" ]
