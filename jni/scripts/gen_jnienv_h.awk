BEGIN {
	# Constants
	# This does not work:
	#     RE_GET_METHOD = /\(JNICALL \*[^)]+\)/
	RE_GET_METHOD = "\\(JNICALL \\*[^)]+\\)"

	# Variables
	state = 0
	pending = ""
	declarations[1] = ""
	decl_index = 1

	# Beginning Stub
	print "#ifndef _GO_JNI_JNIENV_H_"
	print "#define _GO_JNI_JNIENV_H_"
	print ""
	print "#include <jni.h>"
}

function save(declaration) {
	declarations[decl_index++] = declaration
}

function get_method(declaration) {
	match(declaration, RE_GET_METHOD)
	return substr(declaration, RSTART + 10, RLENGTH - 11)
}

END {
	# Notice: "index" is a built-in function.
	method = ""
	family = ""
	declaration = ""
	# j MUST be declared here, since when declarations[i+1] is accessed,
	# an entry will be added to the array
	j = length(declarations)
	for (i = 1; i <= j; i++) {
		declaration = declarations[i]
		next_declaration = declarations[i+1]
		method = get_method(declaration)
		next_method = get_method(next_declaration)
		if (next_method == method "V") {
			print ""
			print "// jni.h:"
			print "//     " declaration
			family = method
		} else if (method == family "V") {
			print "//     " declaration
		} else if (method == family "A") {
			print "//     " declaration
			sub(RE_GET_METHOD, "_GoJni" method, declaration)
			# Note: There may be more than one level of redirection
			# (pointer to pointer to something)
			gsub(/ +\*\*/, "** ", declaration)
			gsub(/ +\*/, "* ", declaration)
			gsub(/\*\* +/, "** ", declaration)
			gsub(/\* +/, "* ", declaration)
			print declaration
			family = ""
		} else {
			print ""
			print "// jni.h:"
			print "//     " declaration
			sub(RE_GET_METHOD, "_GoJni" method, declaration)
			# Note: There may be more than one level of redirection
			# (pointer to pointer to something)
			gsub(/ +\*\*/, "** ", declaration)
			gsub(/ +\*/, "* ", declaration)
			gsub(/\*\* +/, "** ", declaration)
			gsub(/\* +/, "* ", declaration)
			print declaration
		}
	}

	# Ending Stub
	print ""
	print "#endif // #ifndef _GO_JNI_JNIENV_H_"
}

/^struct JNINativeInterface_ {$/, /^};$/ {
	sub(/^[\t ]+/, "", $0)
	if (state == 0) {
		if ($0 !~ /JNICALL/) {
			next
		}
		if ($0 ~ /;$/) {
			save($0)
		} else {
			pending = $0
			state = 1
		}
	} else {
		pending = pending $0
		if ($0 ~ /,$/) {
			pending = pending " "
		} else if ($0 ~ /;$/) {
			save(pending)
			state = 0
		}
	}
}
