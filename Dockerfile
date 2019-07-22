FROM golang:1.11
ADD . .
CMD './start.sh'
EXPOSE 9000
