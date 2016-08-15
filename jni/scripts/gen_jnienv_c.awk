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

func get_method(declaration) {
	match(declaration, /_GoJni[^(]+/)
	return substr(declaration, RSTART + 6, RLENGTH - 6)
}

func get_parameter_list(declaration) {
	match(declaration, /\([^)]+\)/)
	return substr(declaration, RSTART + 1, RLENGTH - 2)
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

/^[^#\/]/ {
	# Method
	m = get_method($0)

	# Parameter List
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
		pns[i] = pn1
		pl = remove_first_parameter(pl)
		++i
	}

	# Invocation
	print ""
	print substr($0, 1, length($0) - 1)
	print "{"
	print "\t" "return (*env)->" m "(" join(pns, ", ") ");"
	print "}"
}
