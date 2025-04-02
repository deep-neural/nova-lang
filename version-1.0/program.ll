@.str.0 = constant [3 x i8] c"ls\00"

declare i32 @printf(i8* %format, ...)

declare i32 @system(i8* %command)

define i32 @list_all_files() {
entry:
	%0 = getelementptr [3 x i8], [3 x i8]* @.str.0, i32 0, i32 0
	%1 = call i32 @system(i8* %0)
	ret i32 0
}

define i32 @main() {
entry:
	%0 = call i32 @list_all_files()
	ret i32 0
}
