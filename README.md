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

    [GET] /user?phone&password - authenticate user, return token

    [GET] /cars - list of available cars for current user

    [POST] /cars - post a new car
    {
        Model   string
        Year    int
        Mileage int
        Image   string [optional]
    }

    *[GET] /car/{id} - show
        full description for chosen car,
        finds all related dates and prices

    >[POST] /car/dates - add new available period. Merged with existing one if they intersect *
    {
        CarId
        DayOfWeek
        Start
        EndTime
    }

    [POST] /car/prices - gets all price units for the car
    {
        CarId
        TimeUnit
        Price
    }


TODO:

    Create table: history of rents

    count total price for period
    [GET] /total
    ?

    add rent
    [POST] /rent
    {

    }

    show list of rents for user
    [GET] rent/{id}

    delete rent
    [DELETE] /rent/{id}

    rent history
    [GET] /history

    delete available time item
    [DELETE] /time/{id} - delete available time item

    change price per time unit
    [PUT] /price/{id} - change price per time unit
    {
        Price float64
    }

    delete price
    [DELETE] / price/{id}

    delete car
    [DELETE] /car/{id} - cascade delete the car, prices and dates for it

