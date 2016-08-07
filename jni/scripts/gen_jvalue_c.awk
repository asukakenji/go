BEGIN {
	types[1] = ""
	fields[1] = ""
	line_index = 1
	decl_index = 1
}

function capitalize(s, count) {
	return toupper(substr(s, 1, count)) substr(s, count + 1)
}

function get_type(line) {
	match($0, /^[a-z]+/)
	return substr($0, RSTART, RLENGTH)
}

function get_field(line) {
	match($0, /[a-z];$/)
	return substr($0, RSTART, RLENGTH - 1)
}

END {
	# Beginning Stub
	print "#include \"jvalue.h\""

	for (i = 1; i <= length(types); i++) {
		print ""
		print "jvalue _GoJni" capitalize(types[i], 2) "ToJValue(" types[i] " " fields[i] ")"
		print "{"
		print "\tjvalue v = {." fields[i] " = " fields[i] "};"
		print "\treturn v;"
		print "}"
	}
	for (i = 1; i <= length(types); i++) {
		print ""
		print types[i] " _GoJniJValueTo" capitalize(types[i], 2) "(jvalue v)"
		print "{"
		print "\treturn v." fields[i] ";"
		print "}"
	}
}

/^typedef union jvalue {$/, /^} jvalue;$/ {
	sub(/^[\t ]+/, "", $0)
	if ($0 !~ /^j/) {
		next
	}

	# Type and Field
	types[decl_index] = get_type($0)
	fields[decl_index] = get_field($0)
	++decl_index
}
