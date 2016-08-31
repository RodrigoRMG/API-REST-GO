package structures

type User struct{
	Usuario string `json:"usuario"`
	Password string `json:"password"`
	Nombres string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Email string `json:"Email"`
	tipoUsuario string `json:"tipoUsuario"`
	activo string `json:"activo"`
	created_at string `json:"created_at"`
	updated_at string `json:"updated_at"`
}