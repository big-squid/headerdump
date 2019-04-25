FROM 864745837241.dkr.ecr.us-west-2.amazonaws.com/kraken-api:latest
ADD . .
CMD './start.sh'
EXPOSE 9000
