FROM ubuntu

ENV PATH /go/bin:/usr/lib/go-1.9/bin:$PATH
ENV GOPATH /go
ENV QMSTR_BUILD_DIR /build

# install golang 1.9
RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository ppa:gophers/archive && \
    apt-get update && \
    apt-get install -y golang-1.9-go autoconf git libio-captureoutput-perl \
    libtool cmake curl libz-dev libssl-dev

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
VOLUME [ "/qmstr" ]

# setup build dir
RUN mkdir -p ${QMSTR_BUILD_DIR}
VOLUME [ "${QMSTR_BUILD_DIR}" ]
RUN mkdir /qmstr-wrapper && \
    for i in gcc g++; do ln -s /go/bin/qmstr-wrapper /qmstr-wrapper/${i}; done

WORKDIR ${QMSTR_BUILD_DIR}

COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY helper-funcs.sh /helper-funcs.sh
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT [ "/docker-entrypoint.sh" ]
