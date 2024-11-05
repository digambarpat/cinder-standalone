# syntax=docker/dockerfile:1

FROM python:3.8


RUN apt-get -y update && apt-get install -y git curl wget lvm2 tgt
COPY --chmod=0755 tgtd.sh /tgtd.sh

CMD ["/tgtd.sh"]

WORKDIR .
COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

EXPOSE 5000

ENTRYPOINT ["gunicorn", "main:app"]