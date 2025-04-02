


$ sudo apt-get install uuid-dev cmake

$ git clone https://github.com/antlr/antlr4.git
$ cd antlr4
$ git checkout 4.12.0

$ cd runtime/Cpp
$ mkdir build && cd build
$ cmake ..
$ make -j4
$ sudo make install

$ cd ../.. # Return to the root antlr4 directory

$ wget https://www.antlr.org/download/antlr-4.12.0-complete.jar
$ java -jar antlr-4.12.0-complete.jar -Dlanguage=Cpp -visitor CustomLanguage.g4

$ java -jar antlr-4.12.0-complete.jar -Dlanguage=Go -visitor CustomLanguage.g4


$ java -jar antlr-4.12.0-complete.jar -Dlanguage=Go -package main -visitor -listener CustomLanguage.g4


(newer jar)
$ wget https://www.antlr.org/download/antlr-4.13.1-complete.jar

$ rm -f *.interp *.tokens *.go

$ java -jar antlr-4.13.1-complete.jar \
  -Dlanguage=Go \
  -package main \
  -visitor \
  -listener \
  CustomLanguage.g4


$ go mod init app
$ go mod tidy


$ go run *.go


$ wget https://apt.llvm.org/llvm.sh
$ chmod +x llvm.sh

$ sudo ./llvm.sh 15

$ clang-15 program.ll -o program