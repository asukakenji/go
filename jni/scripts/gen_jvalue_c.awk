BEGIN {
	# types[1] = ""
	# fields[1] = ""
	decl_index = 0
}

function capitalize(s, count) {
	return toupper(substr(s, 1, count)) substr(s, count + 1)
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
	print "#include \"jvalue.h\""

	for (i = 1; i <= decl_index; ++i) {
		print ""
		print "jvalue _GoJni" capitalize(types[i], 2) "ToJValue(" types[i] " " fields[i] ")"
		print "{"
		print "\tjvalue v = {." fields[i] " = " fields[i] "};"
		print "\treturn v;"
		print "}"
	}
	for (i = 1; i <= decl_index; ++i) {
		print ""
		print types[i] " _GoJniJValueTo" capitalize(types[i], 2) "(jvalue v)"
		print "{"
		print "\treturn v." fields[i] ";"
		print "}"
	}
}

/^typedef union jvalue {$/, /^} jvalue;$/ {
	if ($0 !~ /^[[:space:]]*j/) {
		next
	}

	# Type and Field
	t = get_type($0)
	f = get_field($0)
	save_type_and_field(t, f)
}
