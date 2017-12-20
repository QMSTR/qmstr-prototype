FROM jenkins/jenkins:lts
USER root

RUN apt-get update && \
    apt-get install -y autoconf git libio-captureoutput-perl \
    libtool cmake wget tar build-essential
#curl libz-dev libssl-dev

RUN cd /tmp && wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz && tar -xvf go1.9.2.linux-amd64.tar.gz && mv go/ /usr/local

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# install ninka
RUN mkdir /ninka && \
    git clone https://github.com/dmgerman/ninka.git /ninka && \
    cd /ninka/comments && make && make install && \
    rm /usr/local/man/man1 && \
    cd /ninka && perl Makefile.PL && make && make install && \
    rm -fr /ninka

# setup qmstr build env
RUN mkdir -p /go/src/qmstr-prototype && \
    mkdir /qmstr && \
    git clone https://github.com/QMSTR/qmstr-prototype.git /qmstr && \
    ln -s /qmstr/qmstr /go/src/qmstr-prototype/qmstr

# build and install qmstr
RUN go get qmstr-prototype/qmstr/qmstr-master && \
    go install qmstr-prototype/qmstr/qmstr-master && \
    go install qmstr-prototype/qmstr/qmstr-wrapper && \
    rm -fr /qmstr/*
RUN mkdir /qmstr-wrapper && \
    for i in gcc g++; do ln -s /go/bin/qmstr-wrapper /qmstr-wrapper/${i}; done

USER jenkins
