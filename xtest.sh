#!/bin/bash

LIMIT_FOR_EACH_CASE=100
count_for_case=0
total=0
report=()

function exec_test() {
    total=`expr ${total} + 1`
    TESTCASE=${1}
    count_for_case=`expr ${count_for_case} + 1`
    if go test ./tests -run ${TESTCASE} ; then
        echo "--- OK ${count_for_case}: ${TESTCASE}"
        report+=("${count_for_case}	${TESTCASE}")
        count_for_case=0
    else
        if [ ${count_for_case} -gt ${LIMIT_FOR_EACH_CASE} ]; then
            echo ">> FAILED: TIMEOUT: ${count_for_case} times for ${TESTCASE} <<"
            exit 1
        fi
        echo "--- TRY AGAIN: ${count_for_case}: ${TESTCASE}"
        exec_test ${TESTCASE}
    fi
}

function gather_testcases() {
  targets=()
  if [ -n "${1}" ]; then
    list=`ls ./tests/*${1}*_test.go`
  else
    list=`ls ./tests/*_test.go`
  fi
  for filename in ${list} ; do
    targets+=(`echo ${filename} | sed -e s/\.\\\/tests\\\/xtest_// | sed -e s/_test\.go//`)
  done
  echo ${targets[@]}
}

function final_report() {
  IFS=''
  echo "[COMPLETED]"
  for r in ${report[@]} ; do
    echo ${r}
  done
  echo "-------------------------"
  echo "${total}	TOTAL"
}

function __main__() {
  testcases=`gather_testcases ${1}`
  for t in ${testcases[@]} ; do
    exec_test ${t}
  done
  final_report
}

__main__ $@
