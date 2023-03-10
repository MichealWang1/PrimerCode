#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <iostream>
using namespace std;
int main() {
	// 1  使用socket
	int sfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sfd == -1) {
		printf(" socket 创建失败 \n");
		perror("socket");
		exit(-1);
	} else {
		cout << "创建socket成功" << sfd << endl;
	}
	// 2 绑定 IP 和端口
	struct sockaddr_in saddr;// 专用 socket 地址
	// 协议族 ipv4的AF_INET 也可以
	//inet_pton(AF_INET, "127.0.0.1", saddr.sin_addr.s_addr);
	saddr.sin_family = PF_INET;
	saddr.sin_addr.s_addr = INADDR_ANY;
	saddr.sin_port = htons(8888);
	int ret = bind(sfd, (struct sockaddr *)&saddr, sizeof(saddr));
	if (ret == -1) {
		printf(" 绑定 IP 和 端口号失败 \n");
		perror("bind");
		exit(-1);
	} else {
		cout << " 绑定 IP 和 端口号 成功" << ret << endl;
	}
	// 3 监听
	ret = listen(sfd, 128);
	if (ret == -1) {
		printf(" 监听失败 \n");
		perror("listen");
		exit(-1);
	} else {
		cout << " 监听成功 " << ret << endl;
	}
	// 4 接收客户端
	struct sockaddr_in client_addr;
	socklen_t addrlen = sizeof(client_addr);
	int newfd = accept(sfd, (struct sockaddr *) &client_addr, &addrlen);
	if (newfd == -1) {
		printf(" 接收客户端失败 \n");
		perror("accept");
		exit(-1);
	} else {
		cout << " 接收客户端成功 " << newfd << endl;
	}
	// 输出连接消息
	char clientIP[16] = "";
	inet_ntop(AF_INET, &client_addr.sin_addr.s_addr, clientIP,sizeof(clientIP));
	unsigned short clientPort = ntohs(client_addr.sin_port);
	printf("客户端IP %s 端口 %d  连接 \n ", clientIP, clientPort);

	//5 接收客消息和发送消息
	char recvBuf[1024] = {0};
	while(1) {
		memset(recvBuf, 0, sizeof(recvBuf));
		int num = read(newfd, recvBuf, sizeof(recvBuf));
		if (num == -1) {
			printf(" 读取消息失败 ");
			perror("read");
			exit(-1);
		} else if (num > 0) {
			printf(" 客户端发来消息 %s \n", recvBuf);
		} else if(recvBuf == "exit") {
			printf(" 客户端发来消息结束消息 %s \n", recvBuf);
			break;
		} else if(num == 0) {
			printf(" 客户端发来消息结束消息 num == 0 \n");
			break;
		} else {
			printf(" 客户端发来消息结束消息 未之消息 num == %d \n", num);
		}
	}
	// 关闭 socket
	close(sfd);
	close(newfd);
	return 0;
}