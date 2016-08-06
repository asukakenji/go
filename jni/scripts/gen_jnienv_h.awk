BEGIN {
	# Constants
	# This does not work:
	#     RE_GET_NAME = /\(JNICALL \*[^)]+\)/
	RE_GET_NAME = "\\(JNICALL \\*[^)]+\\)"

	# Variables
	state = 0
	pending = ""
	declarations[0] = ""
	decl_index = 0

	# Beginning Stub
	print "#ifndef _GO_JNI_JNIENV_H_"
	print "#define _GO_JNI_JNIENV_H_"
	print ""
	print "#include <jni.h>"
	print ""
}

function save(declaration) {
	declarations[decl_index++] = declaration
}

function get_name(declaration) {
	match(declaration, RE_GET_NAME)
	return substr(declaration, RSTART + 10, RLENGTH - 11)
}

END {
	# Notice: "index" is a built-in function.
	# Notice: AWK arrays are associative, so j is needed.
	name = ""
	family = ""
	declaration = ""
	j = length(declarations)
	for (i = 0; i < j; ++i) {
		declaration = declarations[i]
		next_declaration = declarations[i+1]
		name = get_name(declaration)
		next_name = get_name(next_declaration)
		if (next_name == name "V") {
			print "// jni.h:"
			print "//     " declaration
			family = name
		} else if (name == family "V") {
			print "//     " declaration
		} else if (name == family "A") {
			print "//     " declaration
			sub(RE_GET_NAME, "_GoJni" name, declaration)
			gsub(/ *\* */, "* ", declaration)
			print declaration
			print ""
			family = ""
		} else {
			print "// jni.h:"
			print "//     " declaration
			sub(RE_GET_NAME, "_GoJni" name, declaration)
			gsub(/ *\* */, "* ", declaration)
			print declaration
			print ""
		}
	}

	# Ending Stub
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
