/*

	BEFORE TALKING ABOUT POINTER
	LETS TALK ABOUT MEMORY MANAGEMENT

	In golang we dont have to do memory management on our own
	go get this shit done by itself

	do to do memory management we have two ways

	//////
	new() -: allocate memory but no INIT(means initialize it)
			 you will get a memory address
			 zeored stroage

	make()-: allocate memory and also INIT
			 you will get a memory address
			 non-zeored storage


	Garbage collection happens automatically on golang
	whenver varibale goes out of scope or nil it get handled by gc

	https://pkg.go.dev/runtime



*/