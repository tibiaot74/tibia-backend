FROM python:3.10.7-alpine3.16
WORKDIR /tests
COPY requirements.txt .
RUN pip3 install -r requirements.txt
WORKDIR /tests/features
COPY ./features/ .
WORKDIR /tests
CMD ["behave"]