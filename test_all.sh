echo "TEST ALL"
one_failed=0
for dir in src/go/*/
do     
    echo ${dir}

    src/testing ${dir}

    if [ $? -eq 0 ]; then
        echo OK
    else
        echo FAIL
        one_failed=1
    fi
done
echo "DONE"

if [ $one_failed -eq 1 ]; then
    echo "failed"
    exit 1
else
    echo "not failed"
    exit 0
fi