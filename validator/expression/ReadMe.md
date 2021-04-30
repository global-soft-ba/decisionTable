# Install 
install antlr4 lib for osx

```
OS X
$ cd /usr/local/lib
$ sudo curl -O https://www.antlr.org/download/antlr-4.9.2-complete.jar
$ export CLASSPATH=".:/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH"
$ alias antlr4='java -jar /usr/local/lib/antlr-4.9.2-complete.jar'
$ alias grun='java org.antlr.v4.gui.TestRig'
```

or 
```
$ brew install antlr
```

#Create Code Generation for Parser

Following command must be executed in directory of the .g4 file.
```
$   antlr4 -Dlanguage=Go -o generated SFeel.g4
```
