
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
    targets+=(`echo ${filename} | sed -e s/xtest_// | sed -e s/_test\.go//`)
  done
fi

for t in ${targets[@]} ; do
  exec_test $t
done
