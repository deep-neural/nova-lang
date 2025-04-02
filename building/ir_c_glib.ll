; ModuleID = 'glib_interval'
source_filename = "glib_interval"

; External function declarations
declare i32 @printf(i8*, ...) nounwind
declare i8* @g_main_loop_new(i8*, i32) nounwind
declare void @g_main_loop_run(i8*) nounwind
declare i32 @g_timeout_add(i32, i32 (i8*)*, i8*) nounwind  ; Corrected declaration

; String format for printing counter
@.str = private unnamed_addr constant [4 x i8] c"%d\0A\00", align 1
@counter_format = global i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i32 0, i32 0)

; Global variables
@counter = global i32 0
@main_loop = global i8* null

; Callback function for interval (correctly returns i32)
define i32 @interval_callback(i8* %user_data) {
entry:
    %counter_val = load i32, i32* @counter
    %new_val = add i32 %counter_val, 1
    store i32 %new_val, i32* @counter
    %fmt = load i8*, i8** @counter_format
    call i32 (i8*, ...) @printf(i8* %fmt, i32 %new_val)
    ret i32 1  ; Continue the timer
}

; Main function
define i32 @main() {
entry:
    ; Create main loop
    %loop = call i8* @g_main_loop_new(i8* null, i32 0)
    store i8* %loop, i8** @main_loop
    
    ; Add timeout with CORRECT FUNCTION TYPE
    %timeout_id = call i32 @g_timeout_add(
        i32 100,
        i32 (i8*)* @interval_callback,  ; Matches the expected signature
        i8* null
    )
    
    ; Run main loop
    %loop_ptr = load i8*, i8** @main_loop
    call void @g_main_loop_run(i8* %loop_ptr)
    
    ret i32 0
}