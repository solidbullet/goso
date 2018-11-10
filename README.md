生成so文件的命令

g++ test_a.cpp test_b.cpp -fPIC -shared -o libtest.so

生成.a文件的命令
gcc -c test_a.cpp
gcc -c test_b.cpp
ar -r libtest.a test_a.o test_b.o

采用动态库编译命令:
g++ test.cpp -o test -L. -ltest

执行：export LD_LIBRARY_PATH=./

最后： ./test
