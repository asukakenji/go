BEGIN {
	# lines[1] = ""
	# types[1] = ""
	# fields[1] = ""
	line_index = 0
	decl_index = 0
}

function capitalize(s, count) {
	return toupper(substr(s, 1, count)) substr(s, count + 1)
}

function save_line(line) {
	lines[++line_index] = line
}

function save_type_and_field(type, field) {
	++decl_index
	types[decl_index] = type
	fields[decl_index] = field
}

function get_type(line) {
	match($0, /[a-z]+/)
	return substr($0, RSTART, RLENGTH)
}

function get_field(line) {
	match($0, /[a-z]+;$/)
	return substr($0, RSTART, RLENGTH - 1)
}

END {
	# Beginning Stub
	print "package main"
	print ""

	# C Implementation - Begin
	print "/*"
	print "// See src/os/user/lookup_unix.go for usage of \"static\""
	print ""
	print "#include <jni.h>"

	# C Implementation - Comment (declaration in <jni.h>)
	print ""
	print "// jni.h:"
	for (i = 1; i <= line_index; ++i) {
		print "// " lines[i]
	}

	# C Implementation - jvalue From Any
	print ""
	print "// jvalue From Any"
	for (i = 1; i <= decl_index; ++i) {
		print ""
		printf("static jvalue _GoJniJValueFrom%s(%s %s)\n", capitalize(types[i], 2), types[i], fields[i])
		print "{"
		printf("\tjvalue v = {.%s = %s};\n", fields[i], fields[i])
		print "\treturn v;"
		print "}"
	}

	# C Implementation - jvalue To Any
	print ""
	print "// jvalue To Any"
	for (i = 1; i <= decl_index; ++i) {
		print ""
		printf("static %s _GoJniJValueTo%s(jvalue v)\n", types[i], capitalize(types[i], 2))
		print "{"
		printf("\treturn v.%s;\n", fields[i])
		print "}"
	}

	# C Implementation - End
	print "*/"
	print "import \"C\""

	# Go Implementation - jvalue From Any
	print ""
	print "// jvalue From Any"
	for (i = 1; i <= decl_index; ++i) {
		print ""
		printf("func JValueFrom%s(%s C.%s) C.jvalue {\n", capitalize(types[i], 2), fields[i], types[i])
		printf("\treturn C._GoJniJValueFrom%s(%s)\n", capitalize(types[i], 2), fields[i])
		print "}"
	}

	# Go Implementation - jvalue To Any
	print ""
	print "// jvalue To Any"
	for (i = 1; i <= decl_index; ++i) {
		print ""
		printf("func JValueTo%s(v C.jvalue) C.%s {\n", capitalize(types[i], 2), types[i])
		printf("\treturn C._GoJniJValueTo%s(v)\n", capitalize(types[i], 2))
		print "}"
	}
}

/^typedef union jvalue {$/, /^} jvalue;$/ {
	# Line
	save_line($0)

	if ($0 !~ /^[[:space:]]*j/) {
		next
	}

	# Type and Field
	t = get_type($0)
	f = get_field($0)
	save_type_and_field(t, f)
}
