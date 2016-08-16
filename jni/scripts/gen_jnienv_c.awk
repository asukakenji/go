BEGIN {
	# lines[1] = ""
	# methods[1] = ""
	# parameter_names[1, 1] = ""
	# param_indices[1] = 1
	line_index = 0
	meth_index = 0
}

# array2d:   The 2D array
# separator: The separator
# index1:    The index of the 1st dimension
# size2:     The size of the 2nd dimension (the last index of array2d[index1])
function join2d(array2d, separator, index1, size2) {
	buffer = array2d[index1, 1]
	for (j = 2; j <= size2; ++j) {
		buffer = buffer separator array2d[index1, j]
	}
	return buffer
}

function save_line(line) {
	lines[++line_index] = line
}

function save_method(method) {
	++meth_index
	methods[meth_index] = method
	param_indices[meth_index] = 0
}

function save_parameter_name(parameter_name) {
	parameter_names[meth_index, ++param_indices[meth_index]] = parameter_name
}

function get_method(declaration) {
	match(declaration, /_GoJni[^(]+/)
	return substr(declaration, RSTART + 6, RLENGTH - 6)
}

function get_parameter_list(declaration) {
	match(declaration, /\([^)]+\)/)
	return substr(declaration, RSTART + 1, RLENGTH - 2)
}

function get_first_parameter_name(parameter_list) {
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

function remove_first_parameter(parameter_list) {
	if (parameter_list == "") {
		return ""
	}
	match(parameter_list, /^[^,]+(, |$)/)
	return substr(parameter_list, RLENGTH + 1)
}

END {
	# Beginning Stub
	print "#include \"jnienv.h\""

	for (i = 1; i <= line_index; ++i) {
		# Argument List
		al = join2d(parameter_names, ", ", i, param_indices[i])

		# Invocation
		print ""
		print substr(lines[i], 1, length(lines[i]) - 1)
		print "{"
		print "\t" "return (*env)->" methods[i] "(" al ");"
		print "}"
	}
}

/^[^#\/]/ {
	# Line
	save_line($0)

	# Method
	m = get_method($0)
	save_method(m)

	# Parameter List
	pl = get_parameter_list($0)
	while (1) {
		pn1 = get_first_parameter_name(pl)
		if (pn1 == "") {
			break
		}
		save_parameter_name(pn1)
		pl = remove_first_parameter(pl)
	}
}
