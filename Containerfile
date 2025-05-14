FROM registry.access.redhat.com/ubi7:latest

RUN yum -y install vim

CMD ["vim", "--help"]
