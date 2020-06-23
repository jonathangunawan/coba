package crud

//SQL Query List for table user
const (
	loginSQL = `SELECT 
		user_id, 
		email, 
		address, 
		password 
	FROM public.user 
	WHERE 
		email = $1 AND 
		password = $2;`

	insertUserSQL = `INSERT INTO public.user (
		email,
		address,
		password,
		create_time
	) VALUES (
		:email,
		:address,
		:password,
		CURRENT_TIMESTAMP
	);`

	updateUserSQL = `UPDATE public.user SET
		email = :email,
		address = :address,
		password = :password,
		update_time = CURRENT_TIMESTAMP
	WHERE 
		user_id = :user_id;`

	deleteUserSQL = `DELETE FROM public.user 
	WHERE
		user_id = $1;`

	createTableUserSQL = `CREATE TABLE IF NOT EXISTS public.user(
		user_id serial PRIMARY KEY,
		email VARCHAR (100) UNIQUE NOT NULL,
		address VARCHAR (200) NOT NULL,
		password VARCHAR (200) NOT NULL,
		create_time TIMESTAMP NOT NULL,
		update_time TIMESTAMP
	);`
)

//Get User for login purpose by email and password
func (c UserCRUD) GetUser(args ...interface{}) (UserRow, error) {
	res := UserRow{}

	err := c.conn.Get(&res, loginSQL, args...)
	if err != nil {
		return UserRow{}, err
	}
	return res, nil
}

//Insert New user data
//Only return error and not returning Row Affected and Last ID
func (c UserCRUD) InsertUser(args UserRow) error {
	_, err := c.conn.NamedExec(insertUserSQL, args)
	if err != nil {
		return err
	}

	return nil
}

//Update user data by user_id
//Only return error and not returning Row Affected and Last ID
func (c UserCRUD) UpdateUser(args UserRow) error {
	_, err := c.conn.NamedExec(updateUserSQL, args)
	if err != nil {
		return err
	}

	return nil
}

//Delete user data by user_id
//Only return error and not returning Row Affected and Last ID
func (c UserCRUD) DeleteUser(args int64) error {
	_, err := c.conn.Exec(deleteUserSQL, args)
	if err != nil {
		return err
	}
	return nil
}

//Create user table if not exist
func (c UserCRUD) CreateTable() error {
	_, err := c.conn.Exec(createTableUserSQL)
	if err != nil {
		return err
	}

	return nil
}
