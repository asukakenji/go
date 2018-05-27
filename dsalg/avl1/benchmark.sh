#!/usr/bin/env bash

# The newn implementation is a little bit slower than the ptrn implementation.
# However, it has a great architectual advantage.
#
# The ptrn implementation requires a pointer-to-pointer-to-node-struct parameter.
# If its type is changed to pointer-to-node-interface, there would be a problem.
# Taking the address of a pointer-to-node-struct does not turn it into a
# pointer-to-node-interface even though pointer-to-node-struct implements
# node-interface.
#
# On the other hand, the newn implementation returns a pointer-to-node-struct.
# If its type is changed to a node-interface, it still works fine, since
# pointer-to-node-struct implements node-interface. The returned value could be
# assigned to a node-interface variable without a type assetion, or to a
# pointer-to-node-struct with a type assertion.

SEED=$((RANDOM * RANDOM + RANDOM))

TAGS=(
    'avl_bools avl_newn avl_noassert avl_mva'
    'avl_bools avl_newn avl_noassert avl_sva'
#    'avl_bools avl_newn avl_assert avl_sva'
    'avl_bools avl_ptrn avl_noassert avl_mva'
    'avl_bools avl_ptrn avl_noassert avl_sva'
    'avl_bools avl_ptrn avl_assert avl_sva'

    'avl_int avl_newn avl_noassert avl_mva'
    'avl_int avl_newn avl_noassert avl_sva'
#    'avl_int avl_newn avl_assert avl_sva'
    'avl_int avl_ptrn avl_noassert avl_mva'
    'avl_int avl_ptrn avl_noassert avl_sva'
    'avl_int avl_ptrn avl_assert avl_sva'
)

for tags in "${TAGS[@]}"
do
    echo "Tags: $tags"
    SEED=$SEED go test --tags "$tags" -bench=.
    echo && echo && echo
done
