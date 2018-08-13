FROM ubuntu:16.04

EXPOSE 9081

WORKDIR /home/msb2pilot
COPY ./conf ./conf
COPY ./msb2pilot .

ENTRYPOINT /home/msb2pilot/msb2pilot
