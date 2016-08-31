package db
import(
"log"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"../structures"
)

var connection *gorm.DB

const engine_sql string="mysql"
const username string ="root"
const password string="domg"
const database string="proyectos"

func InitializeDatabase(){
connection = connectORM(createString())
log.Println("conexion exitosa")
}

func CloseConnection(){
	connection.Close()
	log.Println("Conexion cerrada")
}

func connectORM(cadena string) *gorm.DB{
	connection, err := gorm.Open(engine_sql,cadena)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return connection
}

func createString() string{
	return username+":"+password+"@/"+database
}

 func Getuser(id string) structures.User{
 	user:=structures.User{}
	connection.Where("usuario=?",id).First(&user)
	return user
 }
	