FROM golang:1.22
ENV INPUT="./csv/hb_test.csv"
WORKDIR /app
COPY . .
RUN go build

RUN ./clialgotool -input ${INPUT} -print
