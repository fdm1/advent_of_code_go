# !/bin/bash

go_dirs=('cmd' 'pkg')
for go_dir in ${go_dirs[@]}; do
  for i in $(ls ./${go_dir}); do
    echo
    echo "-----"
    go_test_cmd="go test ./${go_dir}/${i}"
    echo "Running ${go_test_cmd} ..."
    $go_test_cmd
    echo "-----"
    echo
  done
done
