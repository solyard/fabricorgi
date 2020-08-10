FROM hyperledger/fabric-tools:2.0 as main
COPY fabricorgi /usr/bin/fabricorgi
ENTRYPOINT [ "fabricorgi" ]