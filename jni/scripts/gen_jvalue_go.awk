BEGIN {
	# types[1] = ""
	# fields[1] = ""
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
}

function capitalize(s, count) {
	return toupper(substr(s, 1, count)) substr(s, count + 1)
}

function save_type_and_field(type, field) {
	types[decl_index] = type
	fields[decl_index] = field
	++decl_index
}

function get_type(line) {
	match($0, /[a-z]+/)
	return substr($0, RSTART, RLENGTH)
}

function get_field(line) {
	match($0, /[a-z]+;$/)
	return substr($0, RSTART, RLENGTH - 1)
}

function wrap(type, expr) {
	if (type != "jobject") {
		return "Go" capitalize(j_to_g_map[type], 1) "(" expr ")"
	}
	return "NewJObject(" expr ")"
}

function unwrap(type, expr) {
	if (type != "jobject") {
		return "Java" capitalize(substr(type, 2), 1) "(" expr ")"
	}
	return expr ".Peer()"
}

END {
	# Beginning Stub
	print "package main"
	print ""
	print "// #include \"jvalue.h\""
	print "import \"C\""

	for (i = 1; i < decl_index; i++) {
		print ""
		print "func " capitalize(types[i], 2) "ToJValue(" fields[i] " C." types[i] ") C.jvalue {"
		print "\treturn C._GoJni" capitalize(types[i], 2) "ToJValue(" fields[i] ")"
		print "}"
	}
	for (i = 1; i < decl_index; i++) {
		print ""
		print "func JValueTo" capitalize(types[i], 2) "(v C.jvalue) C." types[i] " {"
		print "\treturn C._GoJniJValueTo" capitalize(types[i], 2) "(v)"
		print "}"
	}
	for (i = 1; i < decl_index; i++) {
		print ""
		print "func JValueFrom" capitalize(j_to_g_map[types[i]], 1) "(" fields[i] " " j_to_g_map[types[i]] ") JValue {"
		print "\treturn JValue{" capitalize(types[i], 2) "ToJValue(" unwrap(types[i], fields[i]) ")}"
		print "}"
	}
	for (i = 1; i < decl_index; i++) {
		print ""
		print "func (value JValue) " capitalize(j_to_g_map[types[i]], 1) "() " j_to_g_map[types[i]] " {"
		print "\treturn " wrap(types[i], "JValueTo" capitalize(types[i], 2) "(value.peer)")
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
