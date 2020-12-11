// https://stackoverflow.com/questions/27914311/get-computer-name-and-logged-user-name
#include <unistd.h>
#include <iostream>

#define HOST_NAME_MAX 1024
#define LOGIN_NAME_MAX 1024

char hostname[HOST_NAME_MAX];
char username[LOGIN_NAME_MAX];
int result;

int main() {
    gethostname(hostname, HOST_NAME_MAX);
    getlogin_r(username, LOGIN_NAME_MAX);
    result = printf("Username: %s \nHostname: %s.\n",username, hostname);
    return 0;
}
