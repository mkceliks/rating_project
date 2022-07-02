# Rating Project For Universities ( Angular and Go )
This is gonna be a rating project for universities but it goes a movie website right now :).

For starting this project;

Download or Clone the code to your desktrop.
# open the web-interface
- <pre><code>ng serve --open</code></pre>
# build the main.go file
- <pre><code>go build</code></pre>
# run the backend server.
- <pre><code>go run main.go</code></pre>
Those all things should write to the terminal of the specific file or folder.

# 25.06.2022 

- Angular environment created.
- service, component, model files are created.
- First connection established with backend side.
- Just created a simple struct that has 3 data variables with string type.

# 26.06.2022

- Necessary imports made like js files,bootstrap.
- Dashboard component generated.
- Movie component,service and model generated and integrated with MongoDB data structs
- Episodes model and service generated to show the movies episodes and integrated with MongoDB data structs

# 27.06.2022

- CORS files added to the go file and synchronized with all system.
- addAnime() function created on backend side. And it integrated with the frontend side. ( Works well! )
- Product adding service and component added to the angular side. Retrieving the data from the POST method and sends it to the backend side. ( /add-anime )
- Navi,category and dashboard components created and updated.
- Necessery importations and changes made on app.module.ts, app.component.html, app-routing.module.ts ( npm Toastr, FormsModules, ReactiveFormsModule and Routings).

# 28.06.2022

- Product-add service created to integrate with backend. Movie service updated to get all the data from database about movies,animes and sports collections.
- Anime and Sport interfaces are created. These interfaces are same with the movie interface.
- Sports and Animes components are created and Movies component updated to just show movie collection.
- Navi, Dashboard and Category updated and integrated. Add button added to the navi. There are 3 forms in add buttons route. So add button can add the data to the database for 3 collection right now. ( Tested and Works Well! )
- Adding functions added to the main.go file and created the API's. 4 Structs added to the main.go file. All of them are integrated.
- Add-product component file edited to add data to 3 collections.
- Routing stuffs are OK! with angular side.

# 29.06.2022

- Episode adding feature added to the backend and frontend. The user can select a movie, anime or sport show from dropdown and can import an episode into it. ( Works Well! )
- The dropdown's options comes from 3 different collections which includes movies, animes and sport shows.
- Models created for episodes in frontend side.
- Tried to create a filter API but can't handle it with yet. There is some code about filtering by id to show episodes when the user clicks to the Watch button. ( NOT WORKING YET! )
- The button added to the navi for importing episodes into movies or something. This button routes to new page for user to import episodes.

# 02.07.2022

- Episode filtering and showing feature added for the sport section.
- Bugs fixed from -29.06.2022- ( filterAPI )
- Dashboard css updated.
- Episodes component and service updated, fixed and integrated with filterAPI.
- Unnecessery comment-lined files deleted from main.go file.
- Models and Server folders created.
- Struct models and functions ( setters, getters ) moved to models and server folders.
