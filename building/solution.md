

1: 
The solution is to use the clang c ast format for our own custom language 
$ clang -Xclang -ast-dump=json -fsyntax-only template.c

2:
Then we use that standard format to output the LLVM IR 

3:
Compile the IR files into on binary


// [modules] important notes
all modules and all files will just become one big LLVM IR .ll file 


// modules
we can also use the clang -Xclang -ast to get header files and use golang program
to automate the process of creating External declarations.

// your repo can now include .h files directly
//    git host                 your repo name
from "gitlab.com/user" import "glib-repo" (  
    version = "2.0",
    branch = "main",
)

; External declarations (using ptr)
declare i32 @printf(ptr, ...)
declare i32 @g_timeout_add(i32, ptr, ptr)
declare ptr @g_main_loop_new(ptr, i32)
declare void @g_main_loop_run(ptr)