package cpprof


/*
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#define MAX_STRING 128

void use_system(void) {
	char command[MAX_STRING];

	sprintf(command, "/proc/%d/status", getpid());
	system(command);
}
*/
import "C"
import (
  "net/http"
)

func init() {
  http.HandleFunc("/debug/pprof/cheap", HeapProfile)
}

func HeapProfile(w http.ResponseWriter, r *http.Request) {
  C.use_system()
}
