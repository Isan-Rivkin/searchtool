
# Plugin1: kube_query

Use kubectl but on all of the available k8s clusters available in the kubeconfig file. 
Currently will query only AWS EKS clusters.

## Install 

### from source 

```bash
git clone
# default location on for Macos /usr/local/bin
make release
# or use make release BIN_DIR='/path/to/kq/binary' for custom kq binary location
```

### Target OS: 

```bash

make build-linux

make build-osx

make build-windows
```

### Usage

It's a `kubectl` wrapper that querys all contexts (e.g all of your clusters).
Just use whatever kubectl command with that.  

```bash
kq <kubectl query>
```



### example

```bash
kq get pods --namespace kube-system
kq get ns
kq describe pod --namespace default
```

# Plugin2: Route53 Query 

Command line utility for Route53. 

Get info about your Route53 records from the Command line - quickly!

# Use Cases

### Where does <some dns record> points to? 

<b> Input </b>

```bash
r53 -r '*.foo.goo.website.com'
``` 

<b> Output </b>

```bash
+---+--------+-----------------------+------------------------+
| # | ZoneId | Record                | Target                 |
+---+--------+-----------------------+------------------------+
| 1 | ABC    | *.foo.goo.website.com | lb1.elb.amazonaws.com  |
+---+--------+-----------------------|------------------------+

```


### How it works

Example pseudocode: 

```python
# i.e https://example.com/p/a?ok=11&not=23
# into example.com 
record = '*.<a>.<b><c>.<d>'
record = strip_non_domain(record)
record = split(record)
 
hasWildCard = record[0] == '*'
len == record.length
# a 
if len == 1 and not hasWildCard: 
    lookup(1)
# *.a
if len == 2 and hasWildCard:  
    lookup(1)
# a.b
if len == 2 and not hasWildCard:  
    lookup(2)
# *.a.b
if len == 3 and hasWildCard: 
    lookup(2)
# a.b.c
if len == 3 and hasWildCard: 
    if not lookup(2)
        lookup(3)
# *.a.b.c.d.e
if len > 3:
    for chunksNum in (2, len):
        if (chunksNum == len and hasWildCard):
            return None
        if res = lookup(chunksNum)
            return res 

def lookup(dnsChunk):
    zoneIds = getZoneIds(dnsChunk)
    for zoneId in zoneIds:
        if zoneId name == dnsChunk: 
            aliasTargets = getAliasTargets(record, zoneId)
            return aliasTargetes
```