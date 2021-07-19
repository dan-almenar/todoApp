# todoApp
A full scale ToDo App that includes backend developed in Golang, web frontend (html/css/JavaScript) and a mobile app (developed in Python with the Kivy framework). This is a work in progress...

The Initial Commit is called Backbone because it sets up the **basic REST API** part of the Project, which for the moment serves all requests on a single path ("/"). The requests methods handlers are defined in the handlers package.
- **GET requests** will have a jsonified list of TodoTask (if any) or an empty list. **TodoTask** are defined in the **data** package.
- **PUT requests** will take a jsonified TodoTask name and urgent fields and add it to *data.csv* file in the home directory. If *data.csv* file doesn't exist, the backend will create one. TodoTask id (int) is automatically generated, so it doesn't need to be passed in the request.
- **DELETE requests** (requires a TodoTask id) will delete the passed TodoTask from the data.csv file if exists.

### This project is a work in progress and will have major changes in it. However, the Backbone works as intended, thus you may clone this repo and play with the backend features as is. It is set to serve from localhost:8080.

The idea in the long run is to add a users handler so every authenticated user has his own data file and can't read/write data from other users.
Also, a web app version will be implemented (a basic html/css/JavaScript frontend that connects with the backend -that will be online at this point- and a mobile app will also be availible (made on Python with the Kivy framework).

Thanks!
### Dan Almenar
