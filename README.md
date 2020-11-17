# FABRICORGI

![Image](fabricorgi.png)

What this tool can do?

- Firstly this tool created for comfortable managing organizations list in your Hyperledger network
- Secondly, this is one of the ways to learn GoLang (for me)
- Third is just a try to make some useful open-source tool for all

---

## HOW TO BUILD

Set the ENV variable CGO_ENABLED to 0 on your system via this command:

```
export CGO_ENABLED=0
```

Then run command in root dir of project:
```
go build cmd/fabricorgi/fabricorgi.go
```

Then run docker build command to build image with fabricorgi:
```
docker build . -f Dockerfile -t fabricorgi:latest
```

## HOW TO USE

Use the output json file produced by `configtxgen` tool and send it into API endpoint depend of your task.
For adding organisation use:
```
example.com:8081/api/v1/addorg/{channel}
```
Where `channel` it's a channel inside that you want an organization to be added

For remove organisation use:
```
example.com:8081/api/v1/removeorg
```

As body for remove organisation method use this structure:
```
{
    OrgName: example1MSP
}
```

---

## ADDITIONAL INFO

I used hyperledger/fabric-tools:2.0 image as base for my application
In the future, I will replace it with a scratch image which contains the only main app

And most important. Don't forget to set ENV variables:
- CORE_PEER_ADDRESS - IP:PORT\DNS name of Peer 
- CORE_PEER_LOCALMSPID - MSP name
- CORE_PEER_MSPCONFIGPATH - Path to MSP folder
- FABRICORGI_ORDERER_IP - Orderer DNS\IP:PORT

(optional)

- CORE_PEER_TLS_ENABLED - Set "true" if u using TLS in HLF
- CORE_PEER_TLS_ROOTCERT_FILE - Path to CA-root certificate

---

## ROADMAP
- [X] Base methods for API
- [X] Get channel name from request
- [ ] Replace binaries via source code of HLF
- [ ] Add authorization for methods
- [ ] Implement SwaggerUI
- [ ] ...
