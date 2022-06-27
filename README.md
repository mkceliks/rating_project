# Rating Project For Universities ( Angular and Go )
This is a rating project for universities

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

