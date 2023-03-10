#include <stdio.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <iostream>

using namespace std;
int main() {
	// 1 创建 socket
	int fd = socket(AF_INET, SOCK_STREAM, 0);
	if (fd == -1) {
		cout << "创建socket失败 \n";
		perror("socket");
		exit(-1);
	}


	return 0;
}
// 静态变量作用
// #include <cstring>
// char * Strtok(char *p, char *sep) 
// {
// 	static char *start;
// 	if(p) start = p;

// 	for(; *start && strchr(sep, *start); ++start);

// 	if (*start == 0) return NULL;

// 	char *q = start;
// 	for(; *start && !strchr(seq, *start); ++start);

// 	if (*start) {
// 		*start = 0;
// 		++start;
// 	}
// 	return q;
// }