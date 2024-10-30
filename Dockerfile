# syntax=docker/dockerfile:1
 
FROM python:3.8 
 
 
RUN apt-get -y update && apt-get install git curl wget lvm2 vim -y
RUN git clone https://github.com/openstack/python-cinderclient
WORKDIR "/python-cinderclient"
RUN pip install -e .
RUN pip install python-brick-cinderclient-ext  

WORKDIR .
RUN wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin

WORKDIR .
COPY requirements.txt .
 
RUN pip install -r requirements.txt
 
COPY . .
 
EXPOSE 5000
 
ENTRYPOINT ["gunicorn", "main:app"]

