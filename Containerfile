FROM quay.io/centos/centos:stream9

RUN dnf -y install vim

CMD ["vim", "--help"]

