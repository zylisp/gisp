FROM scratch
MAINTAINER ZYLISP

COPY bin/zylisp-linux /bin/zylisp
ENTRYPOINT ["zylisp", "-ast"]