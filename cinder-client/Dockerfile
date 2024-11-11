# syntax=docker/dockerfile:1
 
FROM python:3.8 
 
 
RUN apt-get -y update && apt-get install git curl wget lvm2 vim -y
RUN git clone https://github.com/openstack/python-cinderclient
WORKDIR "/python-cinderclient"
RUN pip install -e .
RUN pip install python-brick-cinderclient-ext  

WORKDIR . 
COPY . .
RUN pip install -r requirements.txt
 
EXPOSE 5000
 
ENTRYPOINT ["gunicorn", "main:app"]

