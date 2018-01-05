/*
 * Copyright (c) 2018 Math Codder <mc@endocode.com>
 * 
 * SPDX-License-Identifier: GPL-3.0
 */

#include <stdlib.h>
#include <stdio.h>
#include <json-c/json.h>
#include "calc.h"

void usage(char * prog);
void printResult(const char * op, int a, int b, int res);

int main(int argc, char ** argv) {
    if (argc != 4) {
        usage(argv[0]);
        return 1;
    }

    char sw = argv[1][0];
    int res, a, b;

    a = atoi(argv[2]);
    b = atoi(argv[3]);

    switch (sw) {
    case 'a':
        res = Add(a, b);
        break;
    case 's':
        res = Subtract(a, b);
        break;
    case 'm':
        res = Multiply(a, b);
        break;
    case 'd':
        res = Devide(a, b);
        break;
    default:
        usage(argv[0]);
        return 1;
    }

    printResult(argv[1], a, b, res);
}

void usage(char * prog) {
    printf("%s <op> <int> <int>\n", prog);
    printf("Op can be one of add,sub,mul,div\n\n");
}

void printResult(const char * op, int a, int b, int res) {
    struct json_object *obj = json_object_new_object();

    json_object_object_add(obj, "op", json_object_new_string(op));
    json_object_object_add(obj, "a", json_object_new_int(a));
    json_object_object_add(obj, "b", json_object_new_int(b));
    json_object_object_add(obj, "res", json_object_new_int(res));

    printf("You calld for op %s with numbers %d and %d and result is: %d\n", op, a, b, res);
    printf("Here's a json for you\n");
    printf("%s\n", json_object_to_json_string(obj));
}
