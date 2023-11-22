FROM golang:1.21.3-alpine3.18

ENV GOSRC=${GOPATH}/src

WORKDIR ${GOSRC}/imgsimk

COPY . .

RUN go mod tidy
RUN go install

CMD [ "imgsimk" ]
