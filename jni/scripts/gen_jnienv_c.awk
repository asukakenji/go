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
	match(parameter_list, /^[^,]+/)
	parameter = substr(parameter_list, 0, RLENGTH)
	split(parameter, array, / /)
	return array[length(array)]
}

func remove_first_parameter(parameter_list) {
	if (parameter_list == "") {
		return ""
	}
	match(parameter_list, /^[^,]+(, )?/)
	return substr(parameter_list, RLENGTH + 1)
}

# AWK cannot return arrays
func get_parameter_names(parameter_list, pnames) {
	for (i in pnames) {
		delete pnames[i]
	}
	i = 1
	while (1) {
		pname = get_first_parameter_name(parameter_list)
		if (pname == "") {
			break
		}
		pnames[i] = pname
		parameter_list = remove_first_parameter(parameter_list)
		++i
	}
}

/^[^#\/]/ {
	method = get_method($0)
	parameter_list = get_parameter_list($0)
	get_parameter_names(parameter_list, pnames)

	# Get rid of the final semi-colon
	print ""
	print substr($0, 1, length($0) - 1)
	print "{"
	print "\treturn (*env)->" method "(" join(pnames, ", ") ");"
	print "}"
}
