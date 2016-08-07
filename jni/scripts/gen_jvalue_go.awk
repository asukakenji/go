BEGIN {
	types[1] = ""
	names[1] = ""
	line_index = 1
	decl_index = 1

	# Java to Go Map
	j_to_g_map["jboolean"] = "bool"
	j_to_g_map["jbyte"] = "int8"
	j_to_g_map["jchar"] = "uint16"
	j_to_g_map["jshort"] = "int16"
	j_to_g_map["jint"] = "int32"
	j_to_g_map["jlong"] = "int64"
	j_to_g_map["jfloat"] = "float32"
	j_to_g_map["jdouble"] = "float64"
	j_to_g_map["jobject"] = "JObject"

	# Go to Java Map
	for (k in j_to_g_map) {
		g_to_k_map[j_to_g_map[k]] = k
	}
}

function get_type(line) {
	match($0, /^[a-z]+/)
	return substr($0, RSTART, RLENGTH)
}

function get_name(line) {
	match($0, /[a-z];$/)
	return substr($0, RSTART, RLENGTH - 1)
}

function capitalize(s, count) {
	return toupper(substr(s, 1, count)) substr(s, count + 1)
}

function wrap(type, name) {
	if (type != "jobject") {
		return "Go" capitalize(j_to_g_map[type], 1) "(" name ")"
	}
	return "NewJObject(" name ")"
}

function unwrap(type, name) {
	if (type != "jobject") {
		return "Java" capitalize(substr(type, 2), 1) "(" name ")"
	}
	return name ".Peer()"
}

END {
	# Beginning Stub
	print "package main"
	print ""
	print "// #include \"jvalue.h\""
	print "import \"C\""

	for (i = 1; i <= length(types); i++) {
		print ""
		print "func " capitalize(types[i], 2) "ToJValue(" names[i] " C." types[i] ") C.jvalue {"
		print "\treturn C._GoJni" capitalize(types[i], 2) "ToJValue(" names[i] ")"
		print "}"
	}
	for (i = 1; i <= length(types); i++) {
		print ""
		print "func JValueTo" capitalize(types[i], 2) "(v C.jvalue) C." types[i] " {"
		print "\treturn C._GoJniJValueTo" capitalize(types[i], 2) "(v)"
		print "}"
	}
	for (i = 1; i <= length(types); i++) {
		print ""
		print "func JValueFrom" capitalize(j_to_g_map[types[i]], 1) "(" names[i] " " j_to_g_map[types[i]] ") JValue {"
		print "\treturn JValue{" capitalize(types[i], 2) "ToJValue(" unwrap(types[i], names[i]) ")}"
		print "}"
	}
	for (i = 1; i <= length(types); i++) {
		print ""
		print "func (value JValue) " capitalize(j_to_g_map[types[i]], 1) "() " j_to_g_map[types[i]] " {"
		print "\treturn " wrap(types[i], "JValueTo" capitalize(types[i], 2) "(value.peer)")
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
