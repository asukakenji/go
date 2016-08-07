BEGIN {
	types[1] = ""
	names[1] = ""
	line_index = 1
	decl_index = 1
}

function get_type(line) {
	match($0, /^[a-z]+/)
	return substr($0, RSTART, RLENGTH)
}

function get_name(line) {
	match($0, /[a-z];$/)
	return substr($0, RSTART, RLENGTH - 1)
}

function capitalize_type(type) {
	return toupper(substr(types[i], 1, 2)) substr(types[i], 3)
}

END {
	# Beginning Stub
	print "#include \"jvalue.h\""

	for (i = 1; i <= length(types); i++) {
		print ""
		print types[i] " _GoJniJValueTo" capitalize_type(types[i]) "(jvalue v)"
		print "{"
		print "\treturn v." names[i] ";"
		print "}"
	}
	for (i = 1; i <= length(types); i++) {
		print ""
		print "jvalue _GoJni" capitalize_type(types[i]) "ToJValue(" types[i] " " names[i] ")"
		print "{"
		print "\tjvalue v = {." names[i] " = " names[i] "};"
		print "\treturn v;"
		print "}"
	}
}

/^typedef union jvalue {$/, /^} jvalue;$/ {
	sub(/^[\t ]+/, "", $0)
	if ($0 !~ /^j/) {
		next
	}

	# Type and Name
	types[decl_index] = get_type($0)
	names[decl_index] = get_name($0)
	++decl_index
}
