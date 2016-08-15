BEGIN {
	# Beginning Stub
	print "#include \"jnienv.h\""
}

func join(array, separator) {
	buffer = array[1]
	for (i = 2; i <= length(array); i++) {
		buffer = buffer separator array[i]
	}
	return buffer
}

func get_return_type(declaration) {
	match(declaration, /^[^ ]+/)
	return substr(declaration, 1, RLENGTH)
}

func get_method(declaration) {
	match(declaration, /_GoJni[^(]+/)
	return substr(declaration, RSTART + 6, RLENGTH - 6)
}

func get_parameter_list(declaration) {
	match(declaration, /\([^)]+\)/)
	return substr(declaration, RSTART + 1, RLENGTH - 2)
}

func get_first_parameter_type(parameter_list) {
	if (parameter_list == "") {
		return ""
	}
	match(parameter_list, / [^ ,]+(, |$)/)
	return substr(parameter_list, 1, RSTART - 1)
}

func get_first_parameter_name(parameter_list) {
	if (parameter_list == "") {
		return ""
	}
	match(parameter_list, / [^ ,]+(, |$)/)
	first_parameter_name = substr(parameter_list, RSTART + 1, RLENGTH - 1)
	len = length(first_parameter_name)
	if (substr(first_parameter_name, len - 1) == ", ") {
		return substr(first_parameter_name, 1, len - 2)
	}
	return first_parameter_name
}

func remove_first_parameter(parameter_list) {
	if (parameter_list == "") {
		return ""
	}
	match(parameter_list, /^[^,]+(, |$)/)
	return substr(parameter_list, RLENGTH + 1)
}

func transform_type(type) {
	if (type == "void") {
		return ""
	}
	if (substr(type, 1, 6) == "const ") {
		type = substr(type, 7)
	}
	len = length(type)
	if (substr(type, len) == "*") {
		type = "*C." substr(type, 1, len - 1)
	} else {
		type = "C." type
	}
	return type
}

func transform_parameter_name(parameter_name) {
	if (parameter_name == "len") {
		return "_len"
	}
	return parameter_name
}

/^\/\/ jni.h:$/ {
	print ""
	print
}

/^\/\/  / {
	print
}

/^[^#\/]/ {
	# Method
	m = get_method($0)
	printf "func " m "("

	# Parameter List
	pl = get_parameter_list($0)
	for (i in ps) {
		delete ps[i]
	}
	i = 1
	while (1) {
		pt1 = get_first_parameter_type(pl)
		if (pt1 == "") {
			break
		}
		pn1 = get_first_parameter_name(pl)
		ps[i] = transform_parameter_name(pn1) " " transform_type(pt1)
		pl = remove_first_parameter(pl)
		++i
	}
	printf join(ps, ", ")

	# Return Type
	rt = get_return_type($0)
	if (rt == "void") {
		print ") {"
	} else {
		print ") " transform_type(rt) " {"
	}

	# Invocation
	pl = get_parameter_list($0)
	for (i in pns) {
		delete pns[i]
	}
	i = 1
	while (1) {
		pn1 = get_first_parameter_name(pl)
		if (pn1 == "") {
			break
		}
		pns[i] = transform_parameter_name(pn1)
		pl = remove_first_parameter(pl)
		++i
	}
	if (rt == "void") {
		print "\t" "C._GoJni" m "(" join(pns, ", ") ")"
	} else {
		print "\t" "return C._GoJni" m "(" join(pns, ", ") ")"
	}

	# End
	print "}"
}
