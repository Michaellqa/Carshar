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

    [POST] /car/dates - add new available period
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

    Расчет цены бронирования
    [GET] /total
    {
        StartTime
        EndTime
    }
    -> $ total float

    Добавить бронирование -
        пользователь видел расчет цены за бронирование,
        делаем повторный расчет на сервере, на случай если цены изменились
        если цена совпадает с TargetPrice, которую видел пользователь - бронируем
        иначе отправляем 1xx ошибку, показываем что цена изменилась
    [POST] /rent
    {
        RenterId
        CarId
        StartTime
        EndTime
        TargetPrice
    }
    -> 1xx

    Удаление


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



    Create table: history of rents