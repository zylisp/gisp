FROM scratch

COPY bin/zylisp-linux /bin/zylisp
ENTRYPOINT ["zylisp"]