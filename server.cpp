#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netinet/in.h>

using namespace std;
int main() {
	// 1  使用socket
	int sfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sfd == -1) {
		printf(" socket 创建失败 \n");
		exit(-1);
	}
	// 2 绑定 IP 和端口
	struct sockaddr_in saddr;// 专用 socket 地址
	// 协议族 ipv4的AF_INET 也可以
	saddr.sin_family = PF_INET;
	saddr.sin_addr.s_addr = INADDR_ANY;
	saddr.sin_port = htons(8888);
	int ret = bind(sfd, (struct sockaddr *)&saddr, sizeof(saddr));
	if (ret == -1) {
		printf(" 绑定 IP 和 端口号失败 \n");
		exit(-1);
	}
	// 3 监听
	ret = listen(sfd, 128);
	if (ret == -1) {
		printf(" 监听失败 \n");
	}
	// 4 接收客户端
	struct sockaddr_in client_addr;
}