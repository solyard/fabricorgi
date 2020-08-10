# FABRICORGI

![Image](fabricorgi.png)

What this tool can do?

- Firstly this tool created for comfortable managing organizations list in your Hyperledger network
- Secondly, this is one of the ways to learn GoLang (for me)
- Third is just a try to make some useful open-source tool for all

---

# HOW TO BUILD

Set the ENV variable CGO_ENABLED to 0 on your system via this command:

```
export CGO_ENABLED=0
```

Then run command in root dir of project:
```
go build fabricorgi.go
```

Then run docker build command to build image with fabricorgi:
```
docker build . -f Dockerfile -t fabricorgi:latest
```
---

# HOW TO USE

Use the output json file produced by `configtxgen` tool and send it into API endpoint depend of your task.
For adding organisation use:
```
example.com:8081/api/v1/addorg
```
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
