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
	struct  sockaddr_in clentaddr;
	memset(&clentaddr, 0, sizeof(clentaddr));
	clentaddr.sin_family = AF_INET;
	inet_pton(AF_INET, "127.0.0.1", &clentaddr.sin_addr.s_addr);
	clentaddr.sin_port = htons(8888);
	int ret = connect(fd, (struct sockaddr *)&clentaddr, sizeof(clentaddr));
	if (ret == -1) {
		cout << "客户端connect连接失败 \n";
		perror("connect");
		exit(-1);
	}
	// 3 通信

	 char recvBuf[1024] = {};
	 char client_data[1024];
	 while(1) {
	 	memset(client_data,0,sizeof(client_data));
	 	printf("请输入消息\n");
	 	// scanf("%s", client_data);
	 	cin >> client_data;
	 	write(fd, client_data, strlen(client_data));
	 	sleep(1);
	 	int ret = read(fd, recvBuf,sizeof(recvBuf));
	 	if (ret == -1) {
	 		perror("read");
	 		exit(-1);
	 	} else if (ret > 0) {
	 		printf(" 服务器信息 %s\n", recvBuf);
	 		continue;
	 	} else if (recvBuf == "exit") {
	 		printf(" 断开连接\n");
	 		break;
	 	}

	 }
	 close(fd);
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