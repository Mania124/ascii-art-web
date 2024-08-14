# ASCII-ART-WEB-DOCKERIZE

## Description
This is a dockerized web-based application developed using Go, HTML, CSS , and Javascript. The platform takes input text from the user then displays its graphical representation using ASCII art. You can share the docker image and run an instance (container) on various machines independent of the runtime environment. This is because each container has the environment variables, configuration, and files required to run the application.


## Usage
1. **Clone Repository**: Clone this repository to your local machine.
   ```bash
   $ git clone https://learn.zone01kisumu.ke/git/hshikuku/ascii-art-web-dockerize.git
   ```
2. **Install Docker and Go**: Ensure you have Go installed on your machine.(Go to official Go website and follow the instructions (based on your OS)). Also, install docker on your machine if it is not available.

3. **Run Program**: To get the app running, navigate to the directory where you've cloned the repo then use the following command:
    ```bash
    $ cd /ascii-art-web-dockerize
    $ ./bash.sh
    ```
    This will execute the instructions outlined in the dockerfile.

4.  **Launch GUI**: Type this command on the address bar of your browser:
    ```bash
    http://localhost:8080 
    ```
If you change the port value in the dockerfile remember to replace 8080 in this command with that value.

5. **Choose banner** : Select your banner of choice in the list provided.

6. **Input text**: In the text area labeled input text, type the text that you would like to be converted.

7. **Execute**: Click on the 'produce art' button. The program will execute and the results returned. If for example your input text is 'hello' and your banner of choice is 'standard' then the output will be:
    ```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
    ``` 
## Implementation Details

### Dockerfile
- This is a text file that contains a set of instructions used to create a Docker image by automating the steps of configuring the environment and installing the necessary software.

### GET request
- This type of request method is implemented when fetching home page template. 
- The response from the http server to the user delivers the home page which displays a form that expects user input. 
- This type of request method is also implemented when displaying the about ascii and about us pages.

### POST request
- The implementation of post method is carried out when sending form data to the server for implementation of the corresponding ascii art.
- Once the form data is parsed in the ASCIIArthandler and deemed to fulfill all the necessary conditions then the ascii art program is called to convert to its corresponding ascii art.
- This data in form of a string is then injected in a template and served back to the user in a readable format.

### Error handling
- The server uses logging functionalities as defined in package log to handle various errors. 
- In addition, the endpoints implement http methods such as http.StatusBadRequest, http.StatusNotFound, and http.StatusMethodNotAllowed that returns the appropriate status codes.

## File System
1. **Dockerfile**
    This is a text file that contains a set of instructions used by the docker engine to create a Docker image. It automates the steps of configuring the environment and installing the necessary software dependencies.
2. **bash.sh**
    This is a script that executes the commands for builing an image and running a container. It therefore gets the application up and running.
3. **dockerignore**
    It contains the files that should be ignored by docker engine.
4. **server.go**

    This is the entry point of the program. It registers the handler functions, serves the static files and also sets up the server.
     
5. **ascii-art**

    This directory contains the banner files functions used to create the graphical representation, and their test files.
-  **banners**: This sub-directory has the banner files that 
contain the ascii-art characters.

-  **art_map_creator.go**: Creates a map of ascii art characters from a chosen banner file.

-  **get_banner_file.go**: Gets the appropriate banner file that matches the choice. Example; if choice is 'standard' the function returns 'standard.txt'

-  **art_maker.go**: Converts the text and returns the output.

-  **integrity_checker.go**: Ensures that the banner files are available and in original condition.

-  **get_banner_file.go**: Retrieves the specific banner file according to the argument passed.

6. **handlers**
-  **handlers.go**: The file contains the handler functions used in the program.

7. **static**

    This directory contains static files such as images, css and js files.

8. **templates**

This contains various html templates used by the GUI.
-  **index.html**:  Displays the home page.

-  **result.html**: This is where the output is displayed.

-  **aboutus.html**:  This page contains a the name of the authors and thier github links.

-  **aboutascii.html**:  This page contains information about the concept of ascii art.

-  **400.html**: Displays http status code 400

-  **404.html**: Displays http status code 404

-  **500.html**: Displays https status code 500

9. **testCases**
  Contains the expected output used in TestHandleASCIIArt

10. **handlers_test.go**
   Tests the functionality of the handler functions.

## Authors

-   **[Hezborn Shikuku](https://github.com/Mania124/ascii-art-web)**
-   **[Rogers Ragwel](https://github.com/oragwelr/ascii-art-web)**
-   **[Rabin Otieno](https://github.com/Rabinnnn/ascii-art-web)**

## License

* This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.