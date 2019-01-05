# Carshar

Now we have authentication logic.

Request specification

    [POST] /users - create new user
    {
    	Name      string
    	Phone     string
    	Password  string
    	BirthDate time.Time
    }

    [GET] /user - find one
    {
        phone     string
        password  string
    }
    ...


TODO:

