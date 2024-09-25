# syntax=docker/dockerfile:1
 
FROM python:3.11
 
WORKDIR .
 
RUN apt-get -y update && apt-get install lvm2 -y

COPY requirements.txt .
 
RUN pip3 install -r requirements.txt
 
COPY . .
 
EXPOSE 5000
 
ENTRYPOINT ["gunicorn", "main:app"]

