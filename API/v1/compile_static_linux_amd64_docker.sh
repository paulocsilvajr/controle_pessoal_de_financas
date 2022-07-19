BASE=$(dirname $0)
BASE=$(readlink -f $BASE)
PASTA=$BASE/bin/API_CPF
ARQUIVO=$PASTA/API_CPF

if [ `ls $BASE/keys/ | wc -l` -ne 0 ]; then
    echo "Keys presentes em pasta 'keys/'"
else
    echo "Pasta keys vazia, deve ser gerados as keys via script 'generate_keys.sh'"
    exit 1
fi

# export GOOS=linux
# export GOARCH=amd64
# export CGO_ENABLED=0

if [ -d $PASTA ]; then
    rm -r $PASTA && echo -e "Limpado $PASTA\n"
fi

mkdir -p $PASTA 2> /dev/null

echo -e "Compilando para $GOOS:$GOARCH\n"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-extldflags "-static"' -o $ARQUIVO $BASE

echo -e "\nCopiando 'keys' para pasta de API compilada" &&
    cp -vr $BASE/keys $PASTA &&
    chmod -R 777 $PASTA

echo
file $ARQUIVO && ldd $ARQUIVO

echo -e "\nArquivos em $PASTA:"
tree -D $PASTA
