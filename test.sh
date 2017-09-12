
count=0
limit=200

function exec_test() {
    if go test ; then
        echo "[OK]"
    else
        if [ $count -gt $limit ]; then
            echo "[FAILED: TIMEOUT!]"
            exit 1
        fi
        count=`expr $count + 1`
        sleep 0.01s
        echo "[NG: TRY AGAIN] ${count}"
        exec_test
    fi
}

exec_test
