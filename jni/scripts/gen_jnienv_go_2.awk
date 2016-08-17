BEGIN {
	# declarations[1] = ""
	# pretty_declarations[i] = ""
	# method_names[1] = ""
	# return_types[1] = ""
	# parameter_types[1, 1] = ""
	# parameter_names[1, 1] = ""
	# param_indices[1] = 1
	state = 0
	decl_index = 0
}

function join(array, separator, size) {
	buffer = array[1]
	for (j = 2; j <= size; ++j) {
		buffer = buffer separator array[j]
	}
	return buffer
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

function save_decl(declaration) {
	declarations[++decl_index] = declaration
}

function append_decl(declaration) {
	declarations[decl_index] = declarations[decl_index] declaration
}

function get_method_name(declaration) {
	match(declaration, /\(JNICALL\* [^)]+\)/)
	return substr(declaration, RSTART + 10, RLENGTH - 11)
}

function get_return_type(declaration) {
	match(declaration, /^.* \(JNICALL/)
	return substr(declaration, 1, RLENGTH - 9)
}

function get_parameter_list(declaration) {
	gsub(/\(JNICALL\* [^)]+\)/, "", declaration)
	match(declaration, /\(.*\);$/)
	return substr(declaration, RSTART + 1, RLENGTH - 3)
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

function print_c_definition(i) {
	pdecl = pretty_declarations[i]
	sub("\\(JNICALL\\* " method_names[i] "\\)", "_GoJni" method_names[i], pdecl)
	alist = join2d(parameter_names, ", ", i, param_indices[i])
	print "static " substr(pdecl, 1, length(pdecl) - 1)
	print "{"
	printf("\t" "return (*env)->%s(%s);" "\n", method_names[i], alist)
	print "}"
}

function print_go_definition(i) {
	# Parameter List
	pl = get_transformed_parameter_list(i)

	# Return Type
	rt = transform_type(return_types[i])

	# Method
	if (rt == "") {
		printf("func %s(%s) {" "\n", method_names[i], pl)
	} else {
		printf("func %s(%s) %s {" "\n", method_names[i], pl, rt)
	}

	# Argument List
	al = get_transformed_argument_list(i)

	# Invocation
	if (rt == "") {
		printf("\t" "C._GoJni%s(%s)" "\n", method_names[i], al)
	} else {
		printf("\t" "return C._GoJni%s(%s)" "\n", method_names[i], al)
	}

	# End
	print "}"
}

END {
	# 2nd Pass
	for (i = 1; i <= decl_index; ++i) {
		# Reformat the pointer asterisk
		pdecl = declarations[i]
		gsub(/ +\*\*/, "** ", pdecl)
		gsub(/ +\*/, "* ", pdecl)
		gsub(/\*\* +/, "** ", pdecl)
		gsub(/\* +/, "* ", pdecl)
		pretty_declarations[i] = pdecl

		# Method Name
		method_names[i] = get_method_name(pdecl)

		# Return Type
		return_types[i] = get_return_type(pdecl)

		# Parameter Types and Names
		pl = get_parameter_list(pdecl)
		for (j = 1; 1; ++j) {
			pt1 = get_first_parameter_type(pl)
			if (pt1 == "") {
				break
			}
			param_indices[i] = j
			parameter_types[i, j] = pt1
			parameter_names[i, j] = get_first_parameter_name(pl)
			pl = remove_first_parameter(pl)
		}
	}

	# Beginning Stub
	print "package main"
	print ""

	# C Implementation - Begin
	print "/*"
	print "// See src/os/user/lookup_unix.go for usage of \"static\""
	print ""
	print "#include <jni.h>"

	# C Implementation - Comments and Definitions
	for (i = 1; i <= decl_index; ++i) {
		if (method_names[i] "V" == method_names[i+1]) {
			# Comment
			print ""
# 			print "// jni.h:"
# 			print "//     " declarations[i]
		} else if (method_names[i] == method_names[i-1] "V") {
			# Comment
# 			print "//     " declarations[i]
		} else if (method_names[i] == method_names[i-2] "A") {
			# Comment
# 			print "//     " declarations[i]

			# Definition
			print_c_definition(i)
		} else {
			# Comment
			print ""
# 			print "// jni.h:"
# 			print "//     " declarations[i]

			# Definition
			print_c_definition(i)
		}
	}

	# C Implementation - End
	print "*/"
	print "import \"C\""
	print ""
	print "import ("
	print "\t" "\"unsafe\""
	print ")"

	# Go Implementation
	for (i = 1; i <= decl_index; ++i) {
		if (method_names[i] "V" == method_names[i+1]) {
			# Comment
			print ""
			print "// jni.h:"
			print "//     " declarations[i]
		} else if (method_names[i] == method_names[i-1] "V") {
			# Comment
			print "//     " declarations[i]
		} else if (method_names[i] == method_names[i-2] "A") {
			# Comment
			print "//     " declarations[i]

			# Definition
			print_go_definition(i)
		} else {
			# Comment
			print ""
			print "// jni.h:"
			print "//     " declarations[i]

			# Definition
			print_go_definition(i)
		}
	}
}

# 1st Pass
/^struct JNINativeInterface_ {$/, /^};$/ {
	sub(/^[[:space:]]+/, "", $0)
	if (state == 0) {
		if ($0 !~ /JNICALL/) {
			next
		}
		save_decl($0)
		if ($0 !~ /;$/) {
			state = 1
		}
	} else {
		if ($0 ~ /,$/) {
			append_decl($0 " ")
		} else if ($0 ~ /;$/) {
			append_decl($0)
			state = 0
		} else {
			printf("Cannot handle line #%d: %s" "\n", NR, $0)
		}
	}
}
