echo "TEST ALL"
one_failed=0
for dir in src/go/*/
do     
    ./testing ${dir}

    if [ $? -eq 0 ]; then
        echo ""
    else
        echo FAIL
        one_failed=1
    fi
done

if [ $one_failed -eq 1 ]; then
    echo "FAILED"
    exit 1
else
    echo "NOT FAILED"
    exit 0
fi