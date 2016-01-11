for i in $(ls *.xml)
do
    echo $i
    xmllint --format $i > $i_temp.xml
    mv $i_temp.xml $i
done
