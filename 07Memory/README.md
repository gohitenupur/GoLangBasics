<!-- Memory Management -->

<!-- http://pkg.go.dev/runtime -->

1 Memory allocation and deallocation happens automatically

2 use 2 methods 
     a) new() - allocate memory but no init
                you will get a memory address
                zero storage
     b) make() - allocate memory and init 
                you will get a memory address
                non -zeroed storage

3 GC happens automatically
Out of scope and nil

