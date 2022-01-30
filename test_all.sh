echo "TEST ALL"
one_failed=0
for dir in src/go/*/
do     
    if [ "$(gofmt -s -l $dir | wc -l)" -gt 0 ]; then 
        echo "formatting failed for" $dir
        gofmt -s -l $dir
        one_failed=1
    fi

    ./src/testing ${dir}

    if [ $? -eq 0 ]; then
        echo ""
    else
        one_failed=1
    fi
done

if [ $one_failed -eq 1 ]; then
    echo "FAIL"
    exit 1
else
    echo "PASS"
    exit 0
fi