#!/usr/bin/env bash

SEED=$((RANDOM * RANDOM + RANDOM))

TAGS=(
    'avl_bools avl_newn noassert mva'
    'avl_bools avl_newn noassert sva'
    'avl_bools avl_ptrn noassert mva'
    'avl_bools avl_ptrn noassert sva'
    'avl_bools avl_ptrn assert sva'
)

for tags in "${TAGS[@]}"
do
    echo "Tags: $tags"
    SEED=$SEED go test --tags "$tags" -bench=.
    echo && echo && echo
done
