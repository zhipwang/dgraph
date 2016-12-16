### To run the code :
1. have antlr-go runtime :
```
go get github.com/antlr/antlr4/runtime/Go/antlr
cd $GOPATH/src/github.com/antlr/antlr4 # enter the antlr4 source directory
git checkout tags/4.6.0 # the go runtime was added in release 4.6.0
```

full information at : https://github.com/antlr/antlr4/blob/master/doc/go-target.md

2. Now run the tests :
```
# in your Dgraph repo
cd antlr4go
go test -bench=. -test.run=OnlyBench
cd ../gql
go test -bench=. -test.run=OnlyBench
# have a look at the queries
```

If you want to setup antlr for youself then :
### download antlr4.6
cd /usr/local/lib
sudo curl -O http://www.antlr.org/download/antlr-4.6-complete.jar

### create the following environment variables :

```
export CLASSPATH=".:/usr/local/lib/antlr-4.6-complete.jar:$CLASSPATH"
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.6-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java org.antlr.v4.gui.TestRig'
```

Now in dgraph/antlr4go directory :

```
antlr4 -Dlanguage=Go GraphQL.g4
```
