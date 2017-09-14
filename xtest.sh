
total=0
count=1

limit_for_each=100
function exec_test() {

    total=`expr $total + 1`

    target=$1

    if go test -run $1; then
        echo "[OK: ${count}] $target"
    else
        if [ $count -gt ${limit_for_each} ]; then
            echo "[FAILED: TIMEOUT: ${count} : total ${total}] $target"
            exit 1
        fi
        echo "[NG: TRY AGAIN: ${count}] ${target}"
        echo "-"
        count=`expr $count + 1`
        exec_test $target
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

echo "[COMPLETED: total ${total} times]"
