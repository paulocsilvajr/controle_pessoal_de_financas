install(){
    echo "$1:"
    go get -u -v $1
}

echo -e "Instalando dependências\n"

install github.com/gorilla/mux
install github.com/lib/pq
install github.com/auth0/go-jwt-middleware
install github.com/jedib0t/go-pretty/table
install github.com/form3tech-oss/jwt-go

# GORM(https://gorm.io/docs/index.html)
install gorm.io/gorm
install gorm.io/driver/postgres
