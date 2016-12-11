#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

#define MAX_STRING 128

void useSystem(void) {
	char command[MAX_STRING];

	sprintf(command, "/proc/%d/status", getpid());
	system(command);
}
