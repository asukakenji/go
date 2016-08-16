BEGIN {
	# comments[1] = ""
	# methods[1] = ""
	# return_types[1] = ""
	# parameter_types[1, 1] = ""
	# parameter_names[1, 1] = ""
	# param_indices[1] = 1
	comm_index = 0
	meth_index = 0
}

function join(array, separator, size) {
	buffer = array[1]
	for (j = 2; j <= size; ++j) {
		buffer = buffer separator array[j]
	}
	return buffer
}

function save_comments(comment) {
	if (comm_index == meth_index) {
		comments[++comm_index] = comment
	} else {
		comments[comm_index] = comments[comm_index] "\n" comment
	}	
}

function save_method(method) {
	++meth_index
	methods[meth_index] = method
	param_indices[meth_index] = 0
}

function save_return_type(return_type) {
	return_types[meth_index] = return_type
}

function save_parameter_type_and_name(parameter_type, parameter_name) {
	++param_indices[meth_index]
	parameter_types[meth_index, param_indices[meth_index]] = parameter_type
	parameter_names[meth_index, param_indices[meth_index]] = parameter_name
}

function get_return_type(declaration) {
	match(declaration, / _GoJni/)
	return substr(declaration, 1, RSTART - 1)
}

function get_method(declaration) {
	match(declaration, /_GoJni[^(]+/)
	return substr(declaration, RSTART + 6, RLENGTH - 6)
}

function get_parameter_list(declaration) {
	match(declaration, /\([^)]+\)/)
	return substr(declaration, RSTART + 1, RLENGTH - 2)
}

function get_first_parameter_type(parameter_list) {
	if (parameter_list == "") {
		return ""
	}
	match(parameter_list, / [^ ,]+(, |$)/)
	return substr(parameter_list, 1, RSTART - 1)
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

function transform_type(type) {
	if (type == "void") {
		return ""
	}
	if (type == "void*") {
		return "unsafe.Pointer"
	}
	if (substr(type, 1, 6) == "const ") {
		type = substr(type, 7)
	}
	asterisks = ""
	while (1) {
		len = length(type)
		if (substr(type, len) == "*") {
			type = substr(type, 1, len - 1)
			asterisks = "*" asterisks
		} else {
			break
		}
	}
	return asterisks "C." type
}

function transform_parameter_name(parameter_name) {
	if (parameter_name == "len") {
		return "_len"
	}
	return parameter_name
}

function get_transformed_parameter_list(mi) {
	for (pi in pa) {
		delete pa[pi]
	}
	for (pi = 1; pi <= param_indices[mi]; ++pi) {
		pt = parameter_types[mi, pi]
		pt1 = parameter_types[mi, pi + 1]
		pn = parameter_names[mi, pi]
		if (pt == pt1) {
			pa[pi] = transform_parameter_name(pn)
		} else {
			pa[pi] = transform_parameter_name(pn) " " transform_type(pt)
		}
	}
	return join(pa, ", ", param_indices[mi])
}

function get_transformed_argument_list(mi) {
	for (ai in aa) {
		delete aa[ai]
	}
	for (ai = 1; ai <= param_indices[mi]; ++ai) {
		pn = parameter_names[mi, ai]
		aa[ai] = transform_parameter_name(pn)
	}
	return join(aa, ", ", param_indices[mi])
}

END {
	# Beginning Stub
	print "package main"
	print ""
	print "// #include \"jnienv.h\""
	print "import \"C\""
	print ""
	print "import ("
	print "\t" "\"unsafe\""
	print ")"

	for (i = 1; i <= meth_index; ++i) {
		# Comment
		print ""
		print comments[i]

		# Parameter List
		pl = get_transformed_parameter_list(i)

		# Return Type
		rt = transform_type(return_types[i])

		# Method
		if (rt == "") {
			print "func " methods[i] "(" pl ") {"
		} else {
			print "func " methods[i] "(" pl ") " rt " {"
		}

		# Argument List
		al = get_transformed_argument_list(i)

		# Invocation
		if (rt == "") {
			print "\t" "C._GoJni" methods[i] "(" al ")"
		} else {
			print "\t" "return C._GoJni" methods[i] "(" al ")"
		}

		# End
		print "}"
	}
}

/^\/\// {
	save_comments($0)
}

/^[^#\/]/ {
	# Method
	m = get_method($0)
	save_method(m)

	# Return Type
	rt = get_return_type($0)
	save_return_type(rt)

	# Parameter List
	pl = get_parameter_list($0)
	while (1) {
		pt1 = get_first_parameter_type(pl)
		if (pt1 == "") {
			break
		}
		pn1 = get_first_parameter_name(pl)
		save_parameter_type_and_name(pt1, pn1)
		pl = remove_first_parameter(pl)
	}
}
