
count=0
limit=600

function exec_test() {
    if go test -run $1; then
        echo "[OK: ${1}]"
    else
        if [ $count -gt $limit ]; then
            echo "[FAILED: TIMEOUT!]"
            exit 1
        fi
        count=`expr $count + 1`
        echo "[NG: TRY AGAIN: ${1}] ${count}"
        echo "-"
        exec_test $1
    fi
}

targets=()
if [ -n "${1}" ]; then
  targets+=($1)
else
  for filename in `ls xtest_*.go`; do
    targets+=(`echo xtest_basename_fields_test.go | sed -e s/xtest_// | sed -e s/\.go//`)
  done
fi

for t in $targets; do
  echo $t
  exec_test $t
done
